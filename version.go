// Responsible for getting the version of an application and generating
// a `build-version.json` file to include in the applications' build process
// for the purpose of displaying version information.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

const buildVersionFile = "build-version.json"

type buildVersion struct {
	CommitHash        string `json:"CommitHash"`
	CurrentVersion    string `json:"currentVersion"`
	NextVersion       string `json:"nextVersion"`
	NextVersionReason string `json:"nextVersionReason"`
}

func versionMain(af *applicationFlags) error {
	// Default to the current working directory, or set it to arg 1 after the version sub command.
	repoPath, _ := os.Getwd()
	if len(af.args) > 0 {
		repoPath = af.args[0]
	}

	return BuildVersionFile(repoPath)
}

// BuildVersionFile build a JSON file with version info.
func BuildVersionFile(repoPath string) error {
	// Check the path exist.
	fileObj, err := os.Stat(repoPath)
	if os.IsNotExist(err) || !fileObj.IsDir() {
		return fmt.Errorf("repository path does not exists: %v", repoPath)
	}

	bvInfo := new(buildVersion)
	bvInfo.CurrentVersion = getCurrentVersion(repoPath)
	bvInfo.NextVersion, bvInfo.NextVersionReason = getNextVersion(repoPath, bvInfo.CurrentVersion)
	// Add commit hash.
	hash, err2 := getCommitHash(repoPath, bvInfo.CurrentVersion)
	if err2 != nil {
		return err2
	}
	bvInfo.CommitHash = hash

	bvJson, err1 := json.Marshal(bvInfo)
	if err1 != nil {
		return fmt.Errorf("could not JSON encode build version info, reason: %v", err1.Error())
	}

	// Write the build version info to a JSON file.
	bvFile := repoPath + PS + buildVersionFile
	if e := os.WriteFile(bvFile, bvJson, dirMode); e != nil {
		return fmt.Errorf("could not build %v, reason: %v", bvFile, e.Error())
	}

	return nil
}

// getCurrentVersion list all versions, sort by semantic version, then give you the one off the top.
func getCurrentVersion(repoPath string) (latestVersion string) {
	latestVersion = "HEAD"

	// Default to HEAD when no tag.
	sco, sce, exitCode, err3 := runRepoCmd(repoPath, "tag", "--sort=-version:refname")
	if err3 != nil {
		latestVersion = ""
		return
	}

	if sce == nil && exitCode == 0 {
		versionsData := bytes.Trim(sco, "\n")
		if len(versionsData) > 0 {
			// Split output into an array by newline.
			versions := bytes.Split(versionsData, []byte("\n"))
			latestVersion = string(versions[0])
		}
	}

	return
}

// getCommitHash returns the git commit has for the given tag.
func getCommitHash(repoPath, tag string) (commitHash string, err error) {
	sco, sce, exitCode, err1 := runRepoCmd(repoPath, "rev-list", "-n", "1", tag)
	if err1 != nil {
		return
	}

	if sce == nil && exitCode == 0 {
		commitHashData := bytes.Trim(sco, "\n")
		if len(commitHashData) > 0 {
			commitHash = string(commitHashData)
		}
	}
	return
}

// getNextVersion Get the next calculated semantic version.
func getNextVersion(repoPath, tag string) (nextVer, nextVerReason string) {
	nextVer = "0.1.0"
	revRange := ""
	// Reduce the revision range from the last tag to the most recent commit.
	if tag != "HEAD" {
		revRange = tag + "..HEAD"
	}

	var sco []byte
	var sce, err1 error
	var exitCode int
	// TODO: Handle really large amounts of git logs efficiently.
	// Look at all commit logs since the last tag (maybe only annotated):
	if revRange == "" {
		sco, sce, exitCode, err1 = runRepoCmd(repoPath, "log", "--format=medium")
	} else {
		sco, sce, exitCode, err1 = runRepoCmd(repoPath, "log", "--format=medium", revRange)
		if err1 != nil {
			_ = fmt.Errorf("error retrieving git logs: %v", err1.Error())
			return
		}
	}

	if sce != nil && exitCode != 0 {
		return
	}

	commitLogs := string(sco)
	if len(commitLogs) < 0 {
		return
	}

	re := regexp.MustCompile(`rel:\s(\d+\.\d+\.\d+)`)
	res := re.FindStringSubmatch(commitLogs)
	// Look for commit message format "rel: x.x.x"
	if len(res) > 0 {
		nextVerReason = "`rel:` type was found in the git logs from the last release to the current HEAD"
		nextVer = res[1]
		return
	}

	ver := strings.Split(tag, ".")
	if tag == "HEAD" || len(ver) < 3 {
		nextVerReason = "no previous tags detected"
		return
	}

	// Look for "BREAKING CHANGE" to increment major version
	if strings.Contains(commitLogs, "BREAKING CHANGE\n") {
		// TODO: Use a lib to handle incrementing the semantic version number.
		ver[0], nextVer, nextVerReason = incrementNumber(ver[0], nextVer, "`BREAKING CHANGE` keyword found in git logs")
		ver[1] = "0"
		ver[2] = "0"

	} else if strings.Contains(commitLogs, "add: ") { // Look for "add:" to increment the minor version
		ver[1], nextVer, nextVerReason = incrementNumber(ver[1], nextVer, "add: keyword found in git logs")
		ver[2] = "0"

	} else { // Increment patch version
		ver[2], nextVer, nextVerReason = incrementNumber(ver[2], nextVer, "no new features or breaking changed detected in the git logs")
	}
	nextVer = strings.Join(ver, ".")

	return
}

// incrementNumber add 1 to a numeric string, on failure return numeric number and the reason it failed.
func incrementNumber(a, nv, nvr string) (string, string, string) {
	ret, err1 := strconv.ParseInt(a, 10, 64)
	if err1 != nil {
		return a, "", fmt.Sprintf("unable to increment %q numeric string by 1, reason: %v", a, err1.Error())
	}

	return strconv.FormatInt(ret+1, 10), nv, nvr
}

// runRepoCmd run a command against the repository.
func runRepoCmd(repoPath string, args ...string) (cmdOut []byte, cmdErr error, exitCode int, err error) {
	// Remember the current working directory
	cwd, err1 := os.Getwd()
	defer os.Chdir(cwd)
	if err1 != nil {
		err = err1
		return
	}
	// Change into the repository directory.
	err2 := os.Chdir(repoPath)
	if err2 != nil {
		err = err2
		return
	}
	// Run an arbitrary git command.
	cmd := exec.Command("git", args...)
	cmdOut, cmdErr = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()

	return
}

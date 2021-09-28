// Responsible for getting the version of an application and generating
// a `version.go` file to include in the applications' go:build process
// for the purpose of displaying version information.

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
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
	//
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

func getNextVersion(repoPath, tag string) (nextVer, nextVerReason string) {
	nextVer = "0.1.0"
	revRange := ""
	// Reduce the revision range from the last tag to the most recent commit.
	if tag != "HEAD" {
		revRange = tag + "..HEAD"
	}
	// TODO: Handle really large amounts of git logs efficiently.
	// Look at all commit logs since the last tag (maybe only annotated):
	sco, sce, exitCode, err1 := runRepoCmd(repoPath, "log", "--format=medium", revRange)
	if err1 != nil {
		return
	}

	if sce == nil && exitCode == 0 {
		commitLogs := string(sco)
		if len(commitLogs) > 0 {
			ver := strings.Split(tag, ".")
			// TODO: Look for "BREAKING CHANGE" to increment major version
			if strings.Contains(commitLogs, "BREAKING CHANGE\n") {
				nextVerReason = "BREAKING CHANGE keyword found in git logs."
				major, err2 := strconv.ParseInt(ver[0], 10, 64)
				if err2 != nil {
					nextVer = ""
					nextVerReason = "unable to convert string to int"
				}
				ver[0] = strconv.FormatInt(major+1, 10)
				// Look for "add:" to increment the minor version
			} else if strings.Contains(commitLogs, "add:") {
				nextVerReason = "add: keyword found in git logs."
				major, err2 := strconv.ParseInt(ver[1], 10, 64)
				if err2 != nil {
					nextVer = ""
					nextVerReason = "unable to convert string to int"
				}
				ver[1] = strconv.FormatInt(major+1, 10)
			} else {
				nextVerReason = "no new features on breaking changed detected in the git logs"
				major, err2 := strconv.ParseInt(ver[2], 10, 64)
				if err2 != nil {
					nextVer = ""
					nextVerReason = "unable to convert string to int"
				}
				ver[2] = strconv.FormatInt(major+1, 10)
			}
			nextVer = strings.Join(ver, ".")
		}
	}

	return
}

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
		fmt.Printf("\nerr2: %v\n", err2.Error())
		return
	}
	// Run an arbitrary git command.
	cmd := exec.Command("git", args...)
	cmdOut, cmdErr = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()

	return
}

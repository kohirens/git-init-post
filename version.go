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

	bvInfo := new(buildVersion)
	bvInfo.CurrentVersion = getCurrentVersion(repoPath)
	bvInfo.NextVersion, bvInfo.NextVersionReason = getNextVersion(repoPath)
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
	sco, sce, exitCode, err3 := runRepoCmd(repoPath, "tag", "--sort=-version:refname")
	if err3 != nil {
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

func getNextVersion(repoPath string) (nextVer, nextVerReason string) {
	nextVer = "0.1.0"

	// TODO: Look at all commit logs since the last tag (maybe only annotated):
	// TODO: Look for "BREAKING CHANGE" to increment major version
	// TODO: Look for "add:" to increment the minor version
	// TODO: Otherwise increment the patch version.

	return
}

func runRepoCmd(repoPath string, args...string) (cmdOut []byte, cmdErr error, exitCode int, err error) {
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
	// Run an arbitrary the git command.
	cmd := exec.Command("git", args...)
	cmdOut, cmdErr = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()

	return
}
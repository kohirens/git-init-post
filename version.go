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
	// default to the current working directory, or set it to arg 1 after the version sub command.
	repoPath, _ := os.Getwd()
	if len(af.args) > 0 {
		repoPath = af.args[0]
	}

	bvInfo := new(buildVersion)
	bvInfo.CurrentVersion = getVersion(repoPath)
	bvInfo.NextVersion, bvInfo.NextVersionReason = getNextVersion(repoPath)
	// TODO: Add commit hash.


	bvJson, err1 := json.Marshal(bvInfo)
	if err1 != nil {
		return fmt.Errorf("could not JSON encode build version info, reason: %v", err1.Error())
	}

	// Write the build version info to a json file.
	bvFile := repoPath + PS + buildVersionFile
	if e := os.WriteFile(bvFile, bvJson, dirMode); e != nil {
		return fmt.Errorf("could not build %v, reason: %v", bvFile, e.Error())
	}

	return nil
}

// getVersion list all versions, sort by semantic version, then give you the one off the top.
func getVersion(repoPath string) (latestVersion string) {

	// Remember the current working directory
	cwd, err1 := os.Getwd()
	defer os.Chdir(cwd)
	if err1 != nil {
		return
	}
	// Change into the repository directory.
	err2 := os.Chdir(repoPath)
	if err2 != nil {
		return
	}

	// Get and sort from the highest to the lowest version.
	cmd := exec.Command("git", "tag", "--sort=-version:refname")

	sco, sce := cmd.CombinedOutput()

	exitCode := cmd.ProcessState.ExitCode()

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

func getCurrentVersion() string {
	currVer := ""

	return currVer
}

func getNextVersion(repoPath string) (nextVer, nextVerReason string) {
	nextVer = "0.1.0"

	// TODO: Look at all commit logs since the last tag (maybe only annotated):
	// TODO: Look for "BREAKING CHANGE" to increment major version
	// TODO: Look for "add:" to increment the minor version
	// TODO: Otherwise increment the patch version.

	return
}

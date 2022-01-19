package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

// getLatestChanges Get changes since the specified tag.
func getLatestChanges(repoPath, commitRange string) ([]byte, error) {
	revRange := "-1" // Default to last commit.
	// set a user specified revision range .
	if commitRange != "" {
		revRange = commitRange
	}

	var sco []byte
	var sce, err1 error
	var exitCode int

	sco, sce, exitCode, err1 = runRepoCmd(repoPath, "log", "--format=medium", revRange)
	if err1 != nil {
		return nil, fmt.Errorf("error retrieving git logs for %s: %v", revRange, err1.Error())
	}

	if sce != nil && exitCode != 0 {
		return nil, fmt.Errorf("sce error retrieving git logs: %v", sce.Error())
	}

	return sco, nil
}

// hasUnreleasedCommitsWithTags Looks for tag in each unreleased commit line which will indicate if there are changes to tag.
func hasUnreleasedCommitsWithTags(repoPath string, af *applicationFlags) bool {
	retVal := false
	repoLogs, err := getLatestChanges(repoPath, *af.taggable.commitRange)

	if err != nil {
		return retVal
	}

	relReg := regexp.MustCompile(`[ \t]+rel:\s{1,4}(\d+\.\d+\.\d+)`)
	tagReg := regexp.MustCompile(`[ \t]+(add|rmv|chg|fix|dep):\s{1,4}.[^\n]+`)

	commitLogs := string(repoLogs)
	// verbosity
	if af.taggable.verbose {
		fmt.Printf("here are the logs:\n%v", commitLogs)
	}
	// Look for commit message format "rel: x.x.x"
	res := relReg.FindStringSubmatch(commitLogs)
	if len(res) > 0 {
		retVal = true
	}
	// Look for commit message format "<tag>: "
	res2 := tagReg.FindStringSubmatch(commitLogs)
	if len(res2) > 0 {
		retVal = true
	}

	return retVal
}

// runRepoCmd run a command against the repository.
func runRepoCmd(repoPath string, args ...string) (cmdOut []byte, cmdErr error, exitCode int, err error) {
	cmd := exec.Command("git", args...)
	cmd.Env = os.Environ()
	cmd.Dir = repoPath
	cmdOut, cmdErr = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()
	return
}

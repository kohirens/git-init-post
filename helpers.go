package main

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

// getLatestChanges Get changes since the specified tag.
func getLatestChanges(repoPath, tag string) ([]byte, error) {
	revRange := ""
	// Reduce the revision range from the last tag to the most recent commit.
	if tag != "HEAD" {
		revRange = tag + "..HEAD"
	}

	var sco []byte
	var sce, err1 error
	var exitCode int
	// Look at all commit logs since the last tag (maybe only annotated) and dump into a file:
	if revRange == "" {
		sco, sce, exitCode, err1 = runRepoCmd(repoPath, "log", "--format=medium")
	} else {
		sco, sce, exitCode, err1 = runRepoCmd(repoPath, "log", "--format=medium", revRange)
		if err1 != nil {
			return nil, fmt.Errorf("error retrieving git logs for %s: %v", revRange, err1.Error())
		}
	}
	if sce != nil && exitCode != 0 {
		return nil, fmt.Errorf("sce error retrieving git logs: %v", sce.Error())
	}

	return sco, nil
}

// hasTags Looks for tag in each commit line to indicate if there are changes to tag.
func hasTags(repoPath, tag string) bool {
	retVal := false
	repoLogs, err := getLatestChanges(repoPath, tag)

	if err != nil {
		return retVal
	}

	relReg := regexp.MustCompile(`[ \t]+rel:\s{1,4}(\d+\.\d+\.\d+)`)
	tagReg := regexp.MustCompile(`[ \t]+(add|rmv|chg|fix|dep):\s{1,4}.[^\n]+`)

	commitLogs := string(repoLogs)
	// Look for commit message format "rel: x.x.x"
	res := relReg.FindStringSubmatch(commitLogs)
	if len(res) > 0 {
		retVal = true
	}
	// Look for commit message format "<tag>>: "
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

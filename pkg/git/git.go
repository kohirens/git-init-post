package git

import (
	"fmt"
	"github.com/kohirens/stdlib/cli"
	"github.com/kohirens/stdlib/log"
	"regexp"
)

// getLatestChanges Get changes since the specified tag.
func getLatestChanges(repoPath, commitRange string) ([]byte, error) {
	revRange := "-1" // Default to last commit.
	// set a user specified revision range .
	if commitRange != "" {
		revRange = commitRange
	}

	//so, se, _, err1 := RunRepoCmd(repoPath, "log", "--format=medium", revRange)
	so, se, _ := cli.RunCommand("git", repoPath, "log", "--format=medium", revRange)
	if se != nil {
		return nil, fmt.Errorf("could not retrieve git logs for %s: %s; %s\n", revRange, so, se.Error())
	}

	if se != nil {
		return nil, fmt.Errorf("se error retrieving git logs: %v", se.Error())
	}

	return so, nil
}

// HasUnreleasedCommitsWithTags Looks for special lines in each commit message, in the revision range, which will indicate if there are changes to tag.
func HasUnreleasedCommitsWithTags(repoPath string, commitRange string) bool {
	repoLogs, err := getLatestChanges(repoPath, commitRange)

	if err != nil {
		return false
	}

	retVal := false
	relReg := regexp.MustCompile(`[ \t]+rel:\s{1,4}(\d+\.\d+\.\d+)`)
	tagReg := regexp.MustCompile(`[ \t]+(add|rmv|chg|fix|dep):\s{1,4}.[^\n]+`)

	commitLogs := string(repoLogs)

	log.Infof("here are the logs:\n%v", commitLogs)

	// Look for commit message format "rel: x.x.x"
	if r := relReg.FindStringSubmatch(commitLogs); len(r) > 0 {
		retVal = true
	}
	// Look for commit message format "<tag>: "
	if r := tagReg.FindStringSubmatch(commitLogs); len(r) > 0 {
		retVal = true
	}

	return retVal
}

package main

import (
	"os"
)

func IsTaggable(af *applicationFlags) bool {
	// Default to the current working directory, or set it from the flag.
	repoPath, _ := os.Getwd()
	if len(af.taggable.repo) > 0 {
		repoPath = af.taggable.repo
	}

	infof("repoPath = %q\n", repoPath)

	revRange := af.taggable.commitRange
	if af.taggable.commitRange == "" {
		revRange = getRevisionRange(repoPath)
	}

	return hasUnreleasedCommitsWithTags(repoPath, revRange)
}

// getRevisionRange determines the revision range based on current.
func getRevisionRange(repoPath string) string {
	currVersion := getCurrentVersion(repoPath)
	revRange := currVersion
	// currVersion defaults to "HEAD" and so get all the commits, if not, then
	if currVersion != "HEAD" {
		// only grab commits from the last tag up-to and including the HEAD.
		revRange = currVersion + "..HEAD"
	}

	return revRange
}

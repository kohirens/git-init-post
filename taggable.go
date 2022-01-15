package main

import (
	"os"
)

func IsTaggable(af *applicationFlags) bool {
	// Default to the current working directory, or set it from the repo flag.
	repoPath, _ := os.Getwd()
	if len(*af.taggable.repo) > 0 {
		repoPath = *af.taggable.repo
	}

	return hasUnreleasedCommitsWithTags(repoPath, *af.taggable.commitRange)
}

package main

import (
	"fmt"
	"os"
)

func IsTaggable(af *applicationFlags) bool {
	// Default to the current working directory, or set it from the flag.
	repoPath, _ := os.Getwd()
	if len(af.taggable.repo) > 0 {
		repoPath = af.taggable.repo
	}

	if af.taggable.verbose {
		fmt.Printf("repoPath = %q\n", repoPath)
	}
	return hasUnreleasedCommitsWithTags(repoPath, af.taggable.commitRange, af.taggable.verbose)
}

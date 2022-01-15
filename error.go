package main

var errors = struct {
	invalidCommitRange,
	subCmdMissing string
}{
	invalidCommitRange: "invalid commit range format. please use for example `commit1..commit2`. Exceptable values are a branch, tag, or hash.",
	subCmdMissing:      "missing a sub command, either \"version\" or \"taggable\"",
}

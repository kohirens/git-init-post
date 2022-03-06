package main

var errors = struct {
	invalidCommitRange,
	subCmdMissing,
	packageNameRequired,
	semverFormtatInvalid string
}{
	invalidCommitRange:   "invalid commit range format. please use for example `commit1..commit2`. Acceptable values are a branch, tag, or hash.",
	subCmdMissing:        "missing a sub command, either \"semver\" or \"taggable\"",
	packageNameRequired:  "packageName flag is required when --format equals \"go\"",
	semverFormtatInvalid: "invalid format set, options are \"go\" or \"JSON\"",
}

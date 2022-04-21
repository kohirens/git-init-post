package main

var errors = struct {
	invalidCommitRange,
	subCmdMissing,
	packageNameRequired,
	semverFormatInvalid,
	noConfPath string
}{
	invalidCommitRange:  "invalid commit range format. please use for example `commit1..commit2`. Acceptable values are a branch, tag, or hash.",
	subCmdMissing:       "missing a sub command, either \"semver\" or \"taggable\" or \"checkConf\"",
	packageNameRequired: "packageName flag is required when --format equals \"go\"",
	semverFormatInvalid: "invalid format set, options are \"go\" or \"JSON\"",
	noConfPath:          "invalid path for git-chglog config.",
}

package main

var usageMsgs = map[string]string{
	"taggable.verbose":   "Turn on verbose output.",
	"semver.save":        "File to save the output to; for example semver.json",
	"semver":             "sub-command generates a file with repository tag information, such as current tag and the next calculated tag (based on special tags in unreleased commits messages).",
	"taggable":           "sub-command returns a `true` or `false` if commits searched indicate they constitute a new release or not respectively. The values `true` and `false` were chosen to avoid confusion with program exit status",
	"commitRange":        "Commit range to search for special tags. Default to only the last commit.",
	"repo":               "Path to a repository for commits to be searched, if not set, then defaults to current directory.",
	"help":               "Print this help info",
	"version":            "Print version information",
	"subCommands":        "\n  sub-commands:\n\n    semver\n\tsee semver --help\n    taggable\n\tsee taggable --help\n    checkConf\n\tsee checkConf --help\n",
	"semver.packageName": "The package name to use for the Go code generated, this option is only used when \"--format go\"",
	"semver.format":      "The format of the semantic version output; options are \"JSON\" or \"go\"",
	"semver.varName":     "Variable name to use in generated code, default to \"appFlags\"",
	"checkConf":          "sub-command runs a check for the git-chglog configuration adds when missing",
	"checkConf.path":     "path to the configuration",
	"checkConf.repo":     "repo to name in the generated configuration, should it be missing.",
}

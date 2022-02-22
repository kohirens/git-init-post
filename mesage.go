package main

var usageMsgs = map[string]string{
	"help.usage":       "",
	"taggable.v":       "Turn on verbose output.",
	"taggable.verbose": "Turn on verbose output.",
	"semver.save":      "Filename to save the JSON output to; for example semver.json",
	"semver":           "sub-command generates a file with repository tag information, such as current tag and the next calculated tag (based on special tags in unreleased commits messages).",
	"taggable":         "sub-command returns a `true` or `false` if commits searched indicate they constitute a new release or not respectively. The values `true` and `false` were chosen to avoid confusion with program exit status",
	"commitRange":      "Commit range to search for special tags. Default to only the last commit.",
	"repo":             "Path to a repository for commits to be searched, if not set, then defaults to current directory.",
}

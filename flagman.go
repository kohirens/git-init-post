package main

// All flag usage/instructions/documentation goes in here.

var flagUsages = map[string]string{
	"help":        "Display usage info for all arguments, flags, and subcommands.",
	"version":     "version\n\tsub-command generates a file with repository tag information, such as current tag and the next calculated tag (based on special tags in unreleased commits messages).",
	"taggable":    "taggable\n\tsub-command returns a `true` or `false` if commits searched indicate they constitute a new release or not respectively. The values `true` and `false` were chosen to avoid confusion with program exit status",
	"commitRange": "-commitRange\n\tCommit range to search for special tags. Default to only the last commit.",
	"repo":        "-repo\n\tPath to a repository for commits to be searched, if not set, then defaults to current directory.",
}

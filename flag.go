package main

import (
	"flag"
	"fmt"
	"regexp"
)

type applicationFlags struct {
	args     []string
	subCmd   string
	version  *versionSubCmd
	taggable *taggableSubCmd
}

type taggableSubCmd struct {
	fs          *flag.FlagSet
	commitRange string
	repo        string
	verbose     bool
}

type versionSubCmd struct {
	fs   *flag.FlagSet
	repo string
}

const (
	taggable = "taggable"
	version  = "version"
)

// define All application flags.
func (af *applicationFlags) define() {
	// version sub-command
	appFlags.version = &versionSubCmd{
		fs: flag.NewFlagSet("version", flag.ContinueOnError),
	}
	af.version.fs.StringVar(&af.version.repo, "repo", "", flagUsages["repo"])
	// taggable sub-command
	af.taggable = &taggableSubCmd{
		fs: flag.NewFlagSet(taggable, flag.ExitOnError),
	}
	af.taggable.fs.StringVar(&af.taggable.commitRange, "commitRange", "", flagUsages["commitRange"])
	af.taggable.fs.StringVar(&af.taggable.repo, "repo", "", flagUsages["repo"])
	af.taggable.fs.BoolVar(&af.taggable.verbose, "v", false, usageMsgs["taggableVerbose"])
	af.taggable.fs.BoolVar(&af.taggable.verbose, "verbose", false, usageMsgs["taggableVerbose"])
}

// check Verify that all flags are set appropriately.
func (af *applicationFlags) check() error {
	if len(af.args) == 0 {
		return fmt.Errorf(errors.subCmdMissing)
	}

	if af.subCmd == taggable {
		rangeFmt := regexp.MustCompile(`[a-zA-Z0-9\.\-_]+\.\.[a-zA-Z0-9\.\-_]+`)
		res := rangeFmt.FindStringSubmatch(af.taggable.commitRange)
		if af.taggable.commitRange != "HEAD" && af.taggable.commitRange != "" && res == nil {
			return fmt.Errorf(errors.invalidCommitRange)
		}
	}

	return nil
}

// parse Pass all remaining non-flag arguments along to the application.
func (af *applicationFlags) parse(cliArgs []string) {
	if len(cliArgs) > 0 {
		af.args = cliArgs[0:]
	}
}

// parseSubcommands Determines if a sub-command was given.
func (af *applicationFlags) parseSubcommands() error {

	if len(af.args) < 1 {
		return nil
	}

	switch af.args[0] {
	case version:
		af.subCmd = version
		return af.version.fs.Parse(af.args[1:])
	case taggable:
		af.subCmd = taggable
		return af.taggable.fs.Parse(af.args[1:])
	}

	return nil
}

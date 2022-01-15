package main

import (
	"flag"
	"fmt"
	"regexp"
)

type applicationFlags struct {
	args     []string
	subCmd   string
	version  *flag.FlagSet
	taggable *taggableSubCmd
}

type taggableSubCmd struct {
	fs          *flag.FlagSet
	commitRange *string
	repo        *string
}

const (
	taggable = "taggable"
	version  = "version"
)

// define All application flags.
func (af *applicationFlags) define() {
	// version sub-command
	appFlags.version = flag.NewFlagSet("version", flag.ContinueOnError)
	// taggable sub-command
	af.taggable = &taggableSubCmd{
		fs: flag.NewFlagSet(taggable, flag.ExitOnError),
	}
	af.taggable.commitRange = af.taggable.fs.String("commitRange", "", flagUsages["taggableCommitRange"])
	af.taggable.repo = af.taggable.fs.String("repo", "", flagUsages["taggableRepo"])
}

// check Verify that all flags are set appropriately.
func (af *applicationFlags) check() error {
	if len(af.args) == 0 {
		return fmt.Errorf(errors.subCmdMissing)
	}

	if af.subCmd == taggable {
		rangeFmt := regexp.MustCompile(`[a-zA-Z0-9\.\-_]+\.\.[a-zA-Z0-9\.\-_]+`)
		res := rangeFmt.FindStringSubmatch(*af.taggable.commitRange)
		if *af.taggable.commitRange != "" && res == nil {
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

	switch af.args[0] {

	case version:
		af.subCmd = version
	case taggable:
		af.subCmd = taggable
		return af.taggable.fs.Parse(af.args[1:])
	}

	return nil
}

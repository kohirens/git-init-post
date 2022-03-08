package main

//go:generate git-tool-belt semver -save info.go -format go -packageName main -varName appConfig

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

type applicationFlags struct {
	args           []string
	help           bool
	subCmd         string
	semver         *semverSubCmd
	taggable       *taggableSubCmd
	version        bool
	CurrentVersion string
	CommitHash     string
}

type taggableSubCmd struct {
	fs          *flag.FlagSet
	commitRange string
	repo        string
	verbose     bool
}

type semverSubCmd struct {
	format      string
	fs          *flag.FlagSet
	packageName string
	repo        string
	save        string
	varName     string
}

const (
	taggable = "taggable"
	cSemver  = "semver"
)

// define All application flags.
func (af *applicationFlags) define() {
	flag.BoolVar(&af.help, "h", false, "")
	flag.BoolVar(&af.help, "help", false, usageMsgs["help"])
	flag.BoolVar(&af.version, "v", false, "")
	flag.BoolVar(&af.version, "version", false, usageMsgs["version"])
	// semver sub-command
	af.semver = &semverSubCmd{
		fs: flag.NewFlagSet("semver", flag.ContinueOnError),
	}
	af.semver.fs.StringVar(&af.semver.repo, "repo", "", usageMsgs["repo"])
	af.semver.fs.StringVar(&af.semver.save, "save", "", usageMsgs["semver.save"])
	af.semver.fs.StringVar(&af.semver.packageName, "packageName", "", usageMsgs["semver.packageName"])
	af.semver.fs.StringVar(&af.semver.format, "format", "JSON", usageMsgs["semver.format"])
	af.semver.fs.StringVar(&af.semver.varName, "varName", "appFlags", usageMsgs["semver.varName"])
	// taggable sub-command
	af.taggable = &taggableSubCmd{
		fs: flag.NewFlagSet(taggable, flag.ExitOnError),
	}
	af.taggable.fs.StringVar(&af.taggable.commitRange, "commitRange", "", usageMsgs["commit.range"])
	af.taggable.fs.StringVar(&af.taggable.repo, "repo", "", usageMsgs["repo"])
	af.taggable.fs.BoolVar(&af.taggable.verbose, "v", false, "")
	af.taggable.fs.BoolVar(&af.taggable.verbose, "verbose", false, usageMsgs["taggable.verbose"])
}

// check Verify that all flags are set appropriately.
func (af *applicationFlags) check() error {
	if appFlags.help {
		flag.PrintDefaults()
		fmt.Printf(usageMsgs["subCommands"])
		os.Exit(0)
	}

	if appFlags.version {
		fmt.Printf("%v, %v\n", af.CurrentVersion, af.CommitHash)
		os.Exit(0)
	}

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

	if af.subCmd == cSemver {
		if af.semver.format != "go" && af.semver.format != "JSON" {
			return fmt.Errorf(errors.semverFormtatInvalid)
		}
		if af.semver.format == "go" && af.semver.packageName == "" {
			return fmt.Errorf(errors.packageNameRequired)
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
	case cSemver:
		af.subCmd = cSemver
		return af.semver.fs.Parse(af.args[1:])
	case taggable:
		af.subCmd = taggable
		return af.taggable.fs.Parse(af.args[1:])
	}

	return nil
}

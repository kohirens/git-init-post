package main

import (
	"flag"
)

type applicationFlags struct {
	args    []string
	subCmd  string
	version *flag.FlagSet
}

func (af *applicationFlags) parse(cliArgs []string) {
	// Pass all remaining non-flag arguments along to the application.
	if len(cliArgs) > 1 {
		af.args = cliArgs[1:]
	}
}

func (af *applicationFlags) parseSubcommands(cliArgs []string) error {

	// Determine if a sub-command was given.
	switch cliArgs[0] {

	case "version":
		af.subCmd = "version"
	}

	return nil
}

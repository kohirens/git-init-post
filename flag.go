package main

import (
	"flag"
)

type applicationFlags struct {
	version *flag.FlagSet
	subCmd string
}

func (af *applicationFlags) parse(cliArgs []string) {
}

func (af *applicationFlags) parseSubcommands(cliArgs []string) error {
	switch cliArgs[0] {

	case "version":
		af.subCmd = "version"
	}

	return nil
}
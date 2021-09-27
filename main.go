package main

import (
	"flag"
	"log"
	"os"
)

const (
	PS      = string(os.PathSeparator)
	dirMode = 0774
)

var appFlags = new(applicationFlags)

func init() {
	appFlags.version = flag.NewFlagSet("version", flag.ContinueOnError)
}

func main() {
	var mainErr error

	defer func() {
		if mainErr != nil {
			log.Fatalln(mainErr)
		}
		os.Exit(0)
	}()

	flag.Parse()

	appFlags.parse(flag.Args())
	e := appFlags.parseSubcommands(flag.Args())
	if e != nil {
		mainErr = e
		return
	}

	if appFlags.subCmd == "version" {
		if e := versionMain(appFlags); e != nil {
			mainErr = e
		}
	}
}

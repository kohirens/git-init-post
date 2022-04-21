package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	PS      = string(os.PathSeparator)
	dirMode = 0774
)

// appFlags Is what you use at runtime, it is the implementation of the applicationFlags type.
var appFlags = new(applicationFlags)

func init() {
	appFlags.define()
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

	if e := appFlags.parseSubcommands(); e != nil {
		mainErr = e
		return
	}

	if e := appFlags.check(); e != nil {
		mainErr = e
		return
	}

	if appFlags.subCmd == cSemver {
		if e := semverMain(appFlags); e != nil {
			mainErr = e
		}
	}

	if appFlags.subCmd == taggable {
		fmt.Printf("%v", IsTaggable(appFlags))
		return
	}

	if appFlags.subCmd == cCheckConfSubCmd {
		if e := addMissingChgLogConfig(appFlags.checkConfSubCmd.path, appFlags.checkConfSubCmd.repo); e != nil {
			mainErr = e
			return
		}
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	PS = string(os.PathSeparator)
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

	e := appFlags.parseSubcommands(flag.Args())
	if e != nil  {
		mainErr = e
		return
	}

	if appFlags.subCmd == "version" {
		fmt.Println(getVersion())
	}
}

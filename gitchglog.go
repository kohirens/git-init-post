package main

import (
	"fmt"
	"github.com/kohirens/stdlib"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// AddMissingChgLogConfig Check for a git-chglog configuration at the specified location, and add when missing.
func addMissingChgLogConfig(conf, gitUrl string) error {
	wd, e1 := os.Getwd()

	if e1 != nil {
		return e1
	}

	infof("working directory = %s\n", wd)

	if stdlib.PathExist(conf) {
		return nil
	}

	confDir := filepath.Dir(conf)
	fmt.Printf("no changelog file found at %q, adding the default...\n", confDir)

	if e := os.MkdirAll(confDir, dirMode); e != nil {
		return fmt.Errorf("clould not make dir %q", confDir)
	}

	outStr := strings.Replace(gitChgLogConf, "${VCS_URL}", "'"+gitUrl+"'", 1)

	if e := ioutil.WriteFile(conf, []byte(outStr), dirMode); e != nil {
		return e
	}
	fmt.Printf("added git-chglog config at %q\n", conf)

	if e := ioutil.WriteFile(confDir+"/CHANGELOG.tpl.md", []byte(gitChgLogChangelog), 0774); e != nil {
		return e
	}

	if !stdlib.PathExist(conf) {
		return fmt.Errorf("unable to add %q", conf)
	}

	infof("found git-chglog config here %q\n", conf)

	return nil
}

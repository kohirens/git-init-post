package main

import (
	"fmt"
	"github.com/kohirens/stdlib"
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
		fmt.Println("found git-chglog configuration")
		return nil
	}

	confDir := filepath.Dir(conf)
	logf("no changelog file found at %q, adding the default...", confDir)

	if e := os.MkdirAll(confDir, dirMode); e != nil {
		return fmt.Errorf("clould not make dir %q", confDir)
	}

	url := convertGitToHttp(gitUrl)
	outStr := strings.Replace(gitChgLogConf, "${VCS_URL}", "'"+url+"'", 1)

	if e := os.WriteFile(conf, []byte(outStr), dirMode); e != nil {
		return e
	}
	logf("added git-chglog config at %q\n", conf)

	if e := os.WriteFile(confDir+"/CHANGELOG.tpl.md", []byte(gitChgLogChangelog), 0774); e != nil {
		return e
	}

	if !stdlib.PathExist(conf) {
		return fmt.Errorf("unable to add %q", conf)
	}

	fmt.Println("added git-chglog configuration")

	return nil
}

func convertGitToHttp(url string) string {

	if strings.HasPrefix(url, "git@") {
		url = strings.Replace(url, ":", "/", 1)
		url = strings.Replace(url, "git@", "https://", 1)
	}
	return url
}

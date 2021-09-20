// Responsible for getting the version of an application and generating
// a `version.go` file to include in the applications' go:build process
// for the purpose of displaying version information.

package main

import (
	"fmt"
	"os/exec"
)

func getVersion() string {
	nextVersion := "0.1.0"

	cmd := exec.Command("git", "tag")

	sco, sce := cmd.CombinedOutput()

	exitCode := cmd.ProcessState.ExitCode()

	if sce == nil && exitCode == 0 && len(sco) > 2 {
		fmt.Printf("sco: %q\n", sco)
		// TODO: Parse the current version
	} else {
	}

	return nextVersion
}
// Responsible for getting the version of an application and generating
// a `version.go` file to include in the applications' go:build process
// for the purpose of displaying version information.

package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

// list all versions, sort by semantic version, then give you the one off the top.
func getVersion() string {
	nextVersion := ""

	cmd := exec.Command("git", "tag")

	sco, sce := cmd.CombinedOutput()

	exitCode := cmd.ProcessState.ExitCode()

	if sce == nil && exitCode == 0 {
		versions := bytes.Trim(sco,"\n")
		fmt.Printf("versions: %q\n", versions)

		if len(versions) > 0 {
			_ = bytes.Split(versions, []byte("\n"))
		}

		// TODO: split output into an array by newline.
		// TODO: Sort by semantic version.
		// TODO: Return the latest one.
	}

	return nextVersion
}


func getCurrentVersion() string {
	currVer := ""

	return currVer
}

func getNextVersion() string {
	nextVer := "0.1.0"

	return nextVer
}
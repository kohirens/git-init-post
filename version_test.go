package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestVersionSubCmd(tester *testing.T) {
	var tests = []struct {
		name     string
		wantCode int
		args     []string
	}{
		{"versionSubCmd", 0, []string{"version"}},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			_ = setupARepository("repo-01")

			cmd := runMain("TestAppMain", test.args)

			cmdOut, cmdErr := cmd.CombinedOutput()

			got := cmd.ProcessState.ExitCode()


			fmt.Printf("exitCode = %v", got)
			if cmdOut != nil {
				fmt.Printf("\nBEGIN sub-command stdout:\n%q", string(cmdOut))
				fmt.Print("END sub-command\n")
			}

			if cmdErr != nil {
				fmt.Printf("\nBEGIN sub-command stderr:\n%q", cmdErr.Error())
				fmt.Print("END sub-command\n")
			}

			if got != test.wantCode {
				t.Errorf("got %q, want %q", got, test.wantCode)
			}
		})
	}
}


func setupARepository(dirName string) string {
	tmpRepoPath := testTmp + PS + dirName

	fmt.Printf("tmpRepoPath = %q\n", tmpRepoPath)

	cmd := exec.Command("git", "init", "--initial-branch", "main", tmpRepoPath)
	_, _ = cmd.CombinedOutput()
	if ec := cmd.ProcessState.ExitCode(); ec != 0 {
		panic("error initializing a temporary repo for a unit test")
	}

	dh, err := os.OpenFile(tmpRepoPath + PS + "README.md", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, dirMode)
	if err != nil {
		panic("error initializing a file in a temporary repo for a unit test")
	}
	dh.Write([]byte("# Testing"))
	dh.Close()

	wd, _ := os.Getwd()
	_ = os.Chdir(tmpRepoPath)

	cmd = exec.Command("git", "add", ".")
	_, _ = cmd.CombinedOutput()
	if ec := cmd.ProcessState.ExitCode(); ec != 0 {
		panic("error adding file to a temporary repo for a unit test")
	}

	cmd = exec.Command("git", "config", "user.email", "gittoolbelt@example.com")
	_, _ = cmd.CombinedOutput()
	if ec := cmd.ProcessState.ExitCode(); ec != 0 {
		panic("error initializing an email for a temporary repo for a unit test")
	}

	cmd = exec.Command("git", "config", "user.name", "Git Toolbelt")
	_, _ = cmd.CombinedOutput()
	if ec := cmd.ProcessState.ExitCode(); ec != 0 {
		panic("error initializing a name for a temporary repo for a unit test")
	}

	cmd = exec.Command("git", "commit", "-m", "first commit")
	_, _ = cmd.CombinedOutput()
	if ec := cmd.ProcessState.ExitCode(); ec != 0 {
		panic("error committing to a temporary repo for a unit test")
	}

	_ = os.Chdir(wd)

	return tmpRepoPath
}
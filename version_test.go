package main

import (
	"fmt"
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

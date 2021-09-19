package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestVersionSubCmd(tester *testing.T) {
	if os.Getenv(SubCmdFlags) != "" {
		args := strings.Split(os.Getenv(SubCmdFlags), " ")
		os.Args = append([]string{os.Args[0]}, args...)

		fmt.Printf("command: %v\n", os.Args)

		main()
	}

	var tests = []struct {
		name string
		wantCode int
		args []string
	}{
		{"versionSubCmd", 0, []string{"version"}},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			cmd := runMain(tester.Name(), test.args)

			cmdOut, cmdErr := cmd.CombinedOutput()

			// get exit code.
			got := cmd.ProcessState.ExitCode()

			if got != test.wantCode {
				t.Errorf("got %q, want %q", got, test.wantCode)
			}

			if cmdOut != nil {
				fmt.Printf("\nBEGIN sub-command stdout:\n%v", string(cmdOut))
				fmt.Print("END sub-command\n")
			}

			if cmdErr != nil {
				fmt.Printf("\nBEGIN sub-command stderr:\n%v", cmdErr.Error())
				fmt.Print("END sub-command\n")
			}
		})
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestVersionSubCmd(tester *testing.T) {
	var tests = []struct {
		name     string
		wantCode int
		args     []string
		version  string
	}{
		{"versionSubCmd", 0, []string{"version"}, "1.0.0"},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			tmpRepo := setupARepository("repo-01")
			test.args = append(test.args, tmpRepo)

			cmd := getTestBinCmd("TestAppMain", test.args)

			cmdOut, cmdErr := cmd.CombinedOutput()

			got := cmd.ProcessState.ExitCode()

			if cmdOut != nil {
				fmt.Printf("\nBEGIN sub-command stdout:\n%v", string(cmdOut))
				fmt.Print("END sub-command\n")
			}

			if cmdErr != nil {
				fmt.Printf("\nBEGIN sub-command stderr:\n%v", cmdErr.Error())
				fmt.Print("END sub-command\n")
			}

			var bv buildVersion

			bvFile := tmpRepo + PS + buildVersionFile
			bvData, _ := os.ReadFile(bvFile)
			if e := json.Unmarshal(bvData, &bv); e != nil {
				t.Errorf("test failed trying to decode %v: %v", bvFile, e.Error())
			}

			if got != test.wantCode {
				t.Errorf("unexpected error on exit. want %q, got %q", test.wantCode, got)
			}
			if bv.CurrentVersion != test.version {
				t.Errorf("unexpected version got %q, want %q", bv.CurrentVersion, test.version)
			}
		})
	}
}

func setupARepository(repoName string) string {
	tmpRepoPath := testTmp + PS + repoName

	srcRepo := "." + PS + fixturesDir + PS + repoName + ".bundle"
	cmd := exec.Command("git", "clone", "-b", "main", srcRepo, tmpRepoPath)
	_, _ = cmd.CombinedOutput()
	if ec := cmd.ProcessState.ExitCode(); ec != 0 {
		log.Panicf("error un-bundling %q to a temporary repo %q for a unit test", srcRepo, tmpRepoPath)
	}

	return tmpRepoPath
}

package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"testing"
)

func TestVersionSubCmd(tester *testing.T) {
	var tests = []struct {
		name        string
		wantCode    int
		args        []string
		version     string
		hash        string
		repo        string
		nextVersion string
	}{
		{"withGitTag", 0, []string{"version"}, "1.0.0", "a7f111c23f68c3b7fb8fefb7b8cd57cd04879f2a", "repo-01", "1.0.1"},
		{"withoutGitTag", 0, []string{"version"}, "HEAD", "bb062f4b4c8df46b08f824c185641747b128ebf8", "repo-02", "0.1.0"},
		{"withReleaseCommitMsg", 0, []string{"version"}, "HEAD", "f34f896ab7b88e49b4b5e45ac0d6385fcf3549c3", "repo-03", "0.2.0"},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			tmpRepo := setupARepository(test.repo)
			test.args = append(test.args, tmpRepo)

			cmd := getTestBinCmd(test.args)

			cmdOut, cmdErr := cmd.CombinedOutput()

			got := cmd.ProcessState.ExitCode()

			// Debug
			showCmdOutput(cmdOut, cmdErr)

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
			if bv.CommitHash != test.hash {
				t.Errorf("unexpected commit hash got %q, want %q", bv.CommitHash, test.hash)
			}
			if bv.NextVersion != test.nextVersion {
				t.Errorf("unexpected next version got %q, want %q", bv.NextVersion, test.nextVersion)
			}
		})
	}
}

func TestVersionSubCmdInvalidInput(tester *testing.T) {
	var tests = []struct {
		name     string
		wantCode int
		args     []string
	}{
		{"notARepo", 1, []string{"version", "repo-00"}},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			cmd := getTestBinCmd(test.args)

			cmdOut, cmdErr := cmd.CombinedOutput()

			got := cmd.ProcessState.ExitCode()

			// Debug
			showCmdOutput(cmdOut, cmdErr)

			if got != test.wantCode {
				t.Errorf("unexpected error on exit. want %q, got %q", test.wantCode, got)
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

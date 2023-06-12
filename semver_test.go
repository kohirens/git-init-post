package main

import (
	"encoding/json"
	help "github.com/kohirens/stdlib/test"
	"path/filepath"

	"os"
	"testing"
)

const (
	ps = string(os.PathSeparator)
)

var (
	tmpDir, _  = filepath.Abs("tmp")
	fixtureDir = "testdata"
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
		{"withGitTag", 0, []string{"semver"}, "1.0.0", "a7f111c23f68c3b7fb8fefb7b8cd57cd04879f2a", "repo-01", "1.0.1"},
		{"withoutGitTag", 0, []string{"semver"}, "HEAD", "bb062f4b4c8df46b08f824c185641747b128ebf8", "repo-02", "0.1.0"},
		{"withReleaseCommitMsg", 0, []string{"semver"}, "HEAD", "f34f896ab7b88e49b4b5e45ac0d6385fcf3549c3", "repo-03", "0.2.0"},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			tmpRepo := help.SetupARepository(test.repo, tmpDir, fixtureDir, ps)
			bvFile := tmpRepo + PS + "build-version.json"
			test.args = append(test.args, "-repo "+tmpRepo, "-save "+bvFile)
			cmd := getTestBinCmd(test.args)

			cmdOut, cmdErr := cmd.CombinedOutput()

			got := cmd.ProcessState.ExitCode()

			// Debug
			showCmdOutput(cmdOut, cmdErr)

			var bv buildVersion

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
		{"notARepo", 1, []string{"semver", "-repo repo-00"}},
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

func TestGetSemverInfo(tester *testing.T) {
	var tests = []struct {
		name        string
		repo        string
		version     string
		hash        string
		nextVersion string
		shouldErr   bool
	}{
		{"notARepo", "repo-dne", "", "", "", true},
		{"withGitTag", "repo-01", "1.0.0", "a7f111c23f68c3b7fb8fefb7b8cd57cd04879f2a", "1.0.1", false},
		{"withoutGitTag", "repo-02", "HEAD", "bb062f4b4c8df46b08f824c185641747b128ebf8", "0.1.0", false},
		{"withReleaseCommitMsg", "repo-03", "HEAD", "f34f896ab7b88e49b4b5e45ac0d6385fcf3549c3", "0.2.0", false},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			tmpRepo := ""
			if test.repo != "repo-dne" {
				tmpRepo = help.SetupARepository(test.repo, tmpDir, fixtureDir, ps)
			}

			bvData, err := GetSemverInfo(tmpRepo)

			if (err != nil) != test.shouldErr {
				t.Errorf("error %v", err.Error())
			}

			if test.repo == "repo-dne" {
				return
			}

			if bvData.CurrentVersion != test.version {
				t.Errorf("unexpected version got %q, want %q", bvData.CurrentVersion, test.version)
			}
			if bvData.CommitHash != test.hash {
				t.Errorf("unexpected commit hash got %q, want %q", bvData.CommitHash, test.hash)
			}
			if bvData.NextVersion != test.nextVersion {
				t.Errorf("unexpected next version got %q, want %q", bvData.NextVersion, test.nextVersion)
			}
		})
	}
}

func TestSettingRelVersion(tester *testing.T) {
	var tests = []struct {
		name   string
		want   string
		bundle string
	}{
		{"withExtra", "1.0.0-rc", "repo-05"},
		{"invalidExtra", "1.0.0", "repo-06"},
		{"relOverBreakingChange", "1.0.0-rc3", "repo-07"},
		{"noCommitToIncrementNextVersion", "1.0.0-rc1", "repo-08"},
		{"nextAfterNonStandardTag", "1.0.0-rc1", "repo-09"},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			repoPath := help.SetupARepository(test.bundle, tmpDir, fixtureDir, ps)

			got, err := GetSemverInfo(repoPath)

			if err != nil {
				t.Errorf("unexpected error %q", err.Error())
			}

			if got.NextVersion != test.want {
				t.Errorf("unexpected next version, want %q, got %q", test.want, got.NextVersion)
			}
		})
	}
}

func TestScrubNumber(tester *testing.T) {
	var tests = []struct {
		name    string
		want    string
		fixture string
	}{
		{"scrubZero", "0", "0"},
		{"scrubFour", "4", "4"},
		{"scrubReleaseCandidate", "22", "22-rc1"},
		{"scrubTwoWhatever", "2", "2.whatever"},
		{"scrubZero", "", ""},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			got := scrubNumber(test.fixture)

			if got != test.want {
				t.Errorf("unexpected next version, want %q, got %q", test.want, got)
			}
		})
	}
}

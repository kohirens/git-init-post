package main

import (
	"log"
	"os/exec"
	"testing"
)

func TestHasChangesToTag(tester *testing.T) {
	var tests = []struct {
		name   string
		want   bool
		tag    string
		bundle string
		repo   string
	}{
		{"withChangeTag", true, "1.0.0", "repo-01", "hasTags-01"},
		{"withoutTags", false, "HEAD", "repo-02", "hasTags-02"},
		{"withReleaseTag", true, "HEAD", "repo-03", "hasTags-03"},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			tmpRepo := setupARepository(test.repo, test.bundle)

			got := hasTags(tmpRepo, test.tag)

			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func setupARepository(repoName, bundleName string) string {
	tmpRepoPath := testTmp + PS + repoName

	srcRepo := "." + PS + fixturesDir + PS + bundleName + ".bundle"
	cmd := exec.Command("git", "clone", "-b", "main", srcRepo, tmpRepoPath)
	_, _ = cmd.CombinedOutput()
	if ec := cmd.ProcessState.ExitCode(); ec != 0 {
		log.Panicf("error un-bundling %q to a temporary repo %q for a unit test", srcRepo, tmpRepoPath)
	}

	return tmpRepoPath
}

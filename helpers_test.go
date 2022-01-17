package main

import (
	"log"
	"os"
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
		{"withChangeTag", true, "1.0.0..HEAD", "repo-01", "hasUnreleasedCommitsWithTags-01"},
		{"withoutTags", false, "HEAD", "repo-02", "hasUnreleasedCommitsWithTags-02"},
		{"withReleaseTag", true, "HEAD", "repo-03", "hasUnreleasedCommitsWithTags-03"},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			tmpRepo := setupARepository(test.repo, test.bundle)

			got := hasUnreleasedCommitsWithTags(tmpRepo, test.tag)

			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

func setupARepository(repoName, bundleName string) string {
	tmpRepoPath := testTmp + PS + repoName

	fileInfo, err1 := os.Stat(tmpRepoPath)
	if (err1 == nil && fileInfo.IsDir()) || os.IsExist(err1) {
		return tmpRepoPath
	}

	srcRepo := "." + PS + fixturesDir + PS + bundleName + ".bundle"
	cmd := exec.Command("git", "clone", "-b", "main", srcRepo, tmpRepoPath)
	_, _ = cmd.CombinedOutput()
	if ec := cmd.ProcessState.ExitCode(); ec != 0 {
		log.Panicf("error un-bundling %q to a temporary repo %q for a unit test", srcRepo, tmpRepoPath)
	}

	return tmpRepoPath
}

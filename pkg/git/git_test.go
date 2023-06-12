// Contains methods that are helpful for testing.
package git

import (
	help "github.com/kohirens/stdlib/test"
	"os"
	"path/filepath"
	"testing"
)

const (
	ps = string(os.PathSeparator)
)

var (
	tmpDir, _  = filepath.Abs("tmp")
	fixtureDir = "testdata"
)

func TestHasChangesToTag(tester *testing.T) {
	var tests = []struct {
		name   string
		want   bool
		tag    string
		bundle string
		repo   string
	}{
		{"withChangeTag", true, "1.0.0..HEAD", "repo-01", "HasUnreleasedCommitsWithTags-01"},
		{"withoutTags", false, "HEAD", "repo-02", "HasUnreleasedCommitsWithTags-02"},
		{"withReleaseTag", true, "HEAD", "repo-03", "HasUnreleasedCommitsWithTags-03"},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			repo := help.SetupARepository(test.bundle, tmpDir, fixtureDir, ps)
			//tmpRepo := setupARepository(test.repo, test.bundle)

			got := HasUnreleasedCommitsWithTags(repo, test.tag)

			if got != test.want {
				t.Errorf("want %v, got %v", test.want, got)
			}
		})
	}
}

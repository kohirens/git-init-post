package main

import "testing"

func TestIsTaggable(tester *testing.T) {
	var tests = []struct {
		name   string
		code   int
		args   []string
		bundle string
		repo   string
		want   string
	}{
		{"manyTagsInCommit", 0, []string{"taggable"}, "repo-01", "taggable-01", "true"},
		{"noTagInCommit", 0, []string{"taggable"}, "repo-02", "taggable-02", "false"},
		{"relIneCommit", 0, []string{"taggable"}, "repo-03", "taggable-03", "true"},
		{"noCurrentTag", 0, []string{"taggable", "-commitRange", "HEAD"}, "repo-04", "taggable-04", "true"},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			tmpRepo := setupARepository(test.repo, test.bundle)
			test.args = append(test.args, "-repo "+tmpRepo)

			cmd := getTestBinCmd(test.args)

			got, cmdErr := cmd.CombinedOutput()

			code := cmd.ProcessState.ExitCode()

			// Debug
			showCmdOutput(got, cmdErr)

			if code != test.code {
				t.Errorf("unexpected error on exit. want %q, got %q", test.code, code)
			}
			if string(got) != test.want {
				t.Errorf("unexpected error on exit. want %q, got %q", test.want, got)
			}
		})
	}
}

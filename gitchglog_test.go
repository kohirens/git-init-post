package main

import (
	"github.com/kohirens/stdlib"
	"testing"
)

func TestAddMissingChgLogConfig(tester *testing.T) {
	var testCases = []struct {
		name string
		conf string
		repo string
	}{
		{"noConf", testTmp + "/.chglog/config.yml", "git@github.com/example/app.git"},
		{"noConfHttps", testTmp + "/https/.chglog/config.yml", "https://github.com/example/app"},
	}

	for _, tc := range testCases {
		tester.Run(tc.name, func(t *testing.T) {
			err := addMissingChgLogConfig(tc.conf, tc.repo)

			if err != nil {
				t.Errorf("unexpected error: %s", err.Error())
			}

			if !stdlib.PathExist(tc.conf) {
				t.Errorf("expected file %v does not exists", tc.conf)
			}
		})
	}
}

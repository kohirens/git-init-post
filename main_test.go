package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"testing"
)

const (
	fixturesDir = "testdata"
	testTmp     = "tmp"
	// SubCmdFlags space separated list of command line flags.
	SubCmdFlags = "RECURSIVE_TEST_FLAGS"
)

func TestMain(m *testing.M) {
	// ONLY run this code when this test binary has been called directly.
	if os.Getenv(SubCmdFlags) != "" {
		runAppMain()
	}

	// delete all tmp files before running all test, but leave them afterward for manual inspection.
	_ = os.RemoveAll(testTmp)
	// Set up a temporary dir for generate files
	_ = os.Mkdir(testTmp, dirMode) // set up a temporary dir for generate files

	// Run all tests
	exitCode := m.Run()
	// Clean up
	os.Exit(exitCode)
}

func xTestCallingMain(tester *testing.T) {
	var tests = []struct {
		name     string
		wantCode int
		args     []string
	}{
		{"versionFlag", 0, []string{"-v"}},
		{"helpFlag", 0, []string{"-h"}},
	}

	for _, test := range tests {
		tester.Run(test.name, func(t *testing.T) {
			cmd := getTestBinCmd("TestAppMain", test.args)

			out, sce := cmd.CombinedOutput()

			// get exit code.
			got := cmd.ProcessState.ExitCode()

			if got != test.wantCode {
				t.Errorf("got %q, want %q", got, test.wantCode)
			}

			if sce != nil {
				fmt.Printf("\nBEGIN sub-command stdout:\n%v\n\n", string(out))
				fmt.Printf("stderr:\n%v\n", sce.Error())
				fmt.Print("\nEND sub-command\n\n")
			}
		})
	}
}

// Used for running the application's main function from other test.
func runAppMain() {
	args := strings.Split(os.Getenv(SubCmdFlags), " ")
	os.Args = append([]string{os.Args[0]}, args...)

	// Debug stmt
	//fmt.Printf("\nos args = %v\n", os.Args)

	main()
}

// getTestBinCmd will run the binary build for these test; if you have a `TestMain`, it will be run automatically.
func getTestBinCmd(testFunc string, args []string) *exec.Cmd {
	// call the generated test binary directly
	// Have it the function runAppMain.
	cmd := exec.Command(os.Args[0], "-args", strings.Join(args, " "))
	// Run in the context of the source directory.
	_, filename, _, _ := runtime.Caller(0)
	cmd.Dir = path.Dir(filename)
	// Pass an environment variable for us to know when we have called this test binary directly.
	subEnvVar := SubCmdFlags + "=" + strings.Join(args, " ")
	cmd.Env = append(os.Environ(), subEnvVar)

	return cmd
}

// quiet Prints output to the OS null space.
func quiet() func() {
	null, _ := os.Open(os.DevNull)
	sOut := os.Stdout
	sErr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	log.SetOutput(null)
	return func() {
		defer null.Close()
		os.Stdout = sOut
		os.Stderr = sErr
		log.SetOutput(os.Stderr)
	}
}

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
	// subCmdFlags space separated list of command line flags.
	subCmdFlags = "RECURSIVE_TEST_FLAGS"
)

func TestMain(m *testing.M) {
	// Only runs when this environment variable is set.
	if os.Getenv(subCmdFlags) != "" {
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
			cmd := getTestBinCmd(test.args)

			_, cmdErr := cmd.CombinedOutput()

			// get exit code.
			got := cmd.ProcessState.ExitCode()

			if got != test.wantCode {
				t.Errorf("got %q, want %q", got, test.wantCode)
			}

			if cmdErr != nil {
				fmt.Printf("\nBEGIN sub-command stderr:\n%v", cmdErr.Error())
				fmt.Print("END sub-command\n")
			}
		})
	}
}

// Used for running the application's main function from other test.
func runAppMain() {
	args := strings.Split(os.Getenv(subCmdFlags), " ")
	os.Args = append([]string{os.Args[0]}, args...)

	// Debug stmt
	//fmt.Printf("\nos args = %v\n", os.Args)

	main()
}

// getTestBinCmd return a command to run your app (test) binary directly; `TestMain`, will be run automatically.
func getTestBinCmd(args []string) *exec.Cmd {
	// call the generated test binary directly
	// Have it the function runAppMain.
	cmd := exec.Command(os.Args[0], "-args", strings.Join(args, " "))
	// Run in the context of the source directory.
	_, filename, _, _ := runtime.Caller(0)
	cmd.Dir = path.Dir(filename)
	// Set an environment variable
	// 1. Only exist for the life of the test that calls this function.
	// 2. Passes arguments/flag to your app
	// 3. Lets TestMain know when to run the main function.
	subEnvVar := subCmdFlags + "=" + strings.Join(args, " ")
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

func showCmdOutput(cmdOut []byte, cmdErr error) {
	if !testing.Verbose() {
		return
	}

	if cmdOut != nil {
		fmt.Printf("\nBEGIN sub-command out:\n%v", string(cmdOut))
		fmt.Print("END sub-command\n")
	}

	if cmdErr != nil {
		fmt.Printf("\nBEGIN sub-command stderr:\n%v", cmdErr.Error())
		fmt.Print("END sub-command\n")
	}
}

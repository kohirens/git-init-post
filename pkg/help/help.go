package help

import (
	"os"
	"os/exec"
)

type ByteBuf struct {
	Buf []byte
}

func (bb *ByteBuf) Write(b []byte) (n int, err error) {
	if bb.Buf == nil {
		bb.Buf = []byte{}
	}

	bb.Buf = append(bb.Buf, b...)

	return len(b), err
}

// RunRepoCmd run a command against the repository.
func RunRepoCmd(repoPath string, args ...string) (cmdOut []byte, cmdErr error, exitCode int, err error) {
	cmd := exec.Command("git", args...)
	cmd.Env = os.Environ()
	cmd.Dir = repoPath
	cmdOut, cmdErr = cmd.CombinedOutput()
	exitCode = cmd.ProcessState.ExitCode()
	return
}

package core

import (
	"os"
	"os/exec"
	"syscall"
)

func CreateIsolatedProc() {
	cmd := exec.Command("ls", "-l")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags:   syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET | syscall.CLONE_NEWIPC | syscall.CLONE_NEWUSER,
		Unshareflags: syscall.CLONE_NEWNS,
	}

	must(cmd.Run())
}

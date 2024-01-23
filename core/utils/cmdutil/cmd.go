package cmdutil

import (
	"bytes"
	"errors"
	"os/exec"
	"time"
)

func ExecBashCmd(cmd string) ([]byte, error) {
	command := exec.Command("bash", "-c", cmd)
	return command.CombinedOutput()
}

func ExecShCmd(cmd string) ([]byte, error) {
	command := exec.Command("sh", "-c", cmd)
	return command.CombinedOutput()
}

func ExecCmd(cmd string, args ...string) ([]byte, error) {
	command := exec.Command(cmd, args...)
	return command.CombinedOutput()
}

func ExecCmdString(cmd string, args []string) (string, error) {
	output, err := exec.Command(cmd, args...).CombinedOutput()
	return string(output), err
}

func ExecCmdWithReadWrite(cmd string, args []string, in *bytes.Buffer, out *bytes.Buffer) error {
	command := exec.Command(cmd, args...)
	command.Stdin = in
	command.Stdout = out
	command.Stderr = out
	err := command.Run()
	return err
}

func ExecCmdWithReadWriteAndTimeout(cmd string, args []string, timeout time.Duration, in *bytes.Buffer, out *bytes.Buffer) error {
	command := exec.Command(cmd, args...)
	command.Stdin = in
	command.Stdout = out
	command.Stderr = out
	err := command.Run()
	return err
}

func CheckExitErr(err error) bool {
	var exErr *exec.ExitError
	return !errors.As(err, &exErr)
}

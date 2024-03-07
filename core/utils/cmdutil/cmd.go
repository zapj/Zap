package cmdutil

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"time"
)

func ExecBashCmd(cmd string) ([]byte, error) {
	command := exec.Command("bash", "-c", cmd)
	return command.CombinedOutput()
}

func ExecBashCmdWithContext(ctx context.Context, cmd string) ([]byte, error) {
	command := exec.CommandContext(ctx, "bash", "-c", cmd)
	return command.CombinedOutput()
}

func ExecShCmd(cmd string) ([]byte, error) {
	command := exec.Command("sh", "-c", cmd)
	return command.CombinedOutput()
}

func ExecShCmdWithContext(ctx context.Context, cmd string) ([]byte, error) {
	command := exec.CommandContext(ctx, "sh", "-c", cmd)
	return command.CombinedOutput()
}

func ExecCmd(cmd string, args ...string) ([]byte, error) {
	command := exec.Command(cmd, args...)
	return command.CombinedOutput()
}

func ExecCmdWithContext(ctx context.Context, cmd string, args ...string) ([]byte, error) {
	command := exec.CommandContext(ctx, cmd, args...)
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

func CallContext(prog string, stdin []byte, timeout int, args ...string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, prog, args...)
	// cmd.SysProcAttr = &syscall.SysProcAttr{Cloneflags: syscall.CLONE_NEWUTS}
	// cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 1000, Gid: 1000}
	in := bytes.NewReader(stdin)
	out := &bytes.Buffer{}
	errs := &bytes.Buffer{}

	cmd.Stdin, cmd.Stdout, cmd.Stderr = in, out, errs

	if err := cmd.Run(); err != nil {
		if len(errs.Bytes()) > 0 {
			return nil, fmt.Errorf("error running %s (%s):\n %v", cmd.Args, err, errs.String())
		}
		return nil, err
	}

	return out.Bytes(), nil
}

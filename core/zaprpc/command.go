package zaprpc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/zapj/zap/core/global"
)

type Command struct {
}

func (c *Command) Call(args Args, resp *Response) error {
	cmd := args.GetString("cmd")
	params := args.GetStringArray("params")
	cc := exec.Command(cmd, params...)
	out, err := cc.CombinedOutput()
	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			resp.Code = 404
			resp.Msg = fmt.Sprintf("%s 命令不存在", cmd)
			return nil
		}
		resp.Code = 500
		resp.Msg = err.Error()
		global.LOG.Infof("%v", err)
		return err
	}
	resp.Code = 0
	resp.Data = string(out)

	return nil
}

func (c *Command) CallContext(args Args, resp *Response) error {
	cmd := args.GetString("cmd")
	params := args.GetStringArray("params")
	stdin := args.GetString("stdin")
	timeout := args.GetInt("timeout")
	if timeout == 0 {
		timeout = 3
	}
	outBytes, err := CallContext(cmd, []byte(stdin), timeout, params...)

	if err != nil {
		if errors.Is(err, exec.ErrNotFound) {
			resp.Code = 404
			resp.Msg = fmt.Sprintf("%s 命令不存在", cmd)
			return nil
		}
		resp.Code = 500
		resp.Msg = err.Error()
		global.LOG.Infof("%v", err)
		return err
	}
	resp.Code = 0
	resp.Data = string(outBytes)

	return nil
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

package zaprpc

import (
	"errors"
	"fmt"
	"os/exec"

	"github.com/zapj/zap/assist/global"
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

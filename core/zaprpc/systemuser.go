package zaprpc

import (
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/sysapi/user"
)

type SystemUser struct {
}

func (c *SystemUser) CreateUser(args Args, resp *Response) error {
	userEntry := user.UserEntry{
		Username: args.GetString("username"),
		Password: args.GetString("password"),
		Uid:      args.GetString("uid"),
		Gid:      args.GetString("gid"),
		Gecos:    "",
		Home:     args.GetString("home"),
		Shell:    "/usr/sbin/nologin",
	}

	err := user.CreateUser(userEntry)

	if err != nil {
		resp.Code = 500
		resp.Msg = err.Error()
		global.LOG.Infof("%v", err)
		return err
	}
	resp.Code = 0
	resp.Msg = "创建成功"
	return nil
}

func (c *SystemUser) ReadShadow(args Args, resp *Response) error {
	users, err := user.ReadSystemShadow()
	if err != nil {
		resp.Msg = err.Error()
		return err
	}
	// global.LOG.Info(users, err)
	resp.SetDataJson(users)
	resp.Msg = ""
	return nil
}

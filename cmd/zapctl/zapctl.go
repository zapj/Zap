package main

import (
	"github.com/zapj/zap/cmd/zapctl/cmd"
	"github.com/zapj/zap/core/base"
	"github.com/zapj/zap/core/conf"
	"github.com/zapj/zap/core/global"
)

var Version = "0.0.0"
var BuildDate = "0000-00-00_00:00:00"

func main() {
	global.ZAP_INFO = &base.ZapInfo{Version: Version, BuildDate: BuildDate, BuildShortDate: BuildDate[0:8]}
	conf.InitEnv()
	// conf.LogInit("zapc.log")
	conf.DbInit()
	cmd.Execute()
}

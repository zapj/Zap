package main

import (
	"github.com/zapj/zap/cmd/zapctl/cmd"
	"github.com/zapj/zap/core"
	"github.com/zapj/zap/core/conf"
	"github.com/zapj/zap/core/global"
)

var Version = "0.0.0"
var BuildDate = "0000-00-00_00:00:00"

func main() {
	global.ZAP_INFO = &core.ZapInfo{Version: Version, BuildDate: BuildDate}
	conf.InitEnv()
	conf.LogInit()
	conf.DbInit()
	cmd.Execute()
}

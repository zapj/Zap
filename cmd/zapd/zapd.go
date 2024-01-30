package main

import (
	"github.com/zapj/zap/cmd/zapd/cmd"
	"github.com/zapj/zap/core"
	"github.com/zapj/zap/core/global"
)

var Version = "0.0.0"
var BuildDate = "0000-00-00_00:00:00"

func main() {
	global.ZAP_INFO = &core.ZapInfo{Version: Version, BuildDate: BuildDate, BuildShortDate: BuildDate[0:8]}
	cmd.Execute()
}

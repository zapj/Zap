package main

import (
	"github.com/zapj/zap/cmd/zapctl/cmd"
	"github.com/zapj/zap/core/conf"
)

func main() {
	conf.LogInit()
	conf.InitEnv()
	conf.DbInit()
	cmd.Execute()
}

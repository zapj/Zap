package main

import (
	"github.com/sevlyar/go-daemon"
	"github.com/zapj/zap/core/conf"
	"github.com/zapj/zap/core/global"

	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {
	conf.LogInit()
	// using standard library "flag" package
	flag.Bool("daemon", false, "daemon mode")
	flag.Bool("debug", false, "debug mode")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	Run_Daemon := viper.GetBool("daemon")
	// Debug_Mode := viper.GetBool("debug")
	if Run_Daemon {
		cntxt := &daemon.Context{
			PidFileName: "zap.pid",
			PidFilePerm: 0644,
			LogFileName: "zap.log",
			LogFilePerm: 0640,
			WorkDir:     "./",
			Umask:       022,
			Args:        []string{},
		}

		d, err := cntxt.Reborn()
		if err != nil {
			global.LOG.Fatal("Unable to run: ", err)
		}
		if d != nil {
			return
		}
		defer cntxt.Release()

		global.LOG.Print("- - - - - - - - - - - - - - -")
		global.LOG.Print("daemon started")
	}

	conf.InitEnv()
	conf.DbInit()

	conf.RpcInit()
	global.LOG.Info("Server exiting")

}

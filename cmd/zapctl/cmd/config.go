package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/magiconair/properties"
	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/pathutil"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置文件",
	// Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := ""
		if len(args) >= 1 {
			action = args[0]
		}
		switch action {
		case "set":
			if len(args) == 3 {
				setConfig(args[1], args[2])
			} else {
				fmt.Println("设置失败")
			}
		case "get":
			if len(args) == 2 {
				getConfig(args[1])
			}
		default:
			defaultPrintAll()
		}
	},
}

func defaultPrintAll() {
	fmt.Println("Config Path : ", pathutil.GetPath("conf/zap.conf"))
	fmt.Println("")
	fmt.Println(global.CONFIG.String())
}

func setConfig(name, value string) {
	global.CONFIG.SetValue(name, value)
	global.CONFIG.WriteSeparator = "="
	f, err := os.Create(pathutil.GetPath("conf/zap.conf"))
	if err != nil {
		slog.Info("配置文件保存失败: %v \n", err)
	}
	defer f.Close()
	global.CONFIG.Write(f, properties.UTF8)
}

func getConfig(name string) {
	if v, ok := global.CONFIG.Get(name); ok {
		fmt.Println(name, " : ", v)
	}
}

package cmd

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/pathutil"
	"gopkg.in/ini.v1"
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
			if len(args) == 4 {
				setSectionConfig(args[1], args[2], args[3])
			} else if len(args) == 3 {
				setConfig(args[1], args[2])
			} else {
				fmt.Println("设置失败")
			}
		case "get":
			if len(args) == 3 {
				getSectionConfig(args[1], args[2])
			} else if len(args) == 2 {
				getConfig(args[1])
			}
		default:
			defaultPrintAll()
		}
	},
}

func defaultPrintAll() {
	fmt.Println("Config Path : ", pathutil.GetPath("conf/zap.ini"))
	fmt.Println("")
	global.CONFIG.WriteTo(os.Stdout)
}

func setSectionConfig(section, name, value string) {
	global.CONFIG.Section(section).Key(name).SetValue(value)
	f, err := os.Create(pathutil.GetPath("conf/zap.ini"))
	if err != nil {
		fmt.Printf("配置文件保存失败: %v \n", err)
	}
	defer f.Close()
	global.CONFIG.WriteTo(f)
}

func setConfig(name, value string) {
	global.CONFIG.Section(ini.DefaultSection).Key(name).SetValue(value)
	f, err := os.Create(pathutil.GetPath("conf/zap.ini"))
	if err != nil {
		slog.Info("配置文件保存失败: %v \n", err)
	}
	defer f.Close()
	global.CONFIG.WriteTo(f)
}

func getSectionConfig(section, name string) {

	fmt.Println(name, " : ", global.CONFIG.Section(section).Key(name).String())
}

func getConfig(name string) {

	fmt.Println(name, " : ", global.CONFIG.Section(ini.DefaultSection).Key(name).String())
}

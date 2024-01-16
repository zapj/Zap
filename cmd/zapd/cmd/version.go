package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	BuildDate string = "20240116"
	Version   string = "0.0.1"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "ZAP控制面板版本号",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("zap_deamon Build %s-%s Build Date %s ,Version %s \n\n", runtime.GOOS, runtime.GOARCH, BuildDate, Version)
	},
}

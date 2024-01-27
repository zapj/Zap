package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/global"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "zapctl version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ZAP Version %s-%s %s BuildDate: %s \n\n", runtime.GOOS, runtime.GOARCH, global.ZAP_INFO.Version, global.ZAP_INFO.BuildDate)
	},
}

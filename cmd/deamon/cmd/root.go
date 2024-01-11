package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zap_deamon",
	Short: "zap_deamon命令行管理工具",
	Run: func(cmd *cobra.Command, args []string) {
		versionCmd.Run(cmd, args)
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

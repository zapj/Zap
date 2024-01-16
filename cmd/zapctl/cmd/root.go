package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "zapctl",
	Short: "zap命令行管理工具",
	Run: func(cmd *cobra.Command, args []string) {
		versionCmd.Run(cmd, args)
		cmd.Help()
	},
}

// var cfgFile string

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.zap_config)")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

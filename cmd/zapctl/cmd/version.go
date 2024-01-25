package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "zapctl version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("ZAP Version %s %s \n\n", runtime.GOARCH, core.Version)
	},
}

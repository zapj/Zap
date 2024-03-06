package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/global"
)

func init() {
	rootCmd.AddCommand(getCmd)
}

var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "zapctl get [zap...]",
	Example: "zapctl get [zap...]",
	Args:    cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "zap":
			fmt.Print(global.ZAP_BASE_DIR)
		}
	},
}

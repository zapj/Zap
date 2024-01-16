package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(keyCmd)
}

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "证书工具",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

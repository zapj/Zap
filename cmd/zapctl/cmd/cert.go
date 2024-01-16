package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(certCmd)
}

var certCmd = &cobra.Command{
	Use:   "cert",
	Short: "证书工具",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

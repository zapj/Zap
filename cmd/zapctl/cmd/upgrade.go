package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/zapi"
)

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade zap",
	Run: func(cmd *cobra.Command, args []string) {
		resp := zapi.Client.Upgrade()

		if resp.HasError() {
			fmt.Printf("Error Message : %s\n", resp.ErrorMessage())
			return
		}
		if resp.Status == 200 {
			r := resp.ZapMap()
			fmt.Printf("code:%s , msg : %s\n", r.GetString("code"), r.GetString("msg"))
		}

	},
}

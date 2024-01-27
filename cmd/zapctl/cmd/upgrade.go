package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/task"
)

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade zap",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := task.Get("/upgrade")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("提交成功，系统更新后自动重启zapd")
	},
}

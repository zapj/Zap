package cmd

import (
	"fmt"

	"github.com/go-zoox/fetch"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(upgradeCmd)
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "upgrade zap",
	Run: func(cmd *cobra.Command, args []string) {
		_, err := fetch.Get("https://127.0.0.1/upgrade", &fetch.Config{UnixDomainSocket: "/usr/local/zap/data/task", TLSInsecureSkipVerify: true})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("提交成功")
	},
}

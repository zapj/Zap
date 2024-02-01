package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/utils/str"
)

func init() {
	genCmd.PersistentFlags().BoolP("u", "u", false, "默认生成无符号的字符串")
	genCmd.PersistentFlags().IntP("length", "l", 20, "字符串长度")
	rootCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "generate rand string",
	Run: func(cmd *cobra.Command, args []string) {
		u, err := cmd.Flags().GetBool("u")
		if err != nil {
			u = true
		}
		length, err := cmd.Flags().GetInt("length")
		if err != nil {
			length = 20
		}
		if u {
			fmt.Println(str.RandUString(length))
		} else {
			fmt.Println(str.RandString(length))
		}
	},
}

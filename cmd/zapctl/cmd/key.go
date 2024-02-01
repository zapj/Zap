package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/ssl"
)

func init() {
	rootCmd.AddCommand(keyCmd)
}

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "generate private key",
	Run: func(cmd *cobra.Command, args []string) {
		pkey, _ := ssl.GeneratePrivateKey("RSA", 2048)
		fmt.Println(string(ssl.ExportPrivateKeyToMemory(pkey)))
	},
}

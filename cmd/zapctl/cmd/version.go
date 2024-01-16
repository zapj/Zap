package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	BuildDate    string
	BuildVersion string
	Version      string
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "zapctl版本号",
	Run: func(cmd *cobra.Command, args []string) {
		// time.Sleep(time.Second * 20)

		fmt.Printf("Build %s-%s-%s-%s ,Version %s \n\n", runtime.GOOS, runtime.GOARCH, BuildVersion, BuildDate, Version)
	},
}

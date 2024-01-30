package cmd

import (
	"os"
	"strconv"

	"github.com/sevlyar/go-daemon"
	"github.com/spf13/cobra"
	"github.com/zapj/zap/core/task"
	"github.com/zapj/zap/core/utils/zap"
)

var taskDaemon bool = false

func init() {
	taskCmd.PersistentFlags().BoolVarP(&taskDaemon, "daemon", "d", false, "start daemon")
	rootCmd.AddCommand(taskCmd)
}

var taskCmd = &cobra.Command{
	Use:   "master",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		if taskDaemon {
			cntxt := &daemon.Context{
				PidFileName: "zap_master.pid",
				PidFilePerm: 0644,
				LogFileName: "zap_master.log",
				LogFilePerm: 0640,
				WorkDir:     "./",
				Umask:       022,
				Args:        []string{},
			}

			d, err := cntxt.Reborn()
			if err != nil {
				cmd.PrintErrln("Unable to run: ", err)
			}
			if d != nil {
				return
			}
			defer cntxt.Release()
		} else {
			pid := os.Getpid()
			os.WriteFile(zap.GetPath("zap_master.pid"), []byte(strconv.Itoa(pid)), 0644)
		}
		task.StartTask()
	},
}

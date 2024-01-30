package task

import (
	"os"
	"os/exec"
	"time"

	"github.com/zapj/zap/core/global"
)

func StartServer() {
	retry := time.Second * 30
	for {
		serverCmd := exec.Command(os.Args[0], "server")
		err := serverCmd.Run()
		if err != nil {
			time.Sleep(retry)
		}
		if global.ZAP_MODE == "DEV" {
			serverCmd.Stdout = os.Stdout
			serverCmd.Stderr = os.Stderr
		}
		global.ServerPID = serverCmd.Process.Pid
		serverCmd.Wait()
	}
}

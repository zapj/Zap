package task

import (
	"os"
	"os/exec"
	"time"

	"github.com/zapj/zap/core/global"
)

func StartServer() {
	for {
		serverCmd := exec.Command(os.Args[0], "server")
		err := serverCmd.Run()
		if err != nil {
			time.Sleep(time.Second * 10)
		}
		global.ServerPID = serverCmd.Process.Pid
		serverCmd.Wait()
	}
}

package task

import (
	"log/slog"
	"os"
	"os/exec"
	"time"

	"github.com/zapj/zap/core/global"
)

func StartServer() {
	retry := time.Second * 30
	for {
		serverCmd := exec.Command(os.Args[0], "server")
		// serverCmd.SysProcAttr = &syscall.SysProcAttr{GidMappingsEnableSetgroups: true}
		// serverCmd.SysProcAttr.Credential = &syscall.Credential{Uid: 65534, Gid: 65534}
		err := serverCmd.Run()
		if err != nil {
			slog.Error("start server", "err", err)
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

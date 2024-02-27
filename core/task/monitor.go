package task

import (
	"context"
	"log/slog"
	"os"
	"os/exec"
	"time"

	"github.com/zapj/zap/core/global"
)

func StartServer(ctx context.Context) {
	retry := time.Second * 30
	for {
		serverCmd := exec.CommandContext(ctx, os.Args[0], "server")
		// serverCmd.SysProcAttr = &syscall.SysProcAttr{GidMappingsEnableSetgroups: true}
		// serverCmd.SysProcAttr.Credential = &syscall.Credential{Uid: 65534, Gid: 65534}
		if global.ZAP_MODE == "DEV" {
			serverCmd.Env = append(os.Environ(), "ZAP_MODE=DEV")
		}
		err := serverCmd.Start()
		if err != nil {
			slog.Error("start server", "err", err)
			time.Sleep(retry)
		}
		if global.ZAP_MODE == "DEV" {
			serverCmd.Stdout = os.Stdout
			serverCmd.Stderr = os.Stderr
		}

		global.ServerPID = serverCmd.Process.Pid
		err = serverCmd.Wait()
		if err != nil {
			slog.Error("server exit", "err", err)
			time.Sleep(retry)
		}

	}
}

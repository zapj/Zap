package daemon

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"github.com/zapj/zap/assist/global"
)

func InitProcess() {
	if syscall.Getppid() == 1 {
		if err := os.Chdir("./"); err != nil {
			panic(err)
		}
		syscall.Umask(0)
		return
	}
	global.LOG.Info().Msg("daemon start")
	fp, err := os.OpenFile("daemon.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = fp.Close()
	}()
	cmd := exec.Command(os.Args[0], os.Args[1:]...)
	cmd.Stdout = fp
	cmd.Stderr = fp
	cmd.Stdin = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	global.LOG.Info().Msgf(fmt.Sprintf("[PID] %d Start At %s\n", cmd.Process.Pid, time.Now().Format("2006-01-02 15:04:05")))

	os.Exit(0)
}

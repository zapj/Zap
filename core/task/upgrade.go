package task

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"syscall"
	"time"

	"github.com/go-zoox/fetch"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/cmdutil"
)

func UpgradeZapd() {
	GL.Store("upgrade", time.Now().Unix())
	defer func() {
		GL.Delete("upgrade")
	}()
	resp, err := fetch.Get(fmt.Sprintf("%s/dist/latest.txt?t=%d", ZAP_MIRROR_URL, time.Now().Unix()))
	if err != nil {
		slog.Info("check version", "err", err)
		return
	}
	zapdVer := string(resp.Body)
	if zapdVer == global.ZAP_INFO.Version {
		return
	}
	fileName := fmt.Sprintf("zap-v%s-%s-%s.tar.gz", zapdVer, runtime.GOOS, runtime.GOARCH)
	tmpFile := "/tmp/" + fileName
	_, err = fetch.Download(fmt.Sprintf("%s/dist/%s", ZAP_MIRROR_URL, fileName), tmpFile)
	if err != nil {
		slog.Info(fmt.Sprintf("download %s", fileName), "err", err)
		return
	}

	_, err = cmdutil.ExecBashCmd("tar zxf " + tmpFile + " -C /usr/local/")
	if err != nil {
		slog.Info("解压缩失败", "err", err)
		return
	}
	os.Remove(tmpFile)
	// cmdutil.ExecBashCmd("systemctl restart zapd.service")
	ServerProc, err := os.FindProcess(global.ServerPID)
	if err != nil {
		slog.Info("find process err", "err", err)
	}
	if err = ServerProc.Signal(syscall.SIGTERM); err != nil {
		slog.Info("signal err, restart zapd.service", "err", err)
		cmdutil.ExecBashCmd("systemctl restart zapd.service")
	}
}

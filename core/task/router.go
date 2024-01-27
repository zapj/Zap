package task

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-zoox/fetch"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/cmdutil"
)

func RegisterRouter(r *gin.RouterGroup) {

	r.GET("/upgrade", func(ctx *gin.Context) {
		go func() {
			resp, err := fetch.Get("https://mirrors.zap.cn/zap/dist/latest.txt?" + strconv.FormatInt(time.Now().Unix(), 10))
			if err != nil {
				global.LOG.Info(err)
				return
			}
			zapdVer := string(resp.Body)
			if zapdVer == global.ZAP_INFO.Version {
				return
			}
			fileName := fmt.Sprintf("zap-v%s-%s-%s.tar.gz", zapdVer, runtime.GOOS, runtime.GOARCH)
			tmpFile := "/tmp/" + fileName
			_, err = fetch.Download("https://mirrors.zap.cn/zap/dist/"+fileName, tmpFile)
			if err != nil {
				global.LOG.Info(err)
				return
			}

			_, err = cmdutil.ExecBashCmd("tar zxf " + tmpFile + " -C /usr/local/")
			if err != nil {
				slog.Info("解压缩失败", err)
			}
			os.Remove(tmpFile)
			// cmdutil.ExecBashCmd("systemctl restart zapd.service")
			ServerProc, err := os.FindProcess(global.ServerPID)
			if err != nil {
				slog.Info("find process err", err)
			}
			if err = ServerProc.Signal(syscall.SIGTERM); err != nil {
				slog.Info("signal err, restart zapd.service", err)
				cmdutil.ExecBashCmd("systemctl restart zapd.service")
			}

		}()
		ctx.JSON(200, gin.H{"code": 0, "msg": "提交成功，系统升级成功后将自动重启控制面板"})
	})
}

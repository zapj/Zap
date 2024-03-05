package task

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/ssl"
	"github.com/zapj/zap/core/utils/pathutil"
)

func StartTaskScheduler() {
	os.Remove(pathutil.GetPath(TASK_SERV_SOCK))
	crtFile, keyFile := pathutil.GetPath("data/server.crt"), pathutil.GetPath("data/server.key")
	if fileutils.IsFile(global.SERVER_CONF.CertFile) && fileutils.IsFile(global.SERVER_CONF.KeyFile) {
		crtFile = global.SERVER_CONF.CertFile
		keyFile = global.SERVER_CONF.KeyFile
	} else {
		ssl.AutoGenCertAndKey(crtFile, keyFile)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	if global.ZAP_MODE == "DEV" {
		gin.SetMode(gin.DebugMode)
	}
	RegisterRouter(&r.RouterGroup)
	srv := &http.Server{
		Handler: r,
	}

	go func() {
		unixListener, err := net.Listen("unix", TASK_SERV_SOCK)
		if err != nil {
			panic(err)
		}
		srv.ServeTLS(unixListener, crtFile, keyFile)
	}()
	go func() {
		StartServer(ctx)
	}()
	StartTaskQueue()
	<-ctx.Done()
	stop()

	slog.Info("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		slog.Info("Server forced to shutdown: %s \n", err)
	}

	slog.Info("Server exiting")
}

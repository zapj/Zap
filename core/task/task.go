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
	"github.com/zapj/zap/core/ssl"
	"github.com/zapj/zap/core/utils/zap"
)

func StartTask() {
	os.Remove(zap.GetPath("data/task"))
	crtFile, keyFile := zap.GetPath("data/server.crt"), zap.GetPath("data/server.key")
	ssl.AutoGenCertAndKey(crtFile, keyFile)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	RegisterRouter(&r.RouterGroup)
	srv := &http.Server{
		Handler: r,
	}

	go func() {
		unixListener, err := net.Listen("unix", zap.GetPath("data/task"))
		if err != nil {
			panic(err)
		}
		srv.ServeTLS(unixListener, crtFile, keyFile)
	}()
	go func() {
		StartServer()
	}()
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

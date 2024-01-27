package task

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/ssl"
)

func StartTask() {
	os.Remove("./data/task")
	crtFile, keyFile := "data/server.crt", "data/server.key"
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
		unixListener, err := net.Listen("unix", "./data/task")
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
	

	fmt.Println("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown: %s \n", err)
	}

	fmt.Println("Server exiting")
}

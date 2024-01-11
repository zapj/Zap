package main

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/webterm"
	"github.com/zapj/zap/web"
)

var (
	BuildDate string
	// GitCommit, GitBranch string
	BuildVersion string
	Version      string
)

func main() {
	fmt.Printf("Build %s-%s-%s-%s ,Version %s \n", runtime.GOOS, runtime.GOARCH, BuildVersion, BuildDate, Version)
	// syscall.Setegid(65534)
	// syscall.Seteuid(65534)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	assetsFS, _ := fs.Sub(web.ASSETS_FS, "static/assets")

	router.StaticFS("/assets", http.FS(assetsFS))
	router.GET("/", func(c *gin.Context) {
		c.Data(200, "text/html", web.INDEX_BYTES)
	})
	router.NoRoute(func(c *gin.Context) {
		accept := c.Request.Header.Get("Accept")
		flag := strings.Contains(accept, "text/html")
		if flag {
			c.Data(200, "text/html", web.INDEX_BYTES)
		}
	})
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Data(200, "image/x-icon", web.FAVICON_BYTES)
	})

	router.GET("/start", func(c *gin.Context) {
		go func() {
			cmd := exec.Command("./zap_cli", "version")
			cmd.SysProcAttr = &syscall.SysProcAttr{GidMappingsEnableSetgroups: true}
			cmd.SysProcAttr.Credential = &syscall.Credential{Uid: 65534, Gid: 65534, NoSetGroups: true}
			if err := cmd.Run(); err != nil {
				global.LOG.Error(err)
			}

		}()
	})

	v1 := router.Group("/api/v1")
	{
		api.RegisterRouter(v1)

	}
	router.GET("/ws", webterm.HandlerLocalWS)

	srv := &http.Server{
		Addr:    ":2828",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.LOG.Errorf("listen: %s\n", err)
		} else {
			global.LOG.Debugf("Listen Port : %s", ":2828")
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the assist it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}

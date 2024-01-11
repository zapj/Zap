package main

import (
	"context"
	"embed"
	"errors"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/webterm"
)

//go:embed static/assets/*
var ASSETS_FS embed.FS

//go:embed static/index.html
var INDEX_BYTES []byte

func main() {

	// syscall.Setegid(65534)
	// syscall.Seteuid(65534)
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	assetsFS, _ := fs.Sub(ASSETS_FS, "static/assets")
	router.StaticFS("/assets", http.FS(assetsFS))
	// router.StaticFS("/", http.FS(FS))
	router.GET("/", func(c *gin.Context) {
		c.Data(200, "text/html", INDEX_BYTES)
		// c.HTML(http.StatusOK, "index.html", gin.H{})
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
		v1.GET("/loadavg", func(c *gin.Context) {
			st, _ := load.Avg()
			c.String(200, "Loadavg  %s", st.String())
		})
		v1.GET("/mem", func(c *gin.Context) {
			v, _ := mem.VirtualMemory()
			c.String(200, v.String())
		})
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

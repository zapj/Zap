package main

import (
	"context"
	"errors"
	"github.com/zapj/zap/assist/conf"
	"github.com/zapj/zap/assist/global"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	conf.RouterInit(router)
	conf.LogInit()

	srv := &http.Server{
		Addr:    ":2728",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.LOG.Fatal().Msgf("listen: %s", err)
		}
	}()

	<-ctx.Done()

	stop()
	global.LOG.Info().Msg("shutting down gracefully, press Ctrl+C again to force")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.LOG.Fatal().Msgf("Server forced to shutdown: %s", err)
	}
	global.LOG.Info().Msg("Server exiting")

}

package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/assist/global"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		time.Sleep(10 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	srv := &http.Server{
		Addr:    ":2828",
		Handler: router,
	}

	// Initializing the assist in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			global.LOG.Error().Msgf("listen: %s\n", err)
		} else {
			global.LOG.Debug().Msgf("Listen Port : %s", ":2828")
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

package zlog

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/zdate"
)

func Infof(format string, args ...any) {
	slog.Default().Info(fmt.Sprintf(format, args...))
}

func Errorf(format string, args ...any) {
	slog.Default().Error(fmt.Sprintf(format, args...))
}

func Debugf(format string, args ...any) {
	slog.Default().Debug(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...any) {
	slog.Default().Warn(fmt.Sprintf(format, args...))
}

func Fatalf(format string, args ...any) {
	slog.Default().Error(fmt.Sprintf(format, args...))
	os.Exit(1)
}

func Fatal(msg string, args ...any) {
	slog.Default().Error(msg, args...)
	os.Exit(1)
}

func AccessLog(c *gin.Context) {
	global.DB.Save(&models.ZapAccessLog{
		Username:   c.GetString("JWT_USERNAME"),
		LogDate:    time.Now().Format(zdate.DATE_TIME_FORMAT),
		IpAddr:     c.ClientIP(),
		UserAgent:  c.Request.UserAgent(),
		RequestUri: c.Request.RequestURI,
	})
}

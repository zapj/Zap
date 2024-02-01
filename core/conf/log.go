package conf

import (
	"log/slog"

	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/pathutil"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LogInit() {
	if global.ZAP_MODE == "DEV" {
	} else {
		logger := &lumberjack.Logger{
			Filename:   pathutil.GetPath("data/logs/zap.log"),
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     30, // days
		}
		slog.SetDefault(slog.New(slog.NewTextHandler(logger, nil)))
	}
}

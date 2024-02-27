package conf

import (
	"log/slog"

	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/pathutil"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LogInit(name string) {
	logger := &lumberjack.Logger{
		Filename:   pathutil.GetPath("data/logs/" + name + ".log"),
		MaxSize:    20, // megabytes
		MaxBackups: 3,
		MaxAge:     30, // days
	}
	showSource := false
	if global.ZAP_MODE == "DEV" {
		showSource = true
	}
	opts := &slog.HandlerOptions{
		AddSource: showSource,
		Level:     slog.LevelInfo,
	}
	slog.SetDefault(slog.New(slog.NewTextHandler(logger, opts)))
}

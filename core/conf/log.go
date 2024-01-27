package conf

import (
	"log/slog"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LogInit() {
	global.LOG = logrus.New()
	if global.ZAP_MODE == "DEV" {
		global.LOG.SetOutput(os.Stdout)
	} else {
		global.LOG.SetOutput(os.Stdout)
		logger := &lumberjack.Logger{
			Filename:   zap.GetPath("data/logs/zap.log"),
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     30, // days
		}

		slog.SetDefault(slog.New(slog.NewTextHandler(logger, nil)))
	}

	// log.SetOutput(&lumberjack.Logger{
	// 	Filename:   zap.GetPath("data/logs/zap.log"),
	// 	MaxSize:    500, // megabytes
	// 	MaxBackups: 3,
	// 	MaxAge:     30, // days
	// })

	global.LOG.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

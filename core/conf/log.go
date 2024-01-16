package conf

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/zapj/zap/core/global"
)

func LogInit() {
	global.LOG.SetOutput(os.Stdout)
	global.LOG.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

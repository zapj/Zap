package conf

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/zapj/zap/assist/global"
)

func LogInit() {
	global.LOG.SetOutput(os.Stderr)
	global.LOG.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

package conf

import (
	"github.com/rs/zerolog"
	"github.com/zapj/zap/assist/global"
	"os"
)

func LogInit() {
	global.LOG = zerolog.New(os.Stderr).With().Timestamp().Logger()
}

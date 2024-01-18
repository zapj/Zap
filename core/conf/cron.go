package conf

import (
	"github.com/robfig/cron/v3"
	"github.com/zapj/zap/core/crons"
	"github.com/zapj/zap/core/global"
)

func InitCrons() {
	global.CRON = cron.New(cron.WithSeconds())
	crons.RegisterCrons()
	global.CRON.Start()
}

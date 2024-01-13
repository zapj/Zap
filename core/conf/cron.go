package conf

import (
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func InitCrons() {
	global.CRON = cron.New(cron.WithSeconds())

	global.CRON.AddFunc("@every 1m", func() {
		avg, err := load.Avg()
		global.LOG.Info("RUN CRON load avg", avg)
		if err == nil {

			global.DB.Save(&models.LoadAvg{
				Load1:  avg.Load1,
				Load5:  avg.Load5,
				Load15: avg.Load15,
			})
		}

	})

	global.CRON.Start()
}

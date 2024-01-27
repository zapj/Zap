package crons

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func RegisterCrons() {
	ZapCrons()
}

func ZapCrons() {
	global.CRON.AddFunc("@every 1m", func() {
		avg, err := load.Avg()
		if err == nil {
			global.DB.Save(&models.ZapLoadAvg{
				Load1:  avg.Load1,
				Load5:  avg.Load5,
				Load15: avg.Load15,
			})
		}

	})

	global.CRON.AddFunc("@every 4s", func() {
		cpuTimesStat, err := cpu.Times(false)
		if err == nil {
			global.CpuTimesStat = cpuTimesStat[0]
		}

	})
}

package conf

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/zapj/zap/core/global"
)

func ServerStart_INIT() {
	cpuTimesStat, err := cpu.Times(false)
	if err == nil {
		global.CpuTimesStat = cpuTimesStat[0]
	}
}

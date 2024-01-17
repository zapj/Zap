package dashboard

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/datahuman"
)

func DashBoardStats(c *gin.Context) {
	memory, err := mem.VirtualMemory()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, commons.JSONError(1, "read memory error"))
	}
	swapmem, _ := mem.SwapMemory()
	st, err := load.Avg()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, commons.JSONError(1, "read memory error"))
	}
	hostInfo, _ := host.Info()

	// 计算CPU 使用率
	//https://stackoverflow.com/questions/11356330/how-to-get-cpu-usage
	//内存使用率
	// 1
	// global.LOG.Info("pec", (float64(memory.Total-memory.Available))/float64(memory.Total)*100.0)
	// 2
	//strconv.FormatFloat(memory.UsedPercent, 'f', 2, 64),
	pCores, _ := cpu.Counts(false)
	lCores, _ := cpu.Counts(true)
	cpuTImeStat, _ := cpu.Times(false)
	cpusInfo, _ := cpu.Info()
	cpuInfo := cpusInfo[0]
	result := gin.H{"code": 0,
		"msg":  "",
		"user": "",
		"user_stats": gin.H{
			"websiteCount": 1,
		},
		"memory": gin.H{
			"total":        humanize.IBytes(memory.Total),
			"used_percent": strconv.FormatFloat(float64(memory.Total-memory.Available)/float64(memory.Total)*100.0, 'f', 2, 64),
			"free":         humanize.IBytes(memory.Free),
			"used":         humanize.IBytes(memory.Used),
			"available":    humanize.IBytes(memory.Available),
		},
		"swapmem": gin.H{
			"total":        humanize.IBytes(swapmem.Total),
			"used":         humanize.IBytes(swapmem.Used),
			"free":         humanize.IBytes(swapmem.Free),
			"used_percent": strconv.FormatFloat(swapmem.UsedPercent, 'f', 2, 64),
		},
		"load_avg": gin.H{
			"load1":  st.Load1,
			"load5":  st.Load5,
			"load15": st.Load15,
		},
		"cpu": gin.H{
			"cpu_name":       cpuInfo.ModelName,
			"Cores":          cpuInfo.Cores,
			"used_percent":   strconv.FormatFloat(datahuman.CpuPercent(global.CpuTimesStat, cpuTImeStat[0]), 'f', 2, 64),
			"physical_count": pCores,
			"logical_count":  lCores,
		},
		"host": gin.H{
			"BootTime": hostInfo.BootTime,
			"Hostname": hostInfo.Hostname,
			// "uptime1":          hostInfo.Uptime,
			"uptime":          humanize.Time(time.Unix(int64(hostInfo.BootTime), 0)),
			"Platform":        hostInfo.Platform,
			"PlatformFamily":  hostInfo.PlatformFamily,
			"PlatformVersion": hostInfo.PlatformVersion,
			"KernelArch":      hostInfo.KernelArch,
			"KernelVersion":   hostInfo.KernelVersion,
			"OS":              hostInfo.OS,
			"Procs":           hostInfo.Procs,
		},
	}

	c.JSON(200, result)
}

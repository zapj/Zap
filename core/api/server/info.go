package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/zapj/zap/core/api/commons"
)

func ServerInfo(c *gin.Context) {

	memory, err := mem.VirtualMemory()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, commons.JSONError(1, "read memory error"))
	}
	swapmem, err := mem.SwapMemory()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, commons.JSONError(1, "read memory error"))
	}
	st, err := load.Avg()
	if err != nil {
		st = &load.AvgStat{Load1: 0, Load5: 0, Load15: 0}
	}
	hostInfo, _ := host.Info()
	cpusInfo, _ := cpu.Info()
	cpuInfo := cpusInfo[0]

	result := gin.H{
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

		"load1":        st.Load1,
		"load5":        st.Load5,
		"load15":       st.Load15,
		"system_load1": (st.Load1 / float64(cpusInfo[0].Cores)) * 100,

		"cpu_name":  cpuInfo.ModelName,
		"cpu_cores": len(cpusInfo),
		"cpu_ghz":   strconv.FormatFloat(cpuInfo.Mhz/1000, 'f', 2, 64),

		"BootTime": hostInfo.BootTime,
		"Hostname": hostInfo.Hostname,
		"uptime":   humanize.Time(time.Unix(int64(hostInfo.BootTime), 0)),
		// "uptime1":         time.Now().Sub(time.Unix(int64(hostInfo.BootTime), 0)).String(),
		"Platform":        hostInfo.Platform,
		"PlatformFamily":  hostInfo.PlatformFamily,
		"PlatformVersion": hostInfo.PlatformVersion,
		"KernelArch":      hostInfo.KernelArch,
		"KernelVersion":   hostInfo.KernelVersion,
		"OS":              hostInfo.OS,
		"Procs":           hostInfo.Procs,

		// "disk": gin.H{
		// 	"used_percent":        strconv.FormatFloat(diskUsage.UsedPercent, 'f', 2, 64),
		// 	"used":                humanize.Bytes(diskUsage.Used),
		// 	"total":               humanize.Bytes(diskUsage.Total),
		// 	"path":                diskUsage.Path,
		// 	"inodes_used_percent": strconv.FormatFloat(diskUsage.InodesUsedPercent, 'f', 2, 64),
		// 	"inodes_used":         humanize.Bytes(diskUsage.InodesUsed),
		// 	"inodes_total":        humanize.Bytes(diskUsage.InodesTotal),
		// 	"inodes_free":         humanize.Bytes(diskUsage.InodesFree),
		// 	"fstype":              diskUsage.Fstype,
		// 	"free":                humanize.Bytes(diskUsage.Free),
		// },
	}
	diskPartStat, _ := disk.Partitions(true)
	result["diskPartStat"] = diskPartStat

	c.JSON(200, gin.H{"code": 0, "msg": "", "data": result})
}

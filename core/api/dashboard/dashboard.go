package dashboard

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
	"github.com/shirou/gopsutil/v3/net"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/datahuman"
)

func DashBoardStats(c *gin.Context) {
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

	var hostInfo *host.InfoStat
	if hostInfoAny, ok := global.CACHE.Get("zap_host_info"); ok {
		hostInfo, _ = hostInfoAny.(*host.InfoStat)
	} else {
		hostInfo, _ = host.Info()
		global.CACHE.SetDefault("zap_host_info", hostInfo)
	}

	var diskUsage *disk.UsageStat
	if diskUsageAny, ok := global.CACHE.Get("zap_disk_usage_root"); ok {
		diskUsage, _ = diskUsageAny.(*disk.UsageStat)
	} else {
		diskUsage, _ = disk.Usage("/")
		global.CACHE.SetDefault("zap_disk_usage_root", diskUsage)
	}

	// 计算CPU 使用率
	//https://stackoverflow.com/questions/11356330/how-to-get-cpu-usage
	//内存使用率
	// 1
	// global.LOG.Info("pec", (float64(memory.Total-memory.Available))/float64(memory.Total)*100.0)
	// 2
	//strconv.FormatFloat(memory.UsedPercent, 'f', 2, 64),
	var pCores int = 0
	if pCoresObj, ok := global.CACHE.Get("zap_cpu_p_core"); ok {
		pCores, _ = pCoresObj.(int)
	} else {
		pCores, _ = cpu.Counts(false)
		global.CACHE.SetDefault("zap_cpu_p_core", pCores)
	}
	var lCores int = 0
	if lCoresObj, ok := global.CACHE.Get("zap_cpu_p_core"); ok {
		lCores, _ = lCoresObj.(int)
	} else {
		lCores, _ = cpu.Counts(true)
		global.CACHE.SetDefault("zap_cpu_l_core", lCores)
	}

	cpuTImeStat, _ := cpu.Times(false)

	var cpuInfo cpu.InfoStat
	if cacheOBj, ok := global.CACHE.Get("zap_cpu_info"); ok {
		cpuInfo, _ = cacheOBj.(cpu.InfoStat)
	} else {
		cpusInfo, _ := cpu.Info()
		cpuInfo = cpusInfo[0]
		global.CACHE.SetDefault("zap_cpu_info", cpuInfo)
	}

	// upTime := time.Now().Sub(time.Unix(int64(hostInfo.BootTime), 0))
	// upTimeStr := fmt.Sprintf("时%d分%d秒%d",upTime.Seconds())
	result := gin.H{"code": 0,
		"msg":  "",
		"user": "",
		"user_stats": gin.H{
			"websiteCount":  1,
			"databaseCount": 2,
			"cronjobCount":  2,
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
			"uptime":   humanize.Time(time.Unix(int64(hostInfo.BootTime), 0)),
			// "uptime1":         time.Now().Sub(time.Unix(int64(hostInfo.BootTime), 0)).String(),
			"Platform":        hostInfo.Platform,
			"PlatformFamily":  hostInfo.PlatformFamily,
			"PlatformVersion": hostInfo.PlatformVersion,
			"KernelArch":      hostInfo.KernelArch,
			"KernelVersion":   hostInfo.KernelVersion,
			"OS":              hostInfo.OS,
			"Procs":           hostInfo.Procs,
		},
		"disk": gin.H{
			"used_percent":        strconv.FormatFloat(diskUsage.UsedPercent, 'f', 2, 64),
			"used":                humanize.Bytes(diskUsage.Used),
			"total":               humanize.Bytes(diskUsage.Total),
			"path":                diskUsage.Path,
			"inodes_used_percent": strconv.FormatFloat(diskUsage.InodesUsedPercent, 'f', 2, 64),
			"inodes_used":         humanize.Bytes(diskUsage.InodesUsed),
			"inodes_total":        humanize.Bytes(diskUsage.InodesTotal),
			"inodes_free":         humanize.Bytes(diskUsage.InodesFree),
			"fstype":              diskUsage.Fstype,
			"free":                humanize.Bytes(diskUsage.Free),
		},
	}
	// disk.IOCounters()
	// result["diskIOCounters"] =

	netstat, _ := net.IOCounters(true)
	result["netIOCounters"] = netstat

	c.JSON(200, result)
}

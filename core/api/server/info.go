package server

import (
	"bytes"
	"fmt"
	"net/http"
	"regexp"
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
	"github.com/shirou/gopsutil/v3/process"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/task"
	"github.com/zapj/zap/core/utils/cmdutil"
	datautils "github.com/zapj/zap/core/utils/data_utils"
	"github.com/zapj/zap/core/utils/time_utils"
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
		"memory_total":         humanize.IBytes(memory.Total),
		"memory_used_percent":  strconv.FormatFloat(float64(memory.Total-memory.Available)/float64(memory.Total)*100.0, 'f', 2, 64),
		"memory_free":          humanize.IBytes(memory.Free),
		"memory_used":          humanize.IBytes(memory.Used),
		"memory_available":     humanize.IBytes(memory.Available),
		"swapmem_total":        humanize.IBytes(swapmem.Total),
		"swapmem_used":         humanize.IBytes(swapmem.Used),
		"swapmem_free":         humanize.IBytes(swapmem.Free),
		"swapmem_used_percent": strconv.FormatFloat(swapmem.UsedPercent, 'f', 2, 64),
		"load1":                st.Load1,
		"load5":                st.Load5,
		"load15":               st.Load15,
		"system_load1":         (st.Load1 / float64(len(cpusInfo))) * 100,

		"cpu_name":  cpuInfo.ModelName,
		"cpu_cores": len(cpusInfo),
		"cpu_ghz":   strconv.FormatFloat(cpuInfo.Mhz/1000, 'f', 2, 64),

		"BootTime": time.Unix(int64(hostInfo.BootTime), 0).Format(time_utils.DATE_TIME_FORMAT),
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
	diskPartStat, _ := disk.Partitions(false)
	// global.LOG.Info("diskPartStat", diskPartStat)
	var diskUsageStat []*disk.UsageStat
	for i := 0; i < len(diskPartStat); i++ {
		usageStat, err := disk.Usage(diskPartStat[i].Mountpoint)
		usageStat.Fstype = diskPartStat[i].Fstype
		if err != nil {
			continue
		}
		diskUsageStat = append(diskUsageStat, usageStat)

	}
	result["disk"] = diskUsageStat

	c.JSON(200, gin.H{"code": 0, "msg": "", "data": result})
}

func ServerProcessList(c *gin.Context) {
	processList, err := process.Processes()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, commons.JSONError(1, "没有权限"))

	}
	var procList []gin.H = make([]gin.H, len(processList))

	for i, proc := range processList {
		pMap := gin.H{
			"pid": proc.Pid,
		}
		pMap["name"], _ = proc.Name()
		create_time, _ := proc.CreateTime()
		pMap["create_time"] = time_utils.FormatUnixMilliToDateTime(create_time)
		pMap["background"], _ = proc.Background()
		cpu_percent, _ := proc.CPUPercent()
		pMap["cpu_percent"] = datautils.FmtPercentFloat64(cpu_percent)
		// pMap["is_running"], _ = proc.IsRunning()
		mem_percent, _ := proc.MemoryPercent()
		pMap["mem_percent"] = datautils.FmtPercentFloat32(mem_percent)
		// pMap["mem_percent"] = datautils.FmtPercentFloat64(pMap["mem_percent"].(float32))
		pMap["cmdline"], _ = proc.Cmdline()
		pMap["username"], _ = proc.Username()

		procList[i] = pMap

	}
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": procList})
}

func ServerNetInterfaces(c *gin.Context) {
	netinterfaceList, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"code": 0, "msg": "", "data": netinterfaceList})
}

func ServerTopInfo(c *gin.Context) {
	out, err := cmdutil.ExecCmd("top", "-bn1")
	// // out, err := cmdutil.ExecCmdString("top", []string{"-bn1","-E","g", -e m})
	if err != nil {
		fmt.Println(err)
	}
	// scanner := bufio.NewScanner(bytes.NewBuffer(out))
	// scanner.Split(bufio.ScanWords)
	// for scanner.Scan() {
	// 	fmt.Printf("%q ", scanner.Text())
	// }

	buf := bytes.NewBuffer(out)
	var header bytes.Buffer
	// var process bytes.Buffer
	//read header
	for {
		line, err := buf.ReadBytes('\n')
		if err != nil {
			break
		}
		header.Write(line)
		if len(line) < 2 {
			buf.ReadString('\n')
			break
		}
	}
	var rows [][]string
	pattern := "\\S+"

	regex := regexp.MustCompile(pattern)

	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}

		rows = append(rows, regex.FindAllString(line, -1))

	}

	c.JSON(200, gin.H{"code": 0, "msg": "", "data": header.String(), "rows": rows})
}

func UpgradeCheck(c *gin.Context) {
	if global.ZAP_INFO.Version == "0.0.0" {
		c.JSON(200, gin.H{"code": 0, "msg": "开发环境不能执行更新操作"})
	}
	_, err := task.Get("/upgrade")
	if err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": "任务执行失败"})
	}
	c.JSON(200, gin.H{"code": 0, "msg": "开发环境不能执行更新操作"})
}

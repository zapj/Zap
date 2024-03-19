package server

import (
	"strconv"
	"strings"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/zapj/zap/core/api/commons"
)

func PortList(c *gin.Context) {
	conns, err := net.Connections("all")
	if err != nil {
		c.JSON(200, commons.Error(1, err.Error(), nil))
		return
	}
	portList := []map[string]string{}
	for _, conn := range conns {
		if conn.Pid == 0 {
			continue
		}
		proc, err := process.NewProcess(conn.Pid)
		if err != nil {
			continue
		}
		procName, err := proc.Name()
		if err != nil {
			continue
		}
		protocol := protocolToString(conn.Type, conn.Laddr.IP)
		localAddr := conn.Laddr.IP + ":" + strconv.Itoa(int(conn.Laddr.Port))
		remoteAddr := conn.Raddr.IP + ":" + strconv.Itoa(int(conn.Raddr.Port))

		portList = append(portList, map[string]string{
			"local_addr":  localAddr,
			"remote_addr": remoteAddr,
			"protocol":    protocol,
			"status":      conn.Status,
			"pid":         strconv.Itoa(int(conn.Pid)),
			"proc_name":   procName,
		})
	}

	c.JSON(200, commons.Success(portList))
}

func protocolToString(connType uint32, ip string) string {
	isIPv6 := strings.Contains(ip, ":")
	switch connType {
	case syscall.SOCK_STREAM: // TCP
		if isIPv6 {
			return "TCP6"
		}
		return "TCP"
	case syscall.SOCK_DGRAM: // UDP
		if isIPv6 {
			return "UDP6"
		}
		return "UDP"
	default:
		return "Unknown"
	}
}

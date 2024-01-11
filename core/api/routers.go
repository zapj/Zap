package api

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
)

func RegisterRouter(c *gin.RouterGroup) {
	c.GET("/meminfo", MemoryUsageInfo)
	c.GET("/loadavg", func(c *gin.Context) {
		st, _ := load.Avg()
		c.String(200, "Loadavg  %s", st.String())
	})
	c.GET("/mem", func(c *gin.Context) {
		v, _ := mem.VirtualMemory()
		c.String(200, v.String())
	})
}

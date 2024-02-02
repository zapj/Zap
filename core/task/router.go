package task

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.GET("/upgrade", func(c *gin.Context) {
		_, ok := GL.Load("upgrade")
		if ok {
			c.JSON(200, gin.H{"code": 0, "msg": "系统升级中"})
			return
		}
		go UpgradeZapd()
		c.JSON(200, gin.H{"code": 0, "msg": "提交成功，系统升级成功后将自动重启控制面板"})
	})
	r.POST("/task/put", func(c *gin.Context) {
		global.DB.Save(&models.ZapTask{Cmd: "test", Retry: 0, Title: "test 1", Status: STATUS_WAIT})
	})
	r.POST("/task/cancel", func(c *gin.Context) {})
	r.POST("/task/list", func(c *gin.Context) {})
	r.POST("/task/refresh", func(c *gin.Context) {
		LoadTask()
	})
}

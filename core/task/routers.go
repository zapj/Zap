package task

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func RegisterRouter(r *gin.RouterGroup) {
	r.POST("/systemctl", SystemdAction)
	r.POST("/upgrade", UpgradeAction)

	r.POST("/task/install/app", AppInstallAction)
	r.POST("/task/cancel", TaskCancelAction)
	r.POST("/task/refresh", func(c *gin.Context) {
		LoadTask()
		c.JSON(200, gin.H{"code": 0, "msg": "刷新成功"})
	})
	r.POST("/task/put", func(c *gin.Context) {
		global.DB.Save(&models.ZapTask{Cmd: "test", Retry: 0, Title: "test 1", Status: STATUS_WAIT})
	})

	r.POST("/app/uninstall", AppUninstallAction)
}

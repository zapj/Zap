package users

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/global"
)

func SyncUserSettings(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": gin.H{
		"version":    global.ZAP_INFO.Version,
		"build_date": global.ZAP_INFO.BuildDate,
		"theme":      "light",
	}})
}

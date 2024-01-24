package users

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core"
)

func SyncUserSettings(c *gin.Context) {
	c.JSON(200, gin.H{"code": 0, "msg": "", "data": gin.H{
		"version": core.Version,
		"theme":   "light",
	}})
}

package appstore

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/workflows"
)

func ListApp(c *gin.Context) {
	workFLows := workflows.ReadZAPWorkflowsInfo()

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": workFLows,
	})
}

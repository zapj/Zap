package appstore

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
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

// 应用安装任务列表
func TaskList(c *gin.Context) {
	result := []models.ZapTask{}
	global.DB.Order("id desc").Find(&result)
	slog.Info("result", "data", result)
	c.JSON(200, commons.Success(result))
}

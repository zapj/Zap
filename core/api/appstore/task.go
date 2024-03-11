package appstore

import (
	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/task"
	"github.com/zapj/zap/core/utils/str"
	"github.com/zapj/zap/core/zapi"
)

// 应用安装任务列表
func TaskList(c *gin.Context) {
	result := []models.ZapTask{}
	global.DB.Order("id desc").Find(&result)
	c.JSON(200, commons.Success(result))
}

func RemoveTask(c *gin.Context) {
	id := c.PostForm("id")
	global.DB.Delete(&models.ZapTask{}, id)
	c.JSON(200, commons.Success("删除成功"))
}

func GenTask(c *gin.Context) {
	global.DB.Save(&models.ZapTask{
		Cmd:      "ls -la",
		Retry:    0,
		Title:    str.RandUString(12),
		Status:   task.STATUS_WAIT,
		TaskType: task.TASK_TYPE_APPSTORE,
		CreateBy: "admin",
	})
	c.JSON(200, commons.Success("生成成功"))
}

func CancelTask(c *gin.Context) {
	id := c.PostForm("id")
	resp := zapi.NewZapi().CancelTask(map[string]any{
		"id": id,
	})
	// if err != nil {
	// 	c.JSON(200, commons.Error(1, "error", err.Error()))
	// 	return
	// }
	c.Data(resp.Status, gin.MIMEJSON, resp.Data)
}

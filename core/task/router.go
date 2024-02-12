package task

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/jsonutil"
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
	r.POST("/task/install/app", func(c *gin.Context) {
		id := c.PostForm("id")
		version := c.PostForm("version")
		action := c.PostForm("action")
		appStore := models.ZapAppStore{}
		if err := global.DB.First(&appStore, "app_id = ?", id).Error; err != nil {
			c.JSON(200, gin.H{"code": 1, "msg": "App无法安装"})
			return
		}
		if result := global.DB.First(&models.ZapApps{}, "name = ? and version = ?", appStore.Name, version); result.RowsAffected > 0 {
			c.JSON(200, gin.H{"code": 1, "msg": "App无法安装"})
			return
		}

		global.DB.Save(&models.ZapTask{
			Cmd:         jsonutil.EncodeToString(map[string]string{"action": action, "app_id": id, "app_version": version}),
			Retry:       0,
			TaskType:    TASK_TYPE_APPSTORE,
			CreateBy:    "admin",
			Title:       fmt.Sprintf("安装 %s", appStore.Title),
			Status:      STATUS_WAIT,
			ExtendsAttr: fmt.Sprint(appStore.Id),
		})

		global.DB.Save(&models.ZapApps{
			AppStoreId:  appStore.Id,
			Name:        appStore.Name,
			Status:      global.APP_STATUS_INSTALL,
			InstallBy:   "admin",
			InstallDate: time.Now(),
			Title:       appStore.Title,
			Description: appStore.Description,
			Version:     version,
		})
		LoadTask()
		c.JSON(200, gin.H{"code": 0, "msg": "提交成功，系统将自动安装App"})
	})
	r.POST("/task/cancel", func(c *gin.Context) {
		id := c.PostForm("id")
		taskId, _ := strconv.Atoi(id)
		CancelTask(uint(taskId))
	})
	// r.POST("/task/list", func(c *gin.Context) {})
	r.POST("/task/refresh", func(c *gin.Context) {
		LoadTask()
		c.JSON(200, gin.H{"code": 0, "msg": "刷新成功"})
	})
}

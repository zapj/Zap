package task

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/base"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/jsonutil"
	"github.com/zapj/zap/core/utils/pathutil"
)

func AppInstallAction(c *gin.Context) {
	id := c.PostForm("id")
	version := c.PostForm("version")
	action := c.PostForm("action")
	appStore := models.ZapAppStore{}
	if err := global.DB.First(&appStore, "app_id = ?", id).Error; err != nil {
		c.JSON(200, gin.H{"code": 1, "msg": "App无法安装"})
		return
	}
	app := models.ZapApps{}
	if result := global.DB.First(&app, "name = ? and version = ?", appStore.Name, version); result.RowsAffected > 0 {
		c.JSON(200, gin.H{"code": 1, "msg": "App已存在或正在安装"})
		return
	}
	app = models.ZapApps{
		AppStoreId:  appStore.Id,
		Name:        appStore.Name,
		Status:      global.APP_STATUS_INSTALL,
		AppType:     appStore.Category,
		InstallBy:   "admin",
		InstallDate: base.LocalTime(time.Now()),
		Title:       appStore.Title,
		Description: appStore.Description,
		Version:     version,
	}
	global.DB.Create(&app)

	global.DB.Save(&models.ZapTask{
		Cmd:         jsonutil.EncodeToString(base.ZapMap{"action": action, "appstore_unique_id": id, "app_id": app.Id, "appstore_id": appStore.Id, "app_version": version}),
		Retry:       0,
		TaskType:    TASK_TYPE_APPSTORE,
		CreateBy:    "admin",
		Title:       fmt.Sprintf("安装 %s", appStore.Title),
		Status:      STATUS_WAIT,
		LogFile:     pathutil.GetPath(fmt.Sprintf("data/logs/install_%d.log", app.Id)),
		TargetDir:   pathutil.GetPath(fmt.Sprintf("data/build/%s-%s", app.Name, app.Version)),
		ExtendsAttr: fmt.Sprint(appStore.Id),
	})

	LoadTask()
	c.JSON(200, gin.H{"code": 0, "msg": "提交成功,系统将自动安装App"})
}

func TaskCancelAction(c *gin.Context) {
	id := c.PostForm("id")
	taskId, _ := strconv.Atoi(id)
	CancelTask(uint(taskId))
}

package appstore

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-zoox/fetch"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/task"
	"github.com/zapj/zap/core/utils/str"
	"github.com/zapj/zap/core/workflows"
)

func ListApp(c *gin.Context) {
	appInfoList := workflows.ReadAppstoreList()
	alreadyInstallApps := []models.ZapApps{}
	global.DB.Select("name,title").Find(&alreadyInstallApps, "status NOT IN ?", []string{global.APP_STATUS_UNINSTALL})
	alreadyApps := make(map[string]string)
	alreadyAppsVersion := make(map[string]string)
	for _, v := range alreadyInstallApps {
		alreadyApps[v.Name] = v.Title
		alreadyAppsVersion[fmt.Sprintf("%s-%s", v.Name, v.Version)] = v.Title
	}
	alreadyInstallApps = nil
	c.JSON(200, gin.H{
		"code":               0,
		"msg":                "OK",
		"data":               appInfoList,
		"alreadyApps":        alreadyApps,
		"alreadyAppsVersion": alreadyAppsVersion,
	})
}

func ListInstalledApp(c *gin.Context) {
	installedApps := []models.ZapApps{}
	global.DB.Find(&installedApps)
	c.JSON(200, commons.Success(installedApps))
}

func UninstallApp(c *gin.Context) {
	id := c.PostForm("id")
	app := models.ZapApps{}
	global.DB.First(&app, id)
	if app.Status == global.APP_STATUS_UNINSTALL {
		c.JSON(200, commons.Error(1, "应用正在卸载", nil))
		return
	}
	app.Status = global.APP_STATUS_UNINSTALL
	global.DB.Save(&app)
	resp, err := task.Post("/app/uninstall", &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		Body: map[string]string{
			"id": id,
		}})

	if err != nil {
		c.JSON(200, commons.Error(1, "应用卸载失败", err.Error()))
		return
	}
	c.Data(200, gin.MIMEJSON, resp.Body)
}

func AppInstall(c *gin.Context) {
	id := c.PostForm("id")
	action := c.PostForm("action")
	version := c.PostForm("version")
	app := &models.ZapAppStore{}
	if err := global.DB.First(app, "app_id = ?", id).Error; err != nil {
		c.JSON(200, commons.Error(1, "应用不存在", nil))
		return
	}
	resp, err := task.Post("/task/install/app", &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		Body: map[string]string{
			"id":      id,
			"action":  action,
			"version": version,
		},
	})

	if err != nil {
		c.JSON(200, commons.Error(1, "error", err.Error()))
		return
	}
	c.Data(200, gin.MIMEJSON, resp.Body)
}

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
	_, err := task.Post("/task/cancel", &fetch.Config{
		Headers: map[string]string{
			"Content-Type": "application/x-www-form-urlencoded",
		},
		Body: map[string]string{
			"id": id,
		},
	})
	if err != nil {
		c.JSON(200, commons.Error(1, "error", err.Error()))
		return
	}
	c.JSON(200, commons.Success("取消成功"))
}

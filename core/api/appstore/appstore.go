package appstore

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/sysapi/process"
	"github.com/zapj/zap/core/task"
	"github.com/zapj/zap/core/utils/str"
	"github.com/zapj/zap/core/workflows"
	"github.com/zapj/zap/core/zapi"
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
	err := global.DB.Find(&installedApps).Error
	if err != nil {
		c.JSON(200, commons.Error(1, "查询应用失败"+err.Error(), installedApps))
		return
	}

	for i, app := range installedApps {
		if app.PidFile == "" {
			app.AppStatus = "stop"
		}
		pid := process.GetPidFile(app.PidFile)
		if pid == 0 {
			app.AppStatus = "stop"
		} else {
			_, err := os.FindProcess(pid)
			if err == nil {
				app.AppStatus = "running"
			}
		}

		installedApps[i].AppStatus = app.AppStatus
		global.DB.Save(app)
	}

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
	// app.Status = global.APP_STATUS_UNINSTALL
	// global.DB.Save(&app)
	resp := zapi.Client.UninstallApp(map[string]any{
		"id": id,
	})

	// if err != nil {
	// 	c.JSON(200, commons.Error(1, "应用卸载失败", err.Error()))
	// 	return
	// }
	c.Data(resp.Status, gin.MIMEJSON, resp.Data)
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
	resp := zapi.Client.InstallApp(map[string]any{
		"id":      id,
		"action":  action,
		"version": version,
	})

	// if err != nil {
	// 	c.JSON(200, commons.Error(1, "error", err.Error()))
	// 	return
	// }
	c.Data(resp.Status, gin.MIMEJSON, resp.Data)
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
	resp := zapi.NewZapi().CancelTask(map[string]any{
		"id": id,
	})
	// if err != nil {
	// 	c.JSON(200, commons.Error(1, "error", err.Error()))
	// 	return
	// }
	c.Data(resp.Status, gin.MIMEJSON, resp.Data)
}

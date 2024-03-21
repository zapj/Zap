package appstore

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/sysapi/process"
	"github.com/zapj/zap/core/utils/propsutil"
	"github.com/zapj/zap/core/workflows"
	"github.com/zapj/zap/core/zapi"
)

func ListApp(c *gin.Context) {
	workflows.ReadAppstoreList()

	appstoreList := []models.ZapAppStore{}
	global.DB.Model(&models.ZapAppStore{}).Find(&appstoreList)
	alreadyInstallApps := getAlreadyInstalledApps()

	for i, asRow := range appstoreList {
		if _, ok := alreadyInstallApps[asRow.Name]; ok {
			if asRow.AllowMultipleInstances == "yes" {
				appstoreList[i].AllowInstallation = "yes"
			} else {
				appstoreList[i].AllowInstallation = "no"
			}
			// appstoreList[i].AllowInstallation = asRow.AllowMultipleInstances
			appstoreList[i].AlreadyInstalledVers = alreadyInstallApps[asRow.Name]
		} else {
			appstoreList[i].AllowInstallation = "yes"
			appstoreList[i].AlreadyInstalledVers = []string{}
		}

		// if asRow.Actions != "" {
		// 	appstoreList[i].ActionsData = jsonutil.DecodeToMap(asRow.Actions)
		// }
	}

	c.JSON(200, gin.H{
		"code": 0,
		"msg":  "OK",
		"data": appstoreList,
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
		} else {
			pid := process.GetPidFile(app.PidFile)
			if pid == 0 {
				app.AppStatus = "stop"
			} else {
				_, err := os.FindProcess(pid)
				if err == nil {
					app.AppStatus = "running"
				}
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

	resp := zapi.Client.UninstallApp(map[string]any{
		"id": id,
	})
	c.Data(resp.Status, gin.MIMEJSON, resp.Data)
}

func AppInstall(c *gin.Context) {
	id := c.PostForm("id")
	action := c.PostForm("action")
	version := c.PostForm("version")
	if action == "" {
		action = "install"
	}
	if id == "" || id == "0" {
		c.JSON(200, commons.Error(1, "应用ID不能为空", nil))
		return
	}
	app := models.ZapAppStore{}
	if err := global.DB.First(&app, "app_id = ?", id).Error; err != nil {
		c.JSON(200, commons.Error(1, "应用不存在", nil))
		return
	}
	resp := zapi.Client.InstallApp(map[string]any{
		"id":      id,
		"action":  action,
		"version": version,
	})

	c.Data(resp.Status, gin.MIMEJSON, resp.Data)
}

func SaveAppSettings(c *gin.Context) {
	id := c.PostForm("id")
	app := models.ZapApps{}
	if err := global.DB.First(&app, id).Error; err != nil {
		c.JSON(200, commons.Error(1, "应用不存在", nil))
		return
	}
	app.PidFile = c.PostForm("pid_file")
	app.ConfigFile = c.PostForm("config_file")
	app.Expose = propsutil.FmtString(c.PostForm("expose"))
	app.Settings = propsutil.FmtString(c.PostForm("settings"))

	if err := global.DB.Save(&app).Error; err != nil {
		c.JSON(200, commons.Error(1, "保存失败", nil))
		return
	}
	c.JSON(200, commons.Success("保存成功"))
}

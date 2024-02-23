package task

import (
	"fmt"
	"log/slog"
	"os/exec"
	"path"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/base"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/jsonutil"
	"github.com/zapj/zap/core/utils/pathutil"
	"github.com/zapj/zap/core/utils/zaputil"
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
			InstallDate: time.Now(),
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
	r.POST("/app/uninstall", func(c *gin.Context) {
		id := c.PostForm("id")
		app := models.ZapApps{}
		slog.Info("app id", "app", app)
		if err := global.DB.First(&app, id).Error; err != nil {
			c.JSON(200, commons.Error(1, "应用不存在", nil))
			return
		}
		if app.Status != global.APP_STATUS_UNINSTALL || app.Name == "" {
			c.JSON(200, commons.Error(1, "应用无法卸载", nil))
			return
		}

		appPath := path.Join(pathutil.GetPath("data/appstore"), app.Name)
		osArch := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)
		// 默认使用action
		scriptName := "uninstall.sh"
		var scriptFile string
		// build.sh
		if fileutils.IsFile(path.Join(appPath, scriptName)) {
			scriptFile = path.Join(appPath, scriptName)
		} else if fileutils.IsFile(path.Join(appPath, fmt.Sprintf("%s_%s.sh", "uninstall", osArch))) { // build_linux_amd64.sh
			scriptFile = path.Join(appPath, fmt.Sprintf("%s_%s.sh", "uninstall", osArch))
		} else {
			c.JSON(200, commons.Error(1, fmt.Sprintf("scriptFile not found %s", scriptName), nil))
			return
		}

		//stop service
		installPath := fmt.Sprintf("/usr/local/apps/%s-%s", app.Name, zaputil.GetMajorMinorVersion(app.Version))
		cmd := exec.Command("bash", scriptFile)
		cmd.Env = append(cmd.Env,
			"APP_PATH="+installPath, //安装目录
			"APP_VERSION="+app.Version,
			"SHORT_VERSION="+zaputil.GetMajorMinorVersion(app.Version),
			"APP_NAME="+app.Name,
			"APP_ID="+id,
		)
		err := cmd.Run()
		if err != nil {
			c.JSON(200, commons.Error(1, "卸载失败:"+err.Error(), nil))
			return
		}
		global.DB.Delete(&app)

		c.JSON(200, commons.Error(0, "卸载成功", nil))
	})
}

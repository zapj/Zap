package task

import (
	"fmt"
	"os/exec"
	"path"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/api/commons"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/pathutil"
	"github.com/zapj/zap/core/utils/zaputil"
)

func AppUninstallAction(c *gin.Context) {
	id := c.PostForm("id")
	app := models.ZapApps{}
	if err := global.DB.First(&app, id).Error; err != nil {
		c.JSON(200, commons.Error(1, "应用不存在", nil))
		return
	}
	// if app.Status != global.APP_STATUS_UNINSTALL || app.Name == "" {
	// 	c.JSON(200, commons.Error(1, "应用无法卸载", nil))
	// 	return
	// }

	appPath := path.Join(pathutil.GetPath("data/appstore"), app.AppType, app.Name)
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
		c.JSON(200, commons.Error(1, fmt.Sprintf("shell script not found %s", scriptName), nil))
		return
	}

	//stop service
	cmd := exec.Command("bash", scriptFile)
	cmd.Env = append(cmd.Env,
		"APP_PATH="+app.InstallDir, //安装目录
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
}

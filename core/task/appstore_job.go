package task

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/jsonutil"
	"github.com/zapj/zap/core/utils/pathutil"
)

type AppStoreJob struct {
	Ctx      context.Context
	TaskData *models.ZapTask
	Cmd      string
}

func (b *AppStoreJob) Execute() error {
	defer func() {
		if err := recover(); err != nil {
			// log.Error(err)
			slog.Error("appstore job execute failed", "err", err)
		}
	}()
	slog.Info("appstore job execute")
	cmdMap := jsonutil.DecodeToZapMap(b.TaskData.Cmd)
	appId := cmdMap.GetString("app_id")
	action := cmdMap.GetString("action")

	appstore := models.ZapAppStore{}
	result := global.DB.First(&appstore, "app_id = ? ", appId)
	if result.RowsAffected < 1 {
		return fmt.Errorf("App_ID : %s 不存在 ", appId)
	}

	app := models.ZapApps{}
	result = global.DB.First(&app, "app_store_id = ?", appstore.Id)
	if result.RowsAffected < 1 {
		return fmt.Errorf("AppStore_ID : %d 不存在 ", appstore.Id)
	}

	options := jsonutil.DecodeToZapMap(appstore.Options)
	actionOptions := options.GetZapMap(action)
	scriptName := actionOptions.GetString(fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH))
	if scriptName == "" {
		scriptName = fmt.Sprintf("%s.sh", action)
	}
	// scriptFile :=
	appPath := path.Join(pathutil.GetPath("data/appstore"), appstore.Name)
	scriptFile := path.Join(appPath, scriptName)
	slog.Error("scriptFile", "file", scriptFile)
	cmd := exec.CommandContext(b.Ctx, "bash", scriptFile)
	envs := []string{
		"APP_PATH=" + appPath,
		"SCRIPT_PATH=" + scriptFile,
		"APPS_DIR=" + global.APPS_DIR,
		"APP_ID=" + fmt.Sprint(appId),
		"APP_NAME=" + app.Name,
		"APP_STORE_ID=" + fmt.Sprint(app.AppStoreId),
		"APP_STORE_NAME=" + appstore.Name,
		"APP_VERSION=" + app.Version,
		"APP_TITLE=" + app.Title,
	}
	cmd.Env = append(os.Environ(), envs...)
	if err := cmd.Run(); err != nil {
		return err
	}
	app.Status = global.APP_STATUS_ACTIVE
	global.DB.Save(&app)
	return nil
}

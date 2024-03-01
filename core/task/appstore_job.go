package task

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"

	cp "github.com/otiai10/copy"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/jsonutil"
	"github.com/zapj/zap/core/utils/pathutil"
)

type AppStoreJob struct {
	Ctx              context.Context
	TaskData         *models.ZapTask
	Cmd              string
	AppId            int    //apps ID
	AppStoreUniqueId string // appstore unique id
	AppStoreId       int    // appstore id
}

func (b *AppStoreJob) Execute() (err error) {
	defer func() {
		if err2 := recover(); err2 != nil {
			if r, ok := err2.(error); ok {
				err = r
			} else {
				err = fmt.Errorf("job execute failed %q", err2)
			}
			stackTrace := make([]byte, 4096)
			length := runtime.Stack(stackTrace, false)
			slog.Error("appstore job execute failed", "err", err2, "strack", stackTrace[:length])
		}
	}()
	cmdMap := jsonutil.DecodeToZapMap(b.TaskData.Cmd)
	b.AppId = cmdMap.GetInt("app_id")
	b.AppStoreUniqueId = cmdMap.GetString("appstore_unique_id")
	b.AppStoreId = cmdMap.GetInt("appstore_id")
	action := cmdMap.GetString("action")
	version := cmdMap.GetString("app_version")

	appstore := models.ZapAppStore{}
	result := global.DB.First(&appstore, "id = ? ", b.AppStoreId)
	if result.RowsAffected < 1 {
		return fmt.Errorf("App_ID : %d 不存在 ", b.AppStoreId)
	}

	app := models.ZapApps{}
	result = global.DB.First(&app, "id = ?", b.AppId)
	if result.RowsAffected < 1 {
		return fmt.Errorf("AppStore_ID : %d 不存在 ", appstore.Id)
	}

	if !fileutils.IsDir(pathutil.GetPath("data/pkg")) {
		os.MkdirAll(pathutil.GetPath("data/pkg"), 0755)
	}
	if !fileutils.IsDir(pathutil.GetPath("data/build")) {
		os.MkdirAll(pathutil.GetPath("data/build"), 0755)
	}

	// options := jsonutil.DecodeToZapMap(appstore.Options)
	// actionOptions := options.GetZapMap(action)

	// appstore script path
	appStoreScriptPath := path.Join(pathutil.GetPath("data/appstore"), appstore.Name)
	osArch := fmt.Sprintf("%s_%s", runtime.GOOS, runtime.GOARCH)
	// 默认使用action
	scriptName := fmt.Sprintf("%s.sh", action)
	var scriptFile string
	appInstallScriptDir := pathutil.GetPath("data/build", fmt.Sprintf("build-%s", appstore.Name))
	if !fileutils.IsDir(appInstallScriptDir) {
		if err := cp.Copy(appStoreScriptPath, appInstallScriptDir); err != nil {
			return err
		}
	}
	// build.sh
	if fileutils.IsFile(path.Join(appInstallScriptDir, scriptName)) {
		scriptFile = path.Join(appInstallScriptDir, scriptName)
	} else if fileutils.IsFile(path.Join(appInstallScriptDir, fmt.Sprintf("%s_%s.sh", action, osArch))) { // build_linux_amd64.sh
		scriptFile = path.Join(appInstallScriptDir, fmt.Sprintf("%s_%s.sh", action, osArch))
	} else {
		return fmt.Errorf("scriptFile not found %s", scriptName)
	}

	// if !strings.HasPrefix(b.TaskData.TargetDir, global.ZAP_BASE_DIR) {
	// 	return fmt.Errorf("target dir is not allowed %s", b.TaskData.TargetDir)
	// }

	// if fileutils.IsDir(b.TaskData.TargetDir) {
	// 	os.RemoveAll(b.TaskData.TargetDir)
	// }

	// os.MkdirAll(b.TaskData.TargetDir, 0755)
	logFile, err := os.Create(b.TaskData.LogFile)
	if err != nil {
		return err
	}
	defer logFile.Close()

	cmd := exec.CommandContext(b.Ctx, "bash", scriptFile)
	appVers := strings.Split(version, ".")
	majorVersion, minorVersion, revisionVersion := "0", "0", "0"
	if len(appVers) > 0 {
		majorVersion = appVers[0]
	}
	if len(appVers) > 1 {
		minorVersion = appVers[1]
	}
	if len(appVers) > 2 {
		revisionVersion = appVers[2]
	}
	platform, family, os_version, _ := host.PlatformInformation()
	envs := []string{
		"ZAPCTL=" + pathutil.GetPath("zapctl"),
		"ZAP_OS=" + runtime.GOOS,
		"ZAP_PATH=" + global.ZAP_BASE_DIR,
		"OS_ARCH=" + runtime.GOARCH,
		"OS_PLATFORM=" + platform,
		"OS_FAMILY=" + family,
		"OS_VERSION=" + os_version,

		// "ZAP_BASE_DIR=" + global.ZAP_BASE_DIR,
		"ZAP_DATA_PATH=" + pathutil.GetPath("data"),
		"APPSTORE_PATH=" + appStoreScriptPath,
		"USERSCRIPT_PATH=" + appStoreScriptPath,
		"APP_PATH=" + fmt.Sprintf("%s/%s-%s.%s", global.APPS_DIR, app.Name, majorVersion, minorVersion),
		"SCRIPT_PATH=" + scriptFile,
		"APPS_DIR=" + global.APPS_DIR,
		"PKG_PATH=" + pathutil.GetPath("data/pkg"),
		// "APP_STORE_ID=" + fmt.Sprint(app.AppStoreId),
		// "APP_STORE_NAME=" + appstore.Name,
		"BUILD_PATH=" + pathutil.GetPath("data/build"),
		"LOG_FILE=" + b.TaskData.LogFile,
		"APP_ID=" + fmt.Sprint(app.Id),
		"APP_NAME=" + app.Name,
		"APP_VERSION=" + version,
		"MAJOR_VERSION=" + majorVersion,
		"MINOR_VERSION=" + minorVersion,
		"REVISION_VERSION=" + revisionVersion,
		"APP_TITLE=" + app.Title,
		"CPU_NUM=" + fmt.Sprint(runtime.NumCPU()),
	}
	// var errBuffer bytes.Buffer
	// errMixed := io.MultiWriter(logFile, &errBuffer)
	cmd.Env = append(os.Environ(), envs...)
	// cmdOutput, err := cmd.StdoutPipe()
	// cmdStderr, err := cmd.StderrPipe()
	cmd.Stderr = logFile
	cmd.Stdout = logFile
	// writer := bufio.NewWriter(logFile)
	// defer writer.Flush()

	if err := cmd.Start(); err != nil {
		// logFile.Write(errBuffer.Bytes())
		// return errors.Join(err, errors.New(errBuffer.String()))
		return err
	}

	// go io.Copy(writer, io.MultiReader(cmdOutput, cmdStderr))

	err = cmd.Wait()
	if err != nil {
		// logFile.Write(errBuffer.Bytes())
		// return errors.Join(err, errors.New(errBuffer.String()))
		return err
	}
	app.Status = global.APP_STATUS_ACTIVE
	global.DB.Save(&app)
	return nil
}

func (b *AppStoreJob) Clean() {
	global.DB.Delete(&models.ZapApps{}, b.AppId)
}

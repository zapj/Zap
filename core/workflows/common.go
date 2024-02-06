package workflows

import (
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/zapj/goutils"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/pathutil"
	"gopkg.in/yaml.v3"
)

var ZAP_SCRIPT_FILES = "data/appstore/**/app.y*ml"

func ReadAppstoreList() []*AppInfo {
	scriptFiles, err := filepath.Glob(pathutil.GetPath(ZAP_SCRIPT_FILES))
	if err != nil {
		slog.Error("Read appstore", "err", err.Error())
		return nil
	}
	var appInfoList []*AppInfo

	var totalApp int64 = 0
	var saveAppFlag = false

	if global.DB.Model(&models.ZapAppStore{}).Count(&totalApp).Error != nil {
		totalApp = 0
	}
	slog.Info("Read appstore", "totalApp", totalApp)
	if len(scriptFiles) != int(totalApp) {
		global.DB.Where("1 = 1").Delete(&models.ZapAppStore{})
		saveAppFlag = true
	}
	for _, ymlPath := range scriptFiles {
		slog.Info("Read appstore", "appstore", ymlPath)
		ymlBytes, _ := os.ReadFile(ymlPath)
		appInfo := &AppInfo{}
		yaml.Unmarshal(ymlBytes, appInfo)
		appInfo.ScriptName, _ = filepath.Rel("data/appstore", ymlPath)
		appInfo.Id = goutils.MD5(appInfo.ScriptName)
		appInfoList = append(appInfoList, appInfo)
		if saveAppFlag {
			global.DB.Save(&models.ZapAppStore{
				AppId:            appInfo.Id,
				AppName:          appInfo.Name,
				Icon:             appInfo.Icon,
				AppTitle:         appInfo.Title,
				AppVersion:       strings.Join(appInfo.Version, ","),
				ScriptName:       appInfo.ScriptName,
				Description:      appInfo.Description,
				Author:           appInfo.Author,
				OrganizationName: appInfo.OrganizationName,
			})
		}
	}
	return appInfoList
}

func ReadZAPWorkflows() []*Workflows {
	scriptFiles, _ := filepath.Glob(ZAP_SCRIPT_FILES)
	var workflows []*Workflows
	for _, ymlPath := range scriptFiles {
		ymlBytes, _ := os.ReadFile(ymlPath)
		flow := &Workflows{}
		yaml.Unmarshal(ymlBytes, flow)
		flow.ScriptName, _ = filepath.Rel("data/appstore", ymlPath)
		flow.Id = goutils.MD5(flow.ScriptName)
		workflows = append(workflows, flow)
	}
	return workflows
}

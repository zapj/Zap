package workflows

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/zapj/goutils"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/jsonutil"
	"github.com/zapj/zap/core/utils/pathutil"
	"github.com/zapj/zap/core/utils/zlog"
	"gopkg.in/yaml.v3"
)

var ZAP_SCRIPT_FILES = "data/appstore/**/app.y*ml"

func ReadAppstoreList() error {
	scriptFiles, err := filepath.Glob(pathutil.GetPath(ZAP_SCRIPT_FILES))
	if err != nil {
		return err
	}

	var totalApp int64 = 0
	var saveAppFlag = false

	if global.DB.Model(&models.ZapAppStore{}).Count(&totalApp).Error != nil {
		totalApp = 0
	}

	if len(scriptFiles) != int(totalApp) {
		global.DB.Where("1 = 1").Delete(&models.ZapAppStore{})
		global.DB.Exec("DELETE FROM SQLITE_SEQUENCE WHERE name='zap_app_stores'")
		saveAppFlag = true
	}
	//force update
	global.DB.Where("1 = 1").Delete(&models.ZapAppStore{})
	global.DB.Exec("DELETE FROM SQLITE_SEQUENCE WHERE name='zap_app_stores'")
	saveAppFlag = true

	for _, ymlPath := range scriptFiles {
		ymlBytes, _ := os.ReadFile(ymlPath)
		appInfo := &AppInfo{}
		yaml.Unmarshal(ymlBytes, appInfo)
		zlog.Infof("app info: %#v", appInfo)
		appInfo.ConfigName = ymlPath
		appInfo.Id = goutils.MD5(appInfo.ConfigName)

		if saveAppFlag {

			global.DB.Save(&models.ZapAppStore{
				AppId:                  appInfo.Id,
				Name:                   appInfo.Name,
				Icon:                   appInfo.Icon,
				Title:                  appInfo.Title,
				Category:               appInfo.Category,
				ConfigName:             appInfo.ConfigName,
				AllowMultipleInstances: appInfo.AllowMultipleInstances,
				Version:                strings.Join(appInfo.Version, ","),
				Tags:                   strings.Join(appInfo.Tags, ","),
				Description:            appInfo.Description,
				Author:                 appInfo.Author,
				OrganizationName:       appInfo.OrganizationName,
				Actions:                jsonutil.EncodeToString(appInfo.Actions),
				Options:                jsonutil.EncodeToString(appInfo.Options),
			})
		}
	}
	return nil
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

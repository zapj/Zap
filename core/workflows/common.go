package workflows

import (
	"os"
	"path/filepath"

	"github.com/zapj/goutils"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/pathutil"
	"gopkg.in/yaml.v3"
)

var ZAP_SCRIPT_FILES = "data/appstore/zap/*.yaml"

func ReadZAPWorkflowsInfo() []*WorkflowsInfo {
	scriptFiles, _ := filepath.Glob(pathutil.GetPath(ZAP_SCRIPT_FILES))
	var workflows []*WorkflowsInfo

	var totalApp int64 = 0
	var saveAppFlag = false
	global.DB.Model(&models.ZapAppStore{}).Count(&totalApp)
	if len(scriptFiles) != int(totalApp) {
		global.DB.Delete(&models.ZapAppStore{})
		saveAppFlag = true
	}
	for _, ymlPath := range scriptFiles {
		ymlBytes, _ := os.ReadFile(ymlPath)
		flow := &WorkflowsInfo{}
		yaml.Unmarshal(ymlBytes, flow)
		flow.ScriptName, _ = filepath.Rel("data/appstore", ymlPath)
		flow.Id = goutils.MD5(flow.ScriptName)
		workflows = append(workflows, flow)
		if saveAppFlag {
			global.DB.Save(&models.ZapAppStore{
				AppId:            flow.Id,
				AppName:          flow.Name,
				Icon:             flow.Icon,
				AppTitle:         flow.Title,
				AppVersion:       flow.Version,
				ScriptName:       flow.ScriptName,
				Description:      flow.Description,
				Author:           flow.Author,
				OrganizationName: flow.OrganizationName,
			})
		}
	}
	return workflows
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

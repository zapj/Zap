package appstore

import (
	"strings"

	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func getAlreadyInstalledApps() map[string][]string {
	var apps []models.ZapApps
	var result = make(map[string][]string, 0)
	global.DB.Select("name,group_concat(version) as version").Group("name").Find(&apps, "status NOT IN ?", []string{global.APP_STATUS_UNINSTALL})
	for _, app := range apps {
		result[app.Name] = strings.Split(app.Version, ",")
	}
	return result
}

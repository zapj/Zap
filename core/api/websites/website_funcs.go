package websites

import (
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func CountWebsite(uid int) int {
	var count int64 = 0
	global.DB.Model(&models.ZapWebSite{}).Where("uid = ?", uid).Count(&count)
	return int(count)
}

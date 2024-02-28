package crons

import (
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func CountCrons(uid int) int {
	var count int64 = 0
	global.DB.Model(&models.ZapCrontab{}).Where("uid = ?", uid).Count(&count)
	return int(count)
}

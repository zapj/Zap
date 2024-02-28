package database

import (
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
)

func CountDataBase(uid int) int {
	var count int64 = 0
	global.DB.Model(&models.ZapDataBase{}).Where("uid = ?", uid).Count(&count)
	return int(count)
}

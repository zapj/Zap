package store

import (
	"github.com/zapj/zap/core/global"
	"gorm.io/gorm"
)

func DB() *gorm.DB {
	return global.DB
}

func StatisticsDB() *gorm.DB {
	return global.StatisticsDB
}

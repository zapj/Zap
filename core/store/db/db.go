package db

import (
	"github.com/zapj/zap/core/global"
	"gorm.io/gorm"
)

func First(dist any, conds ...any) *gorm.DB {
	return global.DB.First(dist, conds...)
}

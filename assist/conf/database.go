package conf

import (
	"github.com/glebarez/sqlite"
	"github.com/zapj/zap/assist/global"
	"gorm.io/gorm"
)

func DbInit() {
	db, err := gorm.Open(sqlite.Open("data/zap.db"), &gorm.Config{})
	if err != nil {
		global.LOG.Error(err)
	}
	global.DB = db
}

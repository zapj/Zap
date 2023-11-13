package conf

import (
	"github.com/glebarez/sqlite"
	"github.com/zapj/zap/assist/global"
	"gorm.io/gorm"
)

func DbInit() {
	db, err := gorm.Open(sqlite.Open("zap.db"), &gorm.Config{})
	if err != nil {

	}
	global.DB = db
}

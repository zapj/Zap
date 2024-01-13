package conf

import (
	"github.com/glebarez/sqlite"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"gorm.io/gorm"
)

func DbInit() {
	db, err := gorm.Open(sqlite.Open("data/zap.db"), &gorm.Config{})
	if err != nil {
		global.LOG.Error(err)
	}

	db.AutoMigrate(&models.LoadAvg{})

	global.DB = db
}

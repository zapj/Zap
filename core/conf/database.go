package conf

import (
	"errors"

	"github.com/glebarez/sqlite"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/bcrypt_util"
	"gorm.io/gorm"
)

func DbInit() {
	db, err := gorm.Open(sqlite.Open("data/zap.db"), &gorm.Config{})
	if err != nil {
		global.LOG.Error(err)
	}
	db.AutoMigrate(&models.ZapUsers{})
	db.AutoMigrate(&models.LoadAvg{})
	db.AutoMigrate(&models.ZapWebSite{})
	db.AutoMigrate(&models.ZapDataBase{})

	user := models.ZapUsers{}
	result := db.First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user.ID = 1
		user.Username = "admin"
		passwd, _ := bcrypt_util.HashPassword("zapzap")
		user.Password = passwd
		user.Home = "/home/" + user.Username
		user.Gid = 1000
		user.Uid = 1000
		user.Shell = "/usr/sbin/nologin"
		db.Save(&user)
	}
	global.DB = db
}

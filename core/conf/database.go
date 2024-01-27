package conf

import (
	"errors"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/bcrypt_util"
	"github.com/zapj/zap/core/utils/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbInit() {
	var err error
	global.DB, err = gorm.Open(sqlite.Open(zap.GetPath("data/zap.db")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		global.LOG.Error("DB init error")
		os.Exit(1)
	}
	global.StatisticsDB, err = gorm.Open(sqlite.Open(zap.GetPath("data/statistics.db")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		global.LOG.Error("DB init error")
		os.Exit(1)
	}

	global.DB.AutoMigrate(&models.ZapUsers{})
	global.DB.AutoMigrate(&models.ZapLoadAvg{})
	global.DB.AutoMigrate(&models.ZapWebSite{})
	global.DB.AutoMigrate(&models.ZapDataBase{})
	global.DB.AutoMigrate(&models.ZapDiskIOCounters{})
	global.DB.AutoMigrate(&models.ZapNetIOCounters{})

	global.DB.AutoMigrate(&models.ZapAppStore{})

	user := models.ZapUsers{}
	result := global.DB.First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user.ID = 1
		user.Username = "admin"
		passwd, _ := bcrypt_util.HashPassword("zapzap")
		user.Password = passwd
		user.Home = "/home/" + user.Username
		user.Gid = 1000
		user.Uid = 1000
		user.Shell = "/usr/sbin/nologin"
		global.DB.Save(&user)
	}
}

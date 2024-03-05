package conf

import (
	"errors"
	"log/slog"
	"os"

	"github.com/glebarez/sqlite"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"github.com/zapj/zap/core/utils/pathutil"
	"github.com/zapj/zap/core/utils/zaputil"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func DbInit() {
	var err error
	global.DB, err = gorm.Open(sqlite.Open(pathutil.GetPath("data/zap.db")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		slog.Error("DB init error", "err", err)
		os.Exit(1)
	}
	global.StatisticsDB, err = gorm.Open(sqlite.Open(pathutil.GetPath("data/statistics.db")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		slog.Error("DB init error")
		os.Exit(1)
	}

	//DB
	global.DB.AutoMigrate(&models.ZapUsers{})
	global.DB.AutoMigrate(&models.ZapWebSite{})
	global.DB.AutoMigrate(&models.ZapDataBase{})
	global.DB.AutoMigrate(&models.ZapAppStore{})
	global.DB.AutoMigrate(&models.ZapApps{})

	//system task job
	global.DB.AutoMigrate(&models.ZapTask{})
	//user cron
	global.DB.AutoMigrate(&models.ZapCrontab{})

	// StatisticsDB
	global.StatisticsDB.AutoMigrate(&models.ZapLoadAvg{})

	global.StatisticsDB.AutoMigrate(&models.ZapDiskIOCounters{})
	global.StatisticsDB.AutoMigrate(&models.ZapNetIOCounters{})
	global.StatisticsDB.AutoMigrate(&models.ZapAccessLog{})

	user := models.ZapUsers{}
	result := global.DB.First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		user.ID = 1
		user.Username = "admin"
		passwd, _ := zaputil.HashPassword("zapzap")
		user.Password = passwd
		user.Home = "/home/" + user.Username
		user.Gid = 1000
		user.Uid = 1000
		user.Shell = "/usr/sbin/nologin"
		global.DB.Save(&user)
	}
}

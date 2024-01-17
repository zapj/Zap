package global

import (
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

var DB *gorm.DB
var LOG *logrus.Logger = logrus.New()
var CONFIG *ini.File
var CRON *cron.Cron

var CpuTimesStat cpu.TimesStat

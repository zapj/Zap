package global

import (
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

var DB *gorm.DB
var StatisticsDB *gorm.DB
var LOG *logrus.Logger = logrus.New()
var CONFIG *ini.File
var CRON *cron.Cron

var CpuTimesStat cpu.TimesStat

var CACHE *cache.Cache

var NET_IO_COUNTERS []net.IOCountersStat

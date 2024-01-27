package global

import (
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/sirupsen/logrus"
	"github.com/zapj/zap/core"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

var ZAP_MODE string
var ZAP_BASE_DIR string
var ServerPID int
var DB *gorm.DB
var StatisticsDB *gorm.DB
var LOG *logrus.Logger
var CONFIG *ini.File
var CRON *cron.Cron

var CpuTimesStat cpu.TimesStat

var CACHE *cache.Cache

var NET_IO_COUNTERS []net.IOCountersStat
var ZAP_INFO *core.ZapInfo

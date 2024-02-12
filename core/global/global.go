package global

import (
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/zapj/zap/core"
	"gopkg.in/ini.v1"
	"gorm.io/gorm"
)

var ZAP_MODE string
var ZAP_BASE_DIR string
var ServerPID int
var DB *gorm.DB
var StatisticsDB *gorm.DB
var CONFIG *ini.File
var CRON *cron.Cron

var CpuTimesStat cpu.TimesStat

var CACHE *cache.Cache

var NET_IO_COUNTERS []net.IOCountersStat
var ZAP_INFO *core.ZapInfo
var SERVER_CONF *ServerConf

var APPS_DIR string = "/usr/local/apps"

// app status
const (
	APP_STATUS_RUNNING   = "running"
	APP_STATUS_INSTALL   = "install"
	APP_STATUS_UNINSTALL = "uninstall"
	APP_STATUS_STOP      = "stop"
	APP_STATUS_ACTIVE    = "active"
)

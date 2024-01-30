package conf

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/zap"
	"gopkg.in/ini.v1"
)

func InitEnv() {
	global.ZAP_MODE = "PROD"
	if os.Getenv("ZAP_MODE") == "DEV" {
		global.ZAP_MODE = "DEV"
	}
	var err error
	execpath, err := os.Executable()
	if err != nil {
		global.ZAP_BASE_DIR = "/usr/local/zap"
	} else {
		global.ZAP_BASE_DIR = filepath.Dir(execpath)
	}

	if !fileutils.IsDir(zap.GetPath("data")) {
		err = os.MkdirAll(zap.GetPath("data/logs"), 0755)
		if err != nil {
			slog.Error("mkdir data/logs", err)
			os.Exit(1)
		}
	}

	global.CONFIG, err = ini.Load(zap.GetPath("conf/zap.ini"))
	if err != nil {
		slog.Error("Fail to read file: %v", err)
	}
	global.SERVER_CONF = &global.ServerConf{
		Host:     global.CONFIG.Section("zapd").Key("host").MustString(""),
		Port:     global.CONFIG.Section("zapd").Key("port").MustInt(2828),
		IPv6:     global.CONFIG.Section("zapd").Key("ipv6").MustBool(true),
		CertFile: global.CONFIG.Section("zapd").Key("cert_file").MustString(zap.GetPath("data/server.crt")),
		KeyFile:  global.CONFIG.Section("zapd").Key("key_file").MustString(zap.GetPath("data/server.key")),
	}

}

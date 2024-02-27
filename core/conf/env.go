package conf

import (
	"log/slog"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/zapj/go-properties"
	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/base"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/cmdutil"
	"github.com/zapj/zap/core/utils/pathutil"
	"github.com/zapj/zap/core/utils/str"
	"github.com/zapj/zap/core/utils/zaputil"
	"github.com/zapj/zap/core/utils/zlog"
)

func InitEnv() {
	global.ZAP_MODE = "PROD"

	var err error
	execpath, err := os.Executable()
	if err != nil {
		global.ZAP_BASE_DIR = "/usr/local/zap"
	} else if strings.Contains(execpath, "go-build") {
		global.ZAP_BASE_DIR, _ = os.Getwd()
	} else {
		global.ZAP_BASE_DIR = filepath.Dir(execpath)
	}

	if !fileutils.IsDir(pathutil.GetPath("data")) {
		err = os.MkdirAll(pathutil.GetPath("data/logs"), 0755)
		if err != nil {
			slog.Error("mkdir data/logs", "err", err)
			os.Exit(1)
		}
	}
	global.CONFIG, err = properties.LoadFile(pathutil.GetPath("conf/zap.conf"), properties.UTF8)
	if err != nil {
		zlog.Errorf("Fail to read file: %v", err)
	}

	global.SERVER_CONF = &base.ServerConf{
		SigningKey: global.CONFIG.GetString("signing_key", str.RandUString(16)),
		Host:       global.CONFIG.GetString("host", "0.0.0.0"),
		Port:       global.CONFIG.GetInt("port", 2800),
		IPv6:       global.CONFIG.GetBool("ipv6", true),
		CertFile:   global.CONFIG.GetString("cert_file", pathutil.GetPath("data/server.crt")),
		KeyFile:    global.CONFIG.GetString("key_file", pathutil.GetPath("data/server.key")),

		WwwRoot:  global.CONFIG.GetString("www_root", "/home/www"),
		WwwUser:  global.CONFIG.GetString("www_user", "www"),
		WwwGroup: global.CONFIG.GetString("www_group", "www"),

		ZapMode: strings.ToUpper(global.CONFIG.GetString("zap_mode", "PROD")),
	}
	global.ZAP_MODE = strings.ToUpper(global.CONFIG.GetString("zap_mode", "PROD"))
	// check user
	wwwUserInfo, err := user.Lookup(global.SERVER_CONF.WwwUser)
	if err != nil {
		_, err = cmdutil.ExecCmd("adduser", "-M", "-s", "/bin/false", global.SERVER_CONF.WwwUser)
		if err != nil {
			slog.Error("adduser www", "err", err)
			os.Exit(1)
		}
		wwwUserInfo, _ = user.Lookup(global.SERVER_CONF.WwwUser)
	}
	global.SERVER_CONF.WwwUserId = zaputil.MustConvertStringToInt(wwwUserInfo.Uid)
	global.SERVER_CONF.WwwGroupId = zaputil.MustConvertStringToInt(wwwUserInfo.Gid)

	if os.Getenv("ZAP_MODE") == "DEV" {
		global.ZAP_MODE = "DEV"
	}

}

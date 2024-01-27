package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/zapj/goutils/fileutils"
	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/utils/zap"
	"gopkg.in/ini.v1"
)

func InitEnv() {
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
			fmt.Println(err)
			os.Exit(1)
		}
	}

	global.CONFIG, err = ini.Load(zap.GetPath("conf/zap.ini"))
	if err != nil {
		global.LOG.Errorf("Fail to read file: %v", err)
	}
	// fmt.Println(global.CONFIG.Section("server").Key("port").MustInt(2828))
}

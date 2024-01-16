package conf

import (
	"github.com/zapj/zap/core/global"
	"gopkg.in/ini.v1"
)

func InitEnv() {
	var err error
	global.CONFIG, err = ini.Load("conf/zap.ini")
	if err != nil {
		global.LOG.Errorf("Fail to read file: %v", err)
	}
	// fmt.Println(global.CONFIG.Section("server").Key("port").MustInt(2828))
}

package task

import (
	"sync"

	"github.com/robfig/cron/v3"
)

var ZAP_MIRROR_URL = "https://mirrors.zap.cn/zap"
var TASK_SERV_SOCK = "/var/run/zap_task"
var GL sync.Map
var globalTask sync.Map

type InstallApp struct {
	ID    int
	Retry int
	C     *cron.Cron
}

type TaskCommand struct {
	Title     string   `json:"title"`
	Cmd       string   `json:"cmd"`
	TargetDir string   `json:"target_dir"`
	LogFile   string   `json:"log_file"`
	Args      []string `json:"args"`
	Env       []string `json:"env"`
}

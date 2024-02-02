package task

import (
	"sync"

	"github.com/robfig/cron/v3"
)

var ZAP_MIRROR_URL = "https://mirrors.zap.cn/zap"
var TASK_SERV_SOCK = "data/task"
var GL sync.Map

type InstallApp struct {
	ID    int
	Retry int
	C     *cron.Cron
}

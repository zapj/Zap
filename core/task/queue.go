package task

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/models"
	"gorm.io/gorm"
)

const (
	STATUS_WAIT     = "wait"
	STATUS_READY    = "ready"
	STATUS_SUSPEND  = "suspend"
	STATUS_RUNNING  = "running"
	STATUS_COMPLETE = "complete"
	STATUS_CANCEL   = "cancel"
)

var globalTaskQueue = make(chan ZapJobInterface, 1)

type ZapJobInterface interface {
	Execute()
}
type ZapJob struct {
	TaskCtx    context.Context
	TaskCancel context.CancelFunc
	JobId      int
	Cmd        string
	StartTime  int64
	EndTime    int64
	Retry      int
	Task       *models.ZapTask
}

func Cancel() {
}

func (z *ZapJob) Cancel() {
	z.TaskCancel()
}

func (z *ZapJob) Execute() {
	//运行中
	z.Task.Status = STATUS_RUNNING
	global.DB.Updates(z.Task)
	//默认 Retry 次数为1
	if z.Retry == 0 {
		z.Retry = 1
	}
	for i := 1; i <= z.Retry; i++ {
		select {
		case <-z.TaskCtx.Done():
			z.Task.Status = STATUS_CANCEL
			global.DB.Updates(z.Task)
			slog.Info("task cancel")
			return
		default:
			//执行任务
			slog.Info("z.Task.Cmd", "cmd", z.Task.Cmd)
			slog.Info("task complete")
			time.Sleep(time.Second * 20)
			z.Task.RetryCount = i //更新执行次数
			z.Task.EndTime = time.Now().Unix()
			z.Task.Status = STATUS_COMPLETE
			global.DB.Updates(z.Task)
			if z.Task.Status == STATUS_COMPLETE {
				LoadTask() //加载下一条任务
				return
			}
		}

	}

}

func NewZapJob() *ZapJob {
	var taskCtx, taskCancel = context.WithCancel(context.Background())
	return &ZapJob{
		TaskCtx:    taskCtx,
		TaskCancel: taskCancel,
		StartTime:  time.Now().Unix(),
		Retry:      0,
	}
}

func AddJob(job ZapJobInterface) {
	globalTaskQueue <- job
}

func LoadTask() {
	taskData := &models.ZapTask{}

	result := global.DB.Order("priority desc, id asc").First(&taskData, "status=?", "wait")
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return
	}
	taskJob := NewZapJob()
	taskJob.Task = taskData
	globalTaskQueue <- taskJob
	taskData.StartTime = taskJob.StartTime
	taskData.Status = STATUS_READY
	global.DB.Updates(taskData)
}

func StartTaskQueue() {
	slog.Info("StartTaskQueue")
	LoadTask()
	go taskQueue()
}

func taskQueue() {
	for v := range globalTaskQueue {
		//读取下一条
		v.Execute()
	}
}

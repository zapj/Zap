package task

import (
	"context"
	"errors"
	"log/slog"
	"os/exec"
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

const (
	TASK_TYPE_INSTALL = "install"
)

const (
	INSTALL_RUNNING_KEY = "install_task_running"
)

var globalTaskQueue = make(chan ZapJobInterface, 10)

// var rwGTQ sync.Mutex

type ZapJobInterface interface {
	Execute()
}
type ZapJob struct {
	TaskCtx    context.Context
	TaskCancel context.CancelFunc
	TaskType   string
	JobId      uint
	Cmd        string
	StartTime  int64
	EndTime    int64
	Retry      int
	TaskData   *models.ZapTask
	Status     bool
}

func CancelTask(id uint) {
	taskJob, ok := globalTask.Load(id)
	if !ok {
		slog.Error("task not found", "id", id)
		return
	}
	slog.Info("taskJob", "taskJob", taskJob)
	taskJob.(*ZapJob).TaskCancel()
	if taskJob.(*ZapJob).TaskType == TASK_TYPE_INSTALL {
		GL.Delete(INSTALL_RUNNING_KEY)
	}
	globalTask.Delete(id)
}

func (z *ZapJob) Cancel() {
	z.TaskCancel()
}

func (z *ZapJob) Execute() {
	//运行中
	z.TaskData.Status = STATUS_RUNNING
	z.TaskData.StartTime = time.Now().Unix()
	global.DB.Save(z.TaskData)
	//默认 Retry 次数为1
	if z.Retry == 0 {
		z.Retry = 1
	}
	for i := 1; i <= z.Retry; i++ {
		select {
		case <-z.TaskCtx.Done():
			z.TaskData.Status = STATUS_CANCEL
			global.DB.Save(z.TaskData)
			globalTask.Delete(z.TaskData.Id)
			if z.TaskType == TASK_TYPE_INSTALL {
				GL.Delete(INSTALL_RUNNING_KEY)
			}
			slog.Info("task cancel")
			return
		default:
			//执行任务
			slog.Info("z.TaskData.Cmd", "cmd", z.TaskData.Cmd, "ID", z.TaskData.Id)

			if err := exec.CommandContext(z.TaskCtx, "sleep", "60").Run(); err != nil {
				slog.Info("task error", "err", err)
			}
			z.TaskData.RetryCount = i //更新执行次数
			z.TaskData.EndTime = time.Now().Unix()
			z.TaskData.Status = STATUS_COMPLETE
			global.DB.Save(z.TaskData)
			globalTask.Delete(z.TaskData.Id)
			if z.TaskType == TASK_TYPE_INSTALL {
				GL.Delete(INSTALL_RUNNING_KEY)
			}
			slog.Info("task complete", "id", z.TaskData.Id)
			if z.TaskData.Status == STATUS_COMPLETE {
				LoadTask()
				return
			}
		}

	}

}

func NewZapJob() *ZapJob {
	taskCtx, taskCancel := context.WithCancel(context.Background())
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

func LoadTask(status ...string) {
	if len(status) == 0 {
		status = append(status, STATUS_WAIT, STATUS_READY)
	}
	// rwGTQ.Lock()
	// defer rwGTQ.Unlock()

	taskList := []models.ZapTask{}

	err := global.DB.Order("id asc,status asc,priority desc").Where("status IN ?", status).Find(&taskList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	for _, v := range taskList {
		taskData := v
		if _, ok := GL.Load(INSTALL_RUNNING_KEY); taskData.TaskType == TASK_TYPE_INSTALL && ok {
			continue
		}
		if _, ok := globalTask.Load(taskData.Id); ok {
			continue
		}
		taskData.Status = STATUS_READY
		global.DB.Save(&taskData)
		taskJob := NewZapJob()
		taskJob.JobId = taskData.Id
		taskJob.TaskData = &taskData
		taskJob.TaskType = taskData.TaskType
		if taskData.TaskType == TASK_TYPE_INSTALL {
			GL.Store("install_task_running", true)
		}
		globalTask.Store(taskData.Id, taskJob)
		select {
		case globalTaskQueue <- taskJob:
		default:
		}

	}

}

func StartTaskQueue() {
	slog.Info("StartTaskQueue")
	LoadTask()
	go taskQueue()
}

func taskQueue() {
	for v := range globalTaskQueue {
		//读取下一条
		go v.Execute()
	}
}

package task

import (
	"context"
	"errors"
	"log/slog"
	"sync"
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
	STATUS_FAILED   = "failed"
)

const (
	TASK_TYPE_APPSTORE = "appstore"
	TASK_TYPE_EMAIL    = "email"
	TASK_TYPE_HTTP     = "http"
)

const (
	INSTALL_RUNNING_KEY = "install_task_running"
)

var globalTaskQueue = make(chan ZapJobInterface, 10)

var enqueueMux sync.Mutex

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

func CancelTask(id uint) bool {
	taskJob, ok := globalTask.Load(id)
	if !ok {
		return false
	}
	taskJob.(*ZapJob).TaskCancel()
	if taskJob.(*ZapJob).TaskType == TASK_TYPE_APPSTORE {
		GL.Delete(INSTALL_RUNNING_KEY)
	}
	globalTask.Delete(id)
	return true
}

func (z *ZapJob) Cancel() {
	z.TaskCancel()
}

func (z *ZapJob) Exit(status string, err error) {
	z.TaskData.EndTime = time.Now().Unix()
	z.TaskData.Status = status
	if err != nil {
		z.TaskData.Error = err.Error()
	}
	global.DB.Save(z.TaskData)
	z.Unlock()
	LoadTask()
}

func (z *ZapJob) ExecuteSuccess() {
	z.TaskData.EndTime = time.Now().Unix()
	z.TaskData.Status = STATUS_COMPLETE
	global.DB.Save(z.TaskData)
	z.Unlock()
}

func (z *ZapJob) Unlock() {
	globalTask.Delete(z.TaskData.Id)
	if z.TaskType == TASK_TYPE_APPSTORE {
		GL.Delete(INSTALL_RUNNING_KEY)
	}
}

func (z *ZapJob) Execute() {
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
			z.Exit(STATUS_CANCEL, nil)
			return
		default:
			z.TaskData.RetryCount = i
			if z.TaskType == TASK_TYPE_APPSTORE {

				appJob := &AppStoreJob{
					Ctx:      z.TaskCtx,
					TaskData: z.TaskData,
					Cmd:      z.TaskData.Cmd,
				}
				if err := appJob.Execute(); err != nil {
					appJob.Clean()
					z.Exit(STATUS_FAILED, err)
					return
				}
				z.Exit(STATUS_COMPLETE, nil)
				return

			} else if z.TaskType == TASK_TYPE_EMAIL {
				emailJob := &EmailJob{
					Ctx: z.TaskCtx,
				}
				emailJob.Execute()
				z.Exit(STATUS_COMPLETE, nil)
			} else if z.TaskType == TASK_TYPE_HTTP {
				httpJob := &HttpJob{
					Ctx: z.TaskCtx,
				}
				httpJob.Execute()
				z.Exit(STATUS_COMPLETE, nil)
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
	enqueueMux.Lock()
	defer enqueueMux.Unlock()

	taskList := []models.ZapTask{}

	err := global.DB.Order("id asc,status asc,priority desc").Where("status IN ?", status).Find(&taskList).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	for _, v := range taskList {
		taskData := v
		if _, ok := GL.Load(INSTALL_RUNNING_KEY); taskData.TaskType == TASK_TYPE_APPSTORE && ok {
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
		select {
		case globalTaskQueue <- taskJob:
			if taskData.TaskType == TASK_TYPE_APPSTORE {
				GL.Store("install_task_running", true)
			}
			globalTask.Store(taskData.Id, taskJob)
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

func isCancel(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

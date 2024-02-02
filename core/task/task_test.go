package task_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zapj/zap/core/task"
)

func TestTaskQueue(t *testing.T) {
	task.StartTaskQueue()
	// job := task.NewZapJob()
	// job.Cmd = "测试1"
	// task.AddJob(job)
	// // <-job.Ctx.Done()

	for i := 0; i < 5; i++ {
		job1 := task.NewZapJob()
		job1.Cmd = fmt.Sprintf("测试%d", i)
		task.AddJob(job1)
	}
	fmt.Println("添加完成$")
	go func() {
		time.Sleep(15 * time.Second)
		job1 := task.NewZapJob()
		job1.Cmd = fmt.Sprintln("测试 15")
		task.AddJob(job1)
	}()
	go func() {
		time.Sleep(10 * time.Second)
		task.Cancel()
	}()
	// <-job1.Ctx.Done()
	fmt.Println("main sleep 20s")
	time.Sleep(time.Second * 20)
}

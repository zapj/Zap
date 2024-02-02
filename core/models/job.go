package models

import "gorm.io/gorm"

type ZapTask struct {
	gorm.Model

	Title string `json:"title"` //任务名称

	TaskType  string `json:"task_type"` //script | install
	Spec      string
	Cmd       string
	LogFile   string //日志文件
	TargetDir string //临时文件目录

	Retry      int //重试次数
	RetryCount int
	Priority   int //优先级 优先被执行

	StartTime int64 `json:"start_time"` //启动时间
	EndTime   int64 `json:"end_time"`   //结束时间

	Status string // Status ： wait | suspend | ready | running | complete | cancel
}

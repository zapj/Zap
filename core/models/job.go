package models

import "time"

type ZapTask struct {
	Id uint `json:"id" gorm:"primarykey;autoIncrement"`

	Title     string `json:"title"`                            //任务名称
	TaskType  string `json:"task_type" gorm:"default:install"` // email | http | install
	Cmd       string `json:"cmd"`                              //命令
	LogFile   string `json:"log_file"`                         //日志文件
	TargetDir string `json:"target_dir"`                       //临时文件目录
	Request   string `json:"request"`                          // json  {header:{},body:{}}

	Retry      int `json:"retry"` //重试次数
	RetryCount int `json:"retry_count"`
	Priority   int `json:"priority" gorm:"default:100;"` //优先级 优先被执行，值越小优先级越高 默认100

	StartTime int64     `json:"start_time"` //启动时间
	EndTime   int64     `json:"end_time"`   //结束时间
	CreateBy  string    `json:"create_by"`  //创建用户
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Status    string    // Status ： wait | suspend | ready | running | complete | cancel | failed
}

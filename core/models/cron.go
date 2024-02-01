package models

import "gorm.io/gorm"

type ZapCrontab struct {
	gorm.Model

	// 用户ID
	Uid uint

	//自定义任务名称
	TaskTitle string

	//定时表达式
	Spec string

	//命令
	Cmd string

	// 定时任务
	TickerDate string

	//执行一次
	Once bool

	//任务状态
	Status int
}

func Get_AllCrontab() {

}

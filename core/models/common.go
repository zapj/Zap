package models

import "gorm.io/gorm"

type ZapWebSite struct {
	gorm.Model
	Uid  uint `gorm:"index"`
	Gid  uint
	Home string
}

type ZapDataBase struct {
	gorm.Model
	Uid      uint   `gorm:"index"`
	Dbname   string `json:"db_name"`
	Dbuser   string `json:"db_user"`
	Dbpasswd string `json:"db_passwd"`
}

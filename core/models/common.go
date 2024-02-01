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

type ZapAppStore struct {
	gorm.Model
	AppId            string
	AppName          string
	AppTitle         string
	AppVersion       string
	Description      string
	Icon             string
	ScriptName       string
	Author           string
	OrganizationName string
}

type ZapAccessLog struct {
	gorm.Model
	Username   string
	LogDate    string
	IpAddr     string
	UserAgent  string
	RequestUri string
}

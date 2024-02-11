package models

import (
	"time"

	"gorm.io/gorm"
)

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
	ConfigName       string `json:"config_name"`
	Category         string `json:"category"`
	Icon             string `json:"icon"`
	Author           string
	OrganizationName string
	Tags             string
	PublishAt        time.Time `json:"publish_at"`
	Actions          string
	Options          string
}

type ZapAccessLog struct {
	gorm.Model
	Username   string
	LogDate    string
	IpAddr     string
	UserAgent  string
	RequestUri string
}

type ZapApps struct {
	Id          uint      `json:"id" gorm:"primarykey;autoIncrement"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Version     string    `json:"version" gorm:"default:0.0.0"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"index;"`
	InstallBy   string    `json:"install_by"`
	InstallDate time.Time `json:"install_date" gorm:"autoCreateTime"`
	UpdateDate  time.Time `json:"update_date" gorm:"autoUpdateTime"`
}

package models

import (
	"time"

	"gorm.io/gorm"
)

type ZapWebSite struct {
	gorm.Model
	Uid         uint   `json:"uid" gorm:"index"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Domain      string `json:"domain"`
	OtherDomain string `json:"other_domain"`
	Home        string `json:"home"`
	Settings    string `json:"settings"`
	Status      string `json:"status" gorm:"index;"` // 状态 active
}

type ZapDataBase struct {
	gorm.Model
	Uid      uint   `gorm:"index"`
	Dbname   string `json:"db_name"`
	Dbuser   string `json:"db_user"`
	Dbpasswd string `json:"db_passwd"`
}

type ZapAppStore struct {
	Id               uint `json:"id" gorm:"primarykey;autoIncrement"`
	AppId            string
	Name             string `json:"app_name" gorm:"unique"`
	Title            string `json:"title"`
	Version          string
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
	CreatedAt        time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time `json:"updated_at" gorm:"autoUpdateTime"`
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
	Id         uint   `json:"id" gorm:"primarykey;autoIncrement"`
	AppStoreId uint   `json:"app_store_id" gorm:"index;"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Version    string `json:"version" gorm:"default:0.0.0"`
	AppType    string `json:"app_type"` // webserver , application , database

	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"index;"` // install : 安装中  active : 已安装
	AppStatus   string    `json:"app_status"`           // running : 运行中  stop : 停止中
	InstallBy   string    `json:"install_by"`
	Instance    string    `json:"instance"`    // default
	InstallDir  string    `json:"install_dir"` //安装目录
	Settings    string    `json:"settings"`    // properties
	ConfigFile  string    `json:"config_file"` //配置文件
	PidFile     string    `json:"pid_file"`
	Error       string    `json:"error"` //错误信息
	InstallDate time.Time `json:"install_date" gorm:"autoCreateTime"`
	UpdateDate  time.Time `json:"update_date" gorm:"autoUpdateTime"`
}

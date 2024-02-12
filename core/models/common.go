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
	Id          uint      `json:"id" gorm:"primarykey;autoIncrement"`
	AppStoreId  uint      `json:"app_store_id" gorm:"index;"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Version     string    `json:"version" gorm:"default:0.0.0"`
	Description string    `json:"description"`
	Status      string    `json:"status" gorm:"index;"`
	InstallBy   string    `json:"install_by"`
	InstallDir  string    `json:"install_dir"` //安装目录
	ConfigDirs  string    `json:"config_dirs"` //配置目录
	OtherFiles  string    `json:"other_files"` // json (file/dir list)
	Expose      string    `json:"expose"`      //暴露端口 json存储
	InstallDate time.Time `json:"install_date" gorm:"autoCreateTime"`
	UpdateDate  time.Time `json:"update_date" gorm:"autoUpdateTime"`
}

package models

import (
	"github.com/magiconair/properties"
	"github.com/zapj/zap/core/base"
)

type ZapApps struct {
	Id         uint   `json:"id" gorm:"primarykey;autoIncrement"`
	AppStoreId uint   `json:"app_store_id" gorm:"index;"`
	Name       string `json:"name"`
	Title      string `json:"title"`
	Version    string `json:"version" gorm:"default:0.0.0"`
	AppType    string `json:"app_type"` // webserver , application , database

	Description string         `json:"description"`
	Status      string         `json:"status" gorm:"index;"` // install : 安装中  active : 已安装
	AppStatus   string         `json:"app_status"`           // running : 运行中  stop : 停止中
	InstallBy   string         `json:"install_by"`
	Instance    string         `json:"instance"`    // default
	InstallDir  string         `json:"install_dir"` //安装目录
	Settings    string         `json:"settings"`    // properties
	Expose      string         `json:"expose"`
	ConfigFile  string         `json:"config_file"` //配置文件
	PidFile     string         `json:"pid_file"`
	Error       string         `json:"error"` //错误信息
	InstallDate base.LocalTime `json:"install_date" gorm:"autoCreateTime"`
	UpdateDate  base.LocalTime `json:"update_date" gorm:"autoUpdateTime"`
}

func (a *ZapApps) ParseExpose() (*properties.Properties, error) {
	return properties.LoadString(a.Expose)

}

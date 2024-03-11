package appstore

import (
	"encoding/json"
	"log/slog"
	"time"

	"gorm.io/gorm"
)

type AppStoreVO struct {
	Id                     uint           `json:"id" gorm:"primarykey;autoIncrement"`
	AppId                  string         `json:"app_id"`
	Name                   string         `json:"app_name" gorm:"unique"`
	Title                  string         `json:"title"`
	Version                string         `json:"version"`
	Description            string         `json:"description"`
	ConfigName             string         `json:"config_name"`
	Category               string         `json:"category"`
	AllowMultipleInstances string         `json:"allow_multiple_instances" gorm:"default:no"`
	Icon                   string         `json:"icon"`
	Author                 string         `json:"author"`
	OrganizationName       string         `json:"organization_name"`
	Tags                   string         `json:"tags"`
	PublishAt              time.Time      `json:"publish_at"`
	Actions                string         `json:"-" gorm:"actions"`
	ActionsData            map[string]any `json:"actions" gorm:"-"`
	Options                string         `json:"options"`
	AlreadyInstalled       string         `json:"already_installed" gorm:"-"`
	AlreadyInstalledVers   string         `json:"already_installed_version" gorm:"-"`
	CreatedAt              time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt              time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
}

func (a *AppStoreVO) AfterFind(tx *gorm.DB) (err error) {
	slog.Info("after find", "row", a)
	if err = json.Unmarshal([]byte(a.Actions), &a.ActionsData); err != nil {
		return err
	}
	return nil
}

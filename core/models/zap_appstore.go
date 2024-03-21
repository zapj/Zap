package models

import (
	"encoding/json"
	"strings"
	"time"

	"gorm.io/gorm"
)

type ZapAppStore struct {
	Id    uint   `json:"id" gorm:"primarykey;autoIncrement"`
	AppId string `json:"app_id"`
	Name  string `json:"app_name" gorm:"unique"`
	Title string `json:"title"`

	Description            string    `json:"description"`
	ConfigName             string    `json:"config_name"`
	Category               string    `json:"category"`
	AllowMultipleInstances string    `json:"allow_multiple_instances" gorm:"default:no"`
	Icon                   string    `json:"icon"`
	Author                 string    `json:"author"`
	OrganizationName       string    `json:"organization_name"`
	PublishAt              time.Time `json:"publish_at"`
	Options                string    `json:"options"`

	Version string `json:"-"`
	Tags    string `json:"-"`
	Actions string `json:"-"`

	VersionData          []string       `json:"version" gorm:"-"`
	TagsData             []string       `json:"tags" gorm:"-"`
	ActionsData          map[string]any `json:"actions" gorm:"-"`
	AllowInstallation    string         `json:"allow_installation" gorm:"-"`
	AlreadyInstalledVers []string       `json:"already_installed_version" gorm:"-"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func (a *ZapAppStore) AfterFind(tx *gorm.DB) (err error) {
	if a.Actions != "" {
		err = json.Unmarshal([]byte(a.Actions), &a.ActionsData)
		if err != nil {
			return err
		}
	}
	if a.Tags != "" {
		a.TagsData = strings.Split(a.Tags, ",")
	}
	if a.Version != "" {
		a.VersionData = strings.Split(a.Version, ",")
	}
	return nil
}

func (a *ZapAppStore) TableName() string {
	return "zap_app_stores"
}

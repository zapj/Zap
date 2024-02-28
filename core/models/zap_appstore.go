package models

import "time"

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

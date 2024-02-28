package models

import (
	"gorm.io/gorm"
)

type ZapAccessLog struct {
	gorm.Model
	Username   string
	LogDate    string
	IpAddr     string
	UserAgent  string
	RequestUri string
}

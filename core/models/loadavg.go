package models

import (
	"gorm.io/gorm"
)

type LoadAvg struct {
	gorm.Model
	ID     int     `gorm:"primaryKey"`
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

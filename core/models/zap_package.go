package models

import "time"

type ZapPackage struct {
	Id              uint      `json:"id" gorm:"primarykey;autoIncrement"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	WebsiteCount    int       `json:"website_count"`    // 网站数量
	StorageSize     int       `json:"storage_size"`     //存储大小
	DiskSize        int       `json:"disk_size"`        //硬盘使用大小
	DatabaseSize    int       `json:"database_size"`    //数据库大小
	DatabaseCount   int       `json:"database_count"`   //数据库总数据量
	MySQLCount      int       `json:"mysql_count"`      //mysql 数据库
	PostgreSQLCount int       `json:"postgresql_count"` //postgresql 数据库
	FreeSSL         int       `json:"free_ssl"`         //免费ssl数量
	BandWidth       int       `json:"band_width"`       // 带宽限制
	DedicatedIP     string    `json:"dedicated_ip"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}

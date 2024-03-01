package models

import (
	"errors"

	"gorm.io/gorm"
)

type ZapUsers struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"-"`
	Uid      uint   `json:"uid"`
	Gid      uint   `json:"gid"`
	Gecos    string `json:"gecos"`
	Home     string `json:"home"`
	Shell    string `json:"shell"`

	// Days since Jan 1, 1970 password was last changed.
	LastChange string

	// -1 if password aging is disabled.
	MinPasswordAge string

	// 用户必须更改密码的天数。
	//
	// -1 is password aging is disabled.
	MaxPasswordAge string

	// 密码到期前的天数（请参阅上面的密码最长使用期限），在此期间应向用户发出警告。
	//
	// -1 is password aging is disabled.
	WarnPeriod string

	// 密码过期后的天数（请参阅上面的密码最长使用期限），在此期间密码仍应被接受。
	//
	// -1 is password aging is disabled.
	InactivityPeriod string

	// 帐户的到期日期，表示为自1970年1月1日以来的天数。
	//
	// -1 is account never expires.
	AccountExpiry string `json:"account_expiry"`

	// Unused now.
	Flags string `json:"flags"`

	//空间大小 kb
	DiskSpaceSize int `json:"disk_space_size"`
	DataBaseCount int `json:"database_count"`
	DataBaseSize  int `json:"database_size"`
}

func (u *ZapUsers) BeforeDelete(tx *gorm.DB) (err error) {
	if u.ID == 1 {
		return errors.New("admin user not allowed to delete")
	}
	return
}

package models

import "gorm.io/gorm"

type ZapUsers struct {
	gorm.Model
	Username string
	Password string `json:"-"`
	Uid      uint
	Gid      uint
	Gecos    string
	Home     string
	Shell    string

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
	AccountExpiry string

	// Unused now.
	Flags string
}

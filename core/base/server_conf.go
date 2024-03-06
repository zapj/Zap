package base

import (
	"path/filepath"
)

type ServerConf struct {
	SigningKey string
	EncryptKey string
	Host       string
	Port       int
	IPv6       bool
	CertFile   string
	KeyFile    string

	// 用户主目录
	WwwRoot    string
	WwwUser    string
	WwwGroup   string
	WwwUserId  int
	WwwGroupId int

	ZapMode string

	WebServer string
}

func (s *ServerConf) GetUserHomeDir(username string) string {
	return filepath.Join(s.WwwRoot, username)
}

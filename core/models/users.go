package models

type Zap_Users struct {
	Username string
	Password string
	Uid      int
	Gid      int

	//主目录
	Home string
}

package user_test

import (
	"testing"

	"github.com/zapj/zap/core/user"
)

func TestParsePasswd(t *testing.T) {
	users, err := user.ReadSystemUsers()
	if err != nil {
		t.Error(err)
	}
	t.Log(users["zap"].Password == "x")
}

func Test_Check_UserExists(t *testing.T) {

	t.Log(user.CheckUserExists("zap"))
}

func Test_ReadUsers(t *testing.T) {
	u := user.UserEntry{
		Username: "zap1",
		Password: "x",
		Uid:      "1000",
		Gid:      "1000",
		Home:     "/home/zap",
		Shell:    "/usr/bin/nologin",
		Gecos:    "",
	}
	t.Log(user.CreateUser(u))
}

func Test_ReadShadows(t *testing.T) {
	t.Log(user.ReadSystemShadow())
}

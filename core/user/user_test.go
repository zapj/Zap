package user_test

import (
	"testing"

	"github.com/zapj/zap/core/user"
)

func TestParsePasswd(t *testing.T) {
	users, err := user.Parse()
	if err != nil {
		t.Error(err)
	}
	t.Log(users["zap"].Password == "x")
}

func Test_Check_UserExists(t *testing.T) {

	t.Log(user.CheckUserExists("zap"))
}

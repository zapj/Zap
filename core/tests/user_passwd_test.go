package tests

import (
	"testing"

	"github.com/zapj/zap/core/utils/zaputil"
)

func TestGenPasswd(t *testing.T) {
	t.Log("GenPasswd")
	t.Log(zaputil.HashPassword("123456"))
}

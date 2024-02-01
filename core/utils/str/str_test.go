package str_test

import (
	"testing"

	"github.com/zapj/zap/core/utils/str"
)

func TestStrRand(t *testing.T) {
	t.Log(str.RandString(32))
	t.Log(str.RandUString(32))
}

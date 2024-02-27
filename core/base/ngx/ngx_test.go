package ngx_test

import (
	"testing"

	"github.com/zapj/zap/core/base/ngx"
)

func TestNgxConfServer(t *testing.T) {
	ngx := ngx.NewNgxConfServer("www.zap.cn", "zap.cn")
	ngx.AccessLog = "logs/access.log"
	ngx.ErrorLog = "logs/error.log"
	ngx.Root = "/home/www/admin/zap.cn"
	ngx.RunDirectory = "public"
	ngx.Listen = []string{"*:80"}

	t.Log(ngx.GenerateToServer())
}

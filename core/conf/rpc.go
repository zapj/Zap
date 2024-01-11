package conf

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/zapj/zap/core/global"
	"github.com/zapj/zap/core/zaprpc"
)

func RpcInit() {
	rpc.RegisterName("Command", new(zaprpc.Command))
	rpc.RegisterName("SystemUser", new(zaprpc.SystemUser))
	rpc.HandleHTTP()
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("RUNING"))
	})
	l, err := net.Listen("tcp", "127.0.0.1:2728")

	if err != nil {
		log.Fatal("listen error:", err)
	}
	global.LOG.Info("Server start : 127.0.0.1:2728")
	http.Serve(l, nil)
}

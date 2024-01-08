package tests

import (
	"log"
	"net/rpc"
	"testing"

	"github.com/zapj/zap/core/rpclient"
	"github.com/zapj/zap/core/zaprpc"
)

func connect() *rpc.Client {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2728")
	if err != nil {
		log.Fatal(err)
	}
	return client
}

func TestZapRPCCall(t *testing.T) {
	t.Log("call server")
	client := connect()
	defer client.Close()

	var res zaprpc.Response
	err := client.Call("Hello.Say", zaprpc.Args{"cmd": "test"}, &res)
	if err != nil {
		t.Log(err)
	}

	t.Log("rpc call result:", res)
}

func Test_ZapCommandCall(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2728")
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	var res zaprpc.Response
	err = client.Call("Command.Call", zaprpc.Args{"cmd": "who", "params": []string{"-b"}}, &res)
	if err != nil {
		t.Log(err)
	}

	t.Log("rpc call result:", res)
}

func Test_ZapCommandCallNotFound(t *testing.T) {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:2728")
	if err != nil {
		t.Error(err)
	}
	defer client.Close()
	var res zaprpc.Response
	err = client.Call("Command.Call", zaprpc.Args{"cmd": "lss", "params": []string{"-la"}}, &res)
	if err != nil {
		t.Logf("%T", err)
		t.Log(err)
	}

	t.Log("rpc call result:", res)
}

func Test_RPC_Client(t *testing.T) {
	t.Log("test RPC Client")
	client := rpclient.NewRpcClient()
	defer client.Close()
	var resp zaprpc.Response
	err := client.Call("Hello.Say", zaprpc.Args{"cmd": "test"}, &resp)
	if err != nil {
		t.Log("call error:", err)
	}
	t.Log("result:", resp)
}

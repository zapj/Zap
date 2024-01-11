package main

import (
	"fmt"

	"github.com/zapj/zap/core/rpclient"
	"github.com/zapj/zap/core/zaprpc"
)

func main() {

	client := rpclient.NewRpcClient()

	args := zaprpc.Args{}
	resp := zaprpc.Response{}
	var err error
	args["cmd"] = "./zap_cli"
	args["params"] = []string{"version"}
	// args["timeout"] = 3
	//
	err = client.Call("Command.CallContext", args, &resp)
	if err != nil {
		fmt.Println("EOF", err)
	}

	fmt.Println(resp.Data)

	// 测试user
	// err := client.Call("SystemUser.ReadShadow", args, &resp)
	// if err != nil {
	// 	fmt.Println("EOF", err)
	// }
	// var v map[string]user.ShadowEntry
	// fmt.Println(resp.GetJSON(&v))
	// fmt.Println(v["zap"])

}

func callCmd() {
	client := rpclient.NewRpcClient()

	args := zaprpc.Args{}
	resp := zaprpc.Response{}
	var err error

	args["cmd"] = "ls"
	err = client.Call("Command.Call", args, &resp)
	if err != nil {
		fmt.Println("EOF", err)
	}

	fmt.Println(resp.GetString())
}

func callTest() {
	client := rpclient.NewRpcClient()

	args := zaprpc.Args{}
	resp := zaprpc.Response{}

	err := client.Call("Command.Test", args, &resp)
	if err != nil {
		fmt.Println("EOF", err)
	}

	fmt.Println(resp.GetString())

}

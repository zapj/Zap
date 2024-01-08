package rpclient

import (
	"fmt"
	"net/rpc"

	"github.com/zapj/zap/assist/global"
)

type RpcClient struct {
	Host   string
	Port   int
	client *rpc.Client
}

func NewRpcClient() *RpcClient {
	return &RpcClient{
		Host: "127.0.0.1",
		Port: 2728,
	}
	//global.CONFIG.Section("assist").Key("port").MustInt(2728)
}

func (r *RpcClient) connect() {
	if r.client != nil {
		return
	}
	var err error
	r.client, err = rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", r.Host, r.Port))
	if err != nil {
		global.LOG.Errorf("connection to assist server failed: %v", err)
	}
}

func (r *RpcClient) Call(serviceMethod string, args any, reply any) error {
	r.connect()
	return r.client.Call(serviceMethod, args, reply)
}

func (r *RpcClient) Go(serviceMethod string, args any, reply any, done chan *rpc.Call) *rpc.Call {
	r.connect()
	return r.client.Go(serviceMethod, args, reply, done)
}

func (r *RpcClient) Close() {
	if r.client != nil {
		r.client.Close()
	}
}

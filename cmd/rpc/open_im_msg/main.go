package main

import (
	rpcChat "Open_IM/internal/rpc/msg"
	"Open_IM/pkg/common/config"
	"flag"
	"fmt"
)

func main() {
	defaultPorts := config.Config.RpcPort.OpenImMessagePort
	rpcPort := flag.Int("port", defaultPorts[0], "rpc listening port")
	flag.Parse()
	fmt.Println("start msg rpc server, port: ", *rpcPort)
	rpcServer := rpcChat.NewRpcChatServer(*rpcPort)
	rpcServer.Run()
}

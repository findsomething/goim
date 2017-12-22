package main

import (
	"goim/libs/net/xrpc"
	"goim/libs/hash/ketama"
	"strings"

	inet "goim/libs/net"
	"qiniupkg.com/x/log.v7"
)

var (
	logicServiceMap = map[string]*xrpc.Clients{}
	logicRing       *ketama.HashRing
)

const (
	logicServicePing = "LogicRPC.Ping"
)

func InitLogic(addrs map[string]string) (err error) {
	var (
		network, addr string
	)
	logicRing = ketama.NewRing(ketama.Base)
	for serverId, bind := range addrs {
		var rpcOptions []xrpc.ClientOptions
		for _, bind = range strings.Split(bind, ",") {
			if network, addr, err = inet.ParseNetwork(bind); err != nil {
				log.Error("inet.ParseNetwork() error(%v)", err)
				return
			}
			options := xrpc.ClientOptions{
				Proto: network,
				Addr:  addr,
			}
			rpcOptions = append(rpcOptions, options)
		}
		rpcClient := xrpc.Dials(rpcOptions)

		rpcClient.Ping(logicServicePing)
		logicRing.AddNode(serverId, 1)
		logicServiceMap[serverId] = rpcClient
		log.Info("router rpc connect: %v", rpcOptions)
	}
	logicRing.Bake()
	return
}

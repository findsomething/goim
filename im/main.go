package main

import (
	"flag"
	"runtime"

	log "github.com/thinkboy/log4go"
	//"goim/libs/perf"
	"fmt"
	"goim/im/tool"
)

var Conf *tool.Config

func main() {
	flag.Parse()
	var err error
	if Conf, err = tool.InitConfig(""); err != nil {
		panic(err)
	}
	fmt.Println(Conf.RPCConf)
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.LoadConfiguration(Conf.Log)
	defer log.Close()
	log.Info("im[%s] start", Ver)

	//perf.Init(Conf.PprofAddrs)
	//// logic rpc
	//if err := InitLogic(Conf.LogicRpcAddrs); err != nil {
	//	log.Warn("logic rpc current can't connect, retry")
	//}
	//if Conf.MonitorOpen {
	//	InitMonitor(Conf.MonitorAddrs)
	//}

	// im rpc
	if err := InitRpc(Conf.RPCConf); err != nil {
		panic(err)
	}
	if err := InitHTTP(); err != nil {
		panic(err)
	}
	InitSignal()
}

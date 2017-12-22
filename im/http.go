package main

import (
	"net/http"
	log "github.com/thinkboy/log4go"
	inet "goim/libs/net"
	"net"
)

func InitHTTP() (err error) {
	var network, addr string
	for i := 0; i < len(Conf.HTTPAddrs); i++ {
		httpServerMux := http.NewServeMux()
		log.Info("start http listen:\"%s\"", Conf.HTTPAddrs[i]);
		if network, addr, err = inet.ParseNetwork(Conf.HTTPAddrs[i]); err != nil {
			log.Error("inet.ParseNetwork() error(%v)", err)
		}
		go httpListen(httpServerMux, network, addr)
	}
	return
}

func httpListen(mux *http.ServeMux, network, addr string) {
	httpServer := &http.Server{Handler: mux, ReadTimeout: Conf.HTTPReadTimeout, WriteTimeout: Conf.HTTPWriteTimeout}
	httpServer.SetKeepAlivesEnabled(true)
	l, err := net.Listen(network, addr)
	if err != nil {
		log.Error("net.Listen(\"%s\", \"%s\") error(%v)", network, addr, err)
		panic(err)
	}
	if err := httpServer.Serve(l); err != nil {
		log.Error("server.Serve() error(%v)", err)
		panic(err)
	}
}

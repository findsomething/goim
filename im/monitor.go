package main

import (
	"net/http"
	"fmt"
	"qiniupkg.com/x/log.v7"
)

type Monitor struct {
}

func InitMonitor(binds []string) {
	m := new(Monitor)
	monitorServeMux := http.NewServeMux()
	monitorServeMux.HandleFunc("/monitor/ping", m.Ping)
	for _, addr := range binds {
		go func(bind string) {
			if err := http.ListenAndServe(bind, monitorServeMux); err != nil {
				log.Error("http.ListenAndServe(\"%s\", pprofServeMux error(%v)", addr, err)
				panic(err)
			}
		}(addr)
	}
}

func (m *Monitor) Ping(w http.ResponseWriter, r *http.Request) {
	for _, c := range logicServiceMap {
		if err := c.Available(); err != nil {
			http.Error(w, fmt.Sprintf("ping rpc error(%v)", err), http.StatusInternalServerError)
			return
		}
	}
	w.Write([]byte("ok"))
}

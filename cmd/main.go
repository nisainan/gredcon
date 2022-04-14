package main

import (
	"flag"
	"github.com/panjf2000/gnet"
	"redcon/handler"
	"redcon/redcon"
)

func main() {
	var port int
	var multicore bool
	flag.IntVar(&port, "port", 6380, "server port")
	flag.BoolVar(&multicore, "multicore", false, "multicore")
	flag.Parse()
	options := redcon.Options{
		Options: gnet.Options{
			Multicore: multicore,
		},
		Port: port,
	}
	handle := handler.NewHandler()
	mux := redcon.NewServeMux()
	mux.HandleFunc("ping", handle.Ping)
	mux.HandleFunc("set", handle.Set)
	mux.HandleFunc("get", handle.Get)
	redcon.ListenAndServe(options, mux.ServeRESP)
}

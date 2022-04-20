package main

import (
	"flag"
	"github.com/nisainan/redcon-gnet/handler"
	"github.com/nisainan/redcon-gnet/redcon"
	"github.com/panjf2000/gnet"
)

func main() {
	var port int
	var multicore bool
	var reuseport bool
	flag.IntVar(&port, "port", 6379, "server port")
	flag.BoolVar(&multicore, "multicore", false, "multicore")
	flag.BoolVar(&reuseport, "reuseport", false, "reuseport")
	flag.Parse()
	options := redcon.Options{
		Options: gnet.Options{
			Multicore: multicore,
			ReusePort: reuseport,
		},
		Port: port,
	}
	handle := handler.NewHandler()
	mux := redcon.NewServeMux()
	mux.HandleFunc("ping", handle.Ping)
	mux.HandleFunc("set", handle.Set)
	mux.HandleFunc("get", handle.Get)
	mux.HandleFunc("echo", handle.Echo)
	redcon.ListenAndServe(options, mux.ServeRESP)
}

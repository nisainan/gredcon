package main

import (
	"flag"
	"github.com/panjf2000/gnet"
	"log"
	"net/http"
	_ "net/http/pprof"
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

	go func() {
		var err error
		var pprofUri = "0.0.0.0:8080"
		defer func() {
			if err != nil {
				log.Printf("pprof listen error: %s\n", err.Error())
			}
		}()
		log.Printf("pprof listen on: %s\n", pprofUri)
		err = http.ListenAndServe(pprofUri, nil)
	}()

	handle := handler.NewHandler()
	mux := redcon.NewServeMux()
	mux.HandleFunc("ping", handle.Ping)
	mux.HandleFunc("set", handle.Set)
	mux.HandleFunc("get", handle.Get)
	redcon.ListenAndServe(options, mux.ServeRESP)
}

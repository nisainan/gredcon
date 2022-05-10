package main

import (
	"flag"
	"github.com/nisainan/gredcon/gredcon"
	"github.com/panjf2000/gnet"
	"sync"
)

func main() {
	var port int
	var multicore bool
	var reuseport bool
	flag.IntVar(&port, "port", 6379, "server port")
	flag.BoolVar(&multicore, "multicore", false, "multicore")
	flag.BoolVar(&reuseport, "reuseport", false, "reuseport")
	flag.Parse()
	options := gredcon.Options{
		Options: gnet.Options{
			Multicore: multicore,
			ReusePort: reuseport,
		},
		Port: port,
	}
	handle := NewHandler()
	mux := gredcon.NewServeMux()
	mux.HandleFunc("ping", handle.Ping)
	mux.HandleFunc("set", handle.Set)
	mux.HandleFunc("get", handle.Get)
	mux.HandleFunc("echo", handle.Echo)
	gredcon.ListenAndServe(options, mux.ServeRESP)
}

type Handler struct {
	itemsMux sync.RWMutex
	items    map[string][]byte
}

func NewHandler() *Handler {
	return &Handler{
		items: make(map[string][]byte),
	}
}

func (h *Handler) Ping(conn gredcon.Conn, cmd gredcon.Command) {
	conn.WriteString("PONG")
}

func (h *Handler) Set(conn gredcon.Conn, cmd gredcon.Command) {
	if len(cmd.Args) != 3 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return
	}

	h.itemsMux.Lock()
	h.items[string(cmd.Args[1])] = cmd.Args[2]
	h.itemsMux.Unlock()

	conn.WriteString("OK")
}

func (h *Handler) Get(conn gredcon.Conn, cmd gredcon.Command) {
	if len(cmd.Args) != 2 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return
	}

	h.itemsMux.RLock()
	val, ok := h.items[string(cmd.Args[1])]
	h.itemsMux.RUnlock()

	if !ok {
		conn.WriteNull()
	} else {
		conn.WriteBulk(val)
	}
}

func (h *Handler) Echo(conn gredcon.Conn, cmd gredcon.Command) {
	if len(cmd.Args) != 2 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return
	}
	conn.WriteBulk(cmd.Args[1])
}

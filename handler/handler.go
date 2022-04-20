package handler

import (
	"github.com/nisainan/redcon-gnet/redcon"
	"sync"
)

type Handler struct {
	itemsMux sync.RWMutex
	items    map[string][]byte
}

func NewHandler() *Handler {
	return &Handler{
		items: make(map[string][]byte),
	}
}

func (h *Handler) Ping(conn redcon.Conn, cmd redcon.Command) {
	conn.WriteString("PONG")
}

func (h *Handler) Set(conn redcon.Conn, cmd redcon.Command) {
	if len(cmd.Args) != 3 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return
	}

	h.itemsMux.Lock()
	h.items[string(cmd.Args[1])] = cmd.Args[2]
	h.itemsMux.Unlock()

	conn.WriteString("OK")
}

func (h *Handler) Get(conn redcon.Conn, cmd redcon.Command) {
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

func (h *Handler) Echo(conn redcon.Conn, cmd redcon.Command) {
	if len(cmd.Args) != 2 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Args[0]) + "' command")
		return
	}
	conn.WriteBulk(cmd.Args[1])
}

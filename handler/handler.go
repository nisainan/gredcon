package handler

import (
	"redcon/pkg/resparse/credis"
	"redcon/redcon"
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

func (h *Handler) Ping(conn redcon.Conn, cmd *credis.Resp) {
	conn.WriteString("PONG")
}

func (h *Handler) Set(conn redcon.Conn, cmd *credis.Resp) {
	if len(cmd.Array) != 3 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Value) + "' command")
		return
	}

	h.itemsMux.Lock()
	h.items[string(cmd.Array[1].Value)] = cmd.Array[2].Value
	h.itemsMux.Unlock()

	conn.WriteString("OK")
}

func (h *Handler) Get(conn redcon.Conn, cmd *credis.Resp) {
	if len(cmd.Array) != 2 {
		conn.WriteError("ERR wrong number of arguments for '" + string(cmd.Value) + "' command")
		return
	}

	h.itemsMux.RLock()
	val, ok := h.items[string(cmd.Array[1].Value)]
	h.itemsMux.RUnlock()

	if !ok {
		conn.WriteNull()
	} else {
		conn.WriteBulk(val)
	}
}

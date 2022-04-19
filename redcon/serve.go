package redcon

import (
	"redcon/pkg/resparse/credis"
	"strings"
)

type ServeMux struct {
	handlers map[string]Handler
}

func NewServeMux() *ServeMux {
	return &ServeMux{
		handlers: make(map[string]Handler),
	}
}

func (m *ServeMux) HandleFunc(command string, handler func(conn Conn, resp *credis.Resp)) {
	if handler == nil {
		panic("redcon: nil handler")
	}
	m.Handle(command, HandlerFunc(handler))
}

func (m *ServeMux) Handle(command string, handler Handler) {
	if command == "" {
		panic("redcon: invalid command")
	}
	if handler == nil {
		panic("redcon: nil handler")
	}
	if _, exist := m.handlers[command]; exist {
		panic("redcon: multiple registrations for " + command)
	}

	m.handlers[command] = handler
}

func (m *ServeMux) ServeRESP(conn Conn, resp *credis.Resp) {
	command := strings.ToLower(string(resp.Array[0].Value))

	if handler, ok := m.handlers[command]; ok {
		handler.ServeRESP(conn, resp)
	} else {
		conn.WriteError("ERR unknown command '" + command + "'")
	}
}

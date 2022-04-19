package redcon

import "redcon/pkg/resparse/credis"

type Handler interface {
	ServeRESP(conn Conn, resp *credis.Resp)
}

type HandlerFunc func(conn Conn, resp *credis.Resp)

func (f HandlerFunc) ServeRESP(conn Conn, resp *credis.Resp) {
	f(conn, resp)
}

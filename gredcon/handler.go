package gredcon

type Handler interface {
	ServeRESP(conn Conn, cmd Command)
}

type HandlerFunc func(conn Conn, cmd Command)

func (f HandlerFunc) ServeRESP(conn Conn, cmd Command) {
	f(conn, cmd)
}

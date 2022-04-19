package redcon

import (
	"fmt"
	"github.com/panjf2000/gnet"
	"log"
	"net"
	"redcon/pkg/resparse/credis"
	"sync"
	"time"
)

type Server struct {
	*gnet.EventServer
	//eng       gnet.Engine
	mu        sync.RWMutex
	net       string
	laddr     string
	handler   func(conn Conn, resp *credis.Resp)
	accept    func(conn Conn) bool
	closed    func(conn Conn, err error)
	conns     map[gnet.Conn]*RedCon
	ln        net.Listener
	done      bool
	idleClose time.Duration

	// AcceptError is an optional function used to handle Accept errors.
	AcceptError func(err error)
	multicore   bool
}

func (s *Server) OnInitComplete(srv gnet.Server) (action gnet.Action) {
	log.Printf("server with multi-core=%t is listening on %s\n", s.multicore, s.laddr)
	return
}

func (s *Server) OnShutdown(svr gnet.Server) {
}

func (s *Server) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.conns[c] = NewRedcon()
	return
}

func (s *Server) OnClosed(c gnet.Conn, err error) (action gnet.Action) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	delete(s.conns, c)
	return gnet.None
}

func (s *Server) React(frame []byte, conn gnet.Conn) (out []byte, action gnet.Action) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	c := s.conns[conn]
	c.rd.Write(frame)
	decoder := credis.NewDecoderSize(c.rd, 1024)
	//cmds, lastbyte, err := ReadCommands(c.rd.Bytes())
	//for {
	resp, _ := decoder.Decode()
	//if err != nil {
	//	c.wr.WriteError("ERR " + err.Error())
	//	return c.wr.Buffer(), gnet.None
	//}
	//if err != nil || resp == nil {
	//	break
	//}
	c.rd.resps = append(c.rd.resps, resp)
	//}
	//c.rd.Reset()
	//if len(lastbyte) > 0 {
	//	c.rd.Write(lastbyte)
	//} else {
	//	for len(c.rd.resps) > 0 {
	//		resp := c.rd.resps[0]
	//		if len(c.rd.resps) == 1 {
	//			c.rd.resps = nil
	//		} else {
	//			c.rd.resps = c.rd.resps[1:]
	//		}
	//		s.handler(c, resp)
	//	}
	//}
	for len(c.rd.resps) > 0 {
		resp := c.rd.resps[0]
		if len(c.rd.resps) == 1 {
			c.rd.resps = nil
		} else {
			c.rd.resps = c.rd.resps[1:]
		}
		s.handler(c, resp)
	}
	if len(c.wr.Buffer()) > 0 {
		defer c.wr.Flush()
		return c.wr.Buffer(), gnet.None
	}
	return
}

func (s *Server) OnTick() (delay time.Duration, action gnet.Action) {
	log.Println("OnTick")
	return
}

func ListenAndServe(options Options, handler func(conn Conn, resp *credis.Resp)) {
	server := &Server{
		laddr:     fmt.Sprintf("tcp://:%d", options.Port),
		handler:   handler,
		conns:     make(map[gnet.Conn]*RedCon, 0),
		multicore: options.Multicore,
	}
	log.Fatal(gnet.Serve(server, server.laddr, gnet.WithOptions(options.Options)))
}

package redcon

import (
	"github.com/panjf2000/gnet"
	"time"
)

type Options struct {
	gnet.Options
	Port int
}

type RedCon struct {
	wr       *Writer
	rd       *Reader
	ctx      interface{}
	detached bool
	closed   bool
	//cmds      []Command
	idleClose time.Duration
}

func (c *RedCon) WriteError(msg string)       { c.wr.WriteError(msg) }
func (c *RedCon) WriteString(str string)      { c.wr.WriteString(str) }
func (c *RedCon) WriteBulk(bulk []byte)       { c.wr.WriteBulk(bulk) }
func (c *RedCon) WriteBulkString(bulk string) { c.wr.WriteBulkString(bulk) }
func (c *RedCon) WriteInt(num int)            { c.wr.WriteInt(num) }
func (c *RedCon) WriteInt64(num int64)        { c.wr.WriteInt64(num) }
func (c *RedCon) WriteUint64(num uint64)      { c.wr.WriteUint64(num) }
func (c *RedCon) WriteArray(count int)        { c.wr.WriteArray(count) }
func (c *RedCon) WriteNull()                  { c.wr.WriteNull() }
func (c *RedCon) WriteRaw(data []byte)        { c.wr.WriteRaw(data) }
func (c *RedCon) WriteAny(v interface{})      { c.wr.WriteAny(v) }
func (c *RedCon) Context() interface{}        { return c.ctx }
func (c *RedCon) SetContext(v interface{})    { c.ctx = v }
func (c *RedCon) SetReadBuffer(n int)         {}
func (c *RedCon) ReadPipeline() []Command {
	cmds := c.rd.cmds
	c.rd.cmds = nil
	return cmds
}
func (c *RedCon) PeekPipeline() []Command {
	return c.rd.cmds
}

func NewRedcon() *RedCon {
	return &RedCon{
		wr: NewWriter(),
		rd: NewReader(),
	}
}

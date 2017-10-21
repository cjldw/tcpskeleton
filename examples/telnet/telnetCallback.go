package telnet

import (
	"fmt"
	"github.com/vvotm/tcpskeleton"
)

type TelnetCallback struct {
}

func (this *TelnetCallback) OnConnect(c *tcpskeleton.Conn) bool {
	addr := c.GetRawConn().RemoteAddr()
	c.PutExtraData("remoteAddr", addr)
	fmt.Println("OnConnect:", addr)
	c.AsyncWritePacket(NewTelnetPacket("unknow", []byte("Welcome to this Telnet Server")))
	return true
}

func (this *TelnetCallback) OnMessage(c *tcpskeleton.Conn, p tcpskeleton.Packet) bool {
	packet := p.(*TelnetPacket)
	command := packet.GetData()
	commandType := packet.GetType()
	switch commandType {
	case "echo":
		c.AsyncWritePacket(NewTelnetPacket("echo", command))
	case "login":
		c.AsyncWritePacket(NewTelnetPacket("login", []byte(string(command)+" has login")))
	case "quit":
		return false
	default:
		c.AsyncWritePacket(NewTelnetPacket("unknow", []byte("unknow command")))
	}
	return true
}

func (this *TelnetCallback) OnClose(c *tcpskeleton.Conn) {
	fmt.Println("OnClose:", c.GetExtraData("remoteAddr"))
}


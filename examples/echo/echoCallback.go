package echo

import (
	"github.com/vvotm/tcpskeleton"
	"fmt"
)

type Callback struct{}

func (this *Callback) OnConnect(c *tcpskeleton.Conn) bool {
	addr := c.GetRawConn().RemoteAddr()
	c.PutExtraData("remoteAddr", addr)
	fmt.Println("OnConnect:", addr)
	return true
}

func (this *Callback) OnMessage(c *tcpskeleton.Conn, p tcpskeleton.Packet) bool {
	echoPacket := p.(*EchoPacket) // accept client send packet
	fmt.Printf("OnMessage:[%v] [%v]\n", echoPacket.GetLength(), string(echoPacket.GetBody()))
	// bala bala ellipsis many business code..
	c.AsyncWritePacket(echoPacket) // send packet back to client
	return true
}

func (this *Callback) OnClose(c *tcpskeleton.Conn) {
	fmt.Println("OnClose:", c.GetExtraData("remoteAddr"))
}

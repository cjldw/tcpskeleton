package diyproto

import (
	"github.com/vvotm/tcpskeleton"
	"fmt"
	"encoding/json"
)

type DiyCallback struct {

}

func (d DiyCallback) OnConnect(conn *tcpskeleton.Conn) bool  {
	fmt.Printf("OnConnect: %s\n", conn.GetRawConn().RemoteAddr().String())
	conn.PutExtraData("remoteAddr", conn.GetRawConn().RemoteAddr().String())
	return true
}

func (d DiyCallback) OnMessage(conn *tcpskeleton.Conn, packet tcpskeleton.Packet) bool {
	diyPacket := packet.(DiyPacket)
	param := make(map[string]interface{})
	if err := json.Unmarshal(diyPacket.Body, &param); err != nil {
		fmt.Println(err)
		return false
	}
	cmdid, ok := param["cmdid"]
	if  !ok {
		conn.AsyncWritePacket(NewDiyPacket("UnKnow Protocol"))
		return true
	}
	fmt.Printf("receive data: %v\n", param)
	conn.AsyncWritePacket(NewDiyPacket(cmdid.(string)))
	return true
}

func (d DiyCallback) OnClose(conn *tcpskeleton.Conn)  {
	fmt.Println("Close:", conn.GetExtraData("remoteAddr"))
	conn.Close()
}

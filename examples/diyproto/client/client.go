package main

import (
	"net"
	"github.com/vvotm/tcpskeleton/examples/diyproto"
	"fmt"
)

func main() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:19898")
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	packet := diyproto.NewDiyPacket("{\"cmdid\":\"hello\"}")
	fmt.Println(packet)
	conn.Write(packet.Serialize())
	receive := make([]byte, 1024)
	length, err := conn.Read(receive)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("发送的协议", string(receive[4:length]))
}

package main

import (
	"fmt"
	"log"
	"net"
	"time"
	"github.com/vvotm/tcpskeleton/examples/echo"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:19898")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	echoProtocol := &echo.EchoProtocol{}
	// ping <--> pong
	for i := 0; i < 3; i++ {
		// write
		conn.Write(echo.NewEchoPacket([]byte("hello"), false).Serialize())
		fmt.Println("ok")
		// read
		p, err := echoProtocol.ReadPacket(conn)
		if err == nil {
			echoPacket := p.(*echo.EchoPacket)
			fmt.Printf("Server reply:[%v] [%v]\n", echoPacket.GetLength(), string(echoPacket.GetBody()))
		}
		time.Sleep(1 * time.Second)
	}
	conn.Close()
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

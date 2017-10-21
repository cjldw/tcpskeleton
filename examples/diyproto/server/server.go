package main

import "github.com/vvotm/tcpskeleton"

func main() {
	conf := tcpskeleton.Config{
		Debug: true,
		Network: "tcp",
		ListenAddr: "127.0.0.1:19898",
		TcpAcceptTimeout: 20,
		TcpPacketWriteTimeout: 10,
		PacketReceiveChanLimit: 20,
		PacketSendChanLimit: 20,
	}
}

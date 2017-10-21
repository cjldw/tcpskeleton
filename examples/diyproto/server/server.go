package main

import (
	"github.com/vvotm/tcpskeleton"
	"github.com/vvotm/tcpskeleton/examples/diyproto"
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	conf := &tcpskeleton.Config{
		Debug: true,
		Network: "tcp",
		ListenAddr: "127.0.0.1:19898",
		TcpAcceptTimeout: 8,
		TcpPacketWriteTimeout: 10,
		PacketReceiveChanLimit: 20,
		PacketSendChanLimit: 20,
	}

	srv := tcpskeleton.NewServer(conf, diyproto.DiyCallback{}, diyproto.DiyProtocol{})
	go srv.Start()
	fmt.Println("listen:", conf.ListenAddr)
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGKILL)
	fmt.Printf("signal: %v\n", <-signalChan)
	srv.Stop()
}



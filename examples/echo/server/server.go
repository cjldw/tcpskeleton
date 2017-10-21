package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"github.com/vvotm/tcpskeleton/examples/echo"
	"github.com/vvotm/tcpskeleton"
)


func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// creates a server
	config := &tcpskeleton.Config{
		ListenAddr: "127.0.0.1:19898",
		Network: "tcp",
		TcpAcceptTimeout: 10, // set not limit tcp accept time
		TcpPacketWriteTimeout: 5, // set max time of write tcp packet
		PacketSendChanLimit: 20,
		PacketReceiveChanLimit: 20,
	}
	srv := tcpskeleton.NewServer(config, &echo.Callback{}, &echo.EchoProtocol{})
	// starts service
	go srv.Start()
	fmt.Println("listening:", config.ListenAddr)
	// catchs system signal
	chSig := make(chan os.Signal)
	signal.Notify(chSig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Signal: ", <-chSig)
	// stops service
	srv.Stop()
}


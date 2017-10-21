package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"github.com/vvotm/tcpskeleton"
	"github.com/vvotm/tcpskeleton/examples/telnet"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	config := &tcpskeleton.Config{
		ListenAddr: "127.0.0.1:19898",
		TcpAcceptTimeout: 10,
		TcpPacketWriteTimeout: 3,
		PacketSendChanLimit:    20,
		PacketReceiveChanLimit: 20,
	}
	srv := tcpskeleton.NewServer(config, &telnet.TelnetCallback{}, &telnet.TelnetProtocol{})
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

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

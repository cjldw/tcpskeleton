package tcpskeleton

import (
	"net"
	"sync"
	"fmt"
	"time"
)

type Config struct {
	ListenAddr             string // listen address
	TcpAcceptTimeout       int    // tcp max acceptTimeout if set zero accept will not timeout
	TcpPacketWriteTimeout  int    // write tcp packet timeout
	PacketSendChanLimit    uint32 // the limit of packet send channel
	PacketReceiveChanLimit uint32 // the limit of packet receive channel
}

type Server struct {
	config    *Config         // server configuration
	callback  ConnCallback    // message callbacks in connection
	protocol  Protocol        // customize packet protocol
	exitChan  chan struct{}   // notify all goroutines to shutdown
	waitGroup *sync.WaitGroup // wait for all goroutines
}

// NewServer creates a server
func NewServer(config *Config, callback ConnCallback, protocol Protocol) *Server {
	return &Server{
		config:    config,
		callback:  callback,
		protocol:  protocol,
		exitChan:  make(chan struct{}),
		waitGroup: &sync.WaitGroup{},
	}
}

// Start starts service
func (s *Server) Start() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", s.config.ListenAddr)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	ThrowErr(err)
	acceptTimeout := time.Duration(s.config.TcpAcceptTimeout) * time.Second
	fmt.Println(acceptTimeout)
	s.waitGroup.Add(1)
	defer func() {
		listener.Close()
		s.waitGroup.Done()
	}()

	for {
		select {
		case <-s.exitChan:
			fmt.Println("exit")
			return
		default:
		}
		listener.SetDeadline(time.Now().Add(acceptTimeout))
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		s.waitGroup.Add(1)
		go func() {
			newConn(conn, s).Do()
			s.waitGroup.Done()
		}()
	}
	fmt.Println("exit2")
}

// Stop stops service
func (s *Server) Stop() {
	close(s.exitChan)
	fmt.Println("close signal")
	s.waitGroup.Wait()
}

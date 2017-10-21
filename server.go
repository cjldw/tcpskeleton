package tcpskeleton

import (
	"net"
	"sync"
	"fmt"
	"time"
	"log"
)

type Config struct {
	Debug bool // Debug or not
	Network string // network "tcp", "tcp4", "udp", "udp4" and so on
	ListenAddr string // listen address
	TcpAcceptTimeout       int    // tcp max acceptTimeout if set zero accept will not timeout
	TcpPacketWriteTimeout  int    // write tcp packet timeout
	PacketSendChanLimit    uint32 // the limit of packet send channel
	PacketReceiveChanLimit uint32 // the limit of packet receive channel
}

type Server struct {
	debug bool
	config    *Config         // server configuration
	callback  ConnCallback    // message callbacks in connection
	protocol  Protocol        // customize packet protocol
	exitChan  chan struct{}   // notify all goroutines to shutdown
	waitGroup *sync.WaitGroup // wait for all goroutines
}

// NewServer creates a server
func NewServer(config *Config, callback ConnCallback, protocol Protocol) *Server {
	return &Server{
		debug: config.Debug,
		config:    config,
		callback:  callback,
		protocol:  protocol,
		exitChan:  make(chan struct{}),
		waitGroup: &sync.WaitGroup{},
	}
}

// Start starts service
func (s *Server) Start() {
	network := "tcp"
	if s.config.Network != "" {
		network = s.config.Network
	}
	tcpAddr, _ := net.ResolveTCPAddr(network, s.config.ListenAddr)
	listener, err := net.ListenTCP(network, tcpAddr)
	ThrowErr(err)
	acceptTimeout := time.Duration(s.config.TcpAcceptTimeout) * time.Second
	s.waitGroup.Add(1)
	defer func() {
		listener.Close()
		s.waitGroup.Done()
	}()
	for {
		select {
		case <-s.exitChan:
			return
		default:
		}
		listener.SetDeadline(time.Now().Add(acceptTimeout))
		conn, err := listener.AcceptTCP()
		if err != nil {
			s.Trace(err)
			continue
		}
		s.waitGroup.Add(1)
		go func() {
			newConn(conn, s).Do()
			s.waitGroup.Done()
		}()
	}
}

// Stop stops service
func (s *Server) Stop() {
	close(s.exitChan)
	s.waitGroup.Wait()
}

func (s *Server) Trace(trace interface{}) {
	if s.debug {
		log.Printf("Trace info: %v\n", trace)
	}
}

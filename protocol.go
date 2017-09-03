package gotcp

import (
	"net"
)

// Packet
type Packet interface {
	Serialize() []byte
}

// Protocol
type Protocol interface {
	ReadPacket(conn *net.TCPConn) (Packet, error)
}

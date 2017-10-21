package tcpskeleton

import "net"

// Protocol
type Protocol interface {
	ReadPacket(conn *net.TCPConn) (Packet, error)
}

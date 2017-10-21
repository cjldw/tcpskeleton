package echo

import (
	"encoding/binary"
	"errors"
	"io"
	"net"
	"github.com/vvotm/tcpskeleton"
)
type EchoProtocol struct {
}

func (this *EchoProtocol) ReadPacket(conn *net.TCPConn) (tcpskeleton.Packet, error) {
	var (
		lengthBytes []byte = make([]byte, 4)
		length      uint32
	)
	// read length
	if _, err := io.ReadFull(conn, lengthBytes); err != nil {
		return nil, err
	}
	const ReceiveMaxLen  = 1024 // max tcp packet length
	if length = binary.BigEndian.Uint32(lengthBytes); length > ReceiveMaxLen {
		return nil, errors.New("the size of packet is larger than the limit")
	}
	buff := make([]byte, 4+length)
	copy(buff[0:4], lengthBytes)
	// read body ( buff = lengthBytes + body )
	if _, err := io.ReadFull(conn, buff[4:]); err != nil {
		return nil, err
	}
	return NewEchoPacket(buff, true), nil
}

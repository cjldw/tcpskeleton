package diyproto

import (
	"net"
	"encoding/binary"
	"errors"
	"io"
	"github.com/vvotm/tcpskeleton"
)

// MAX_RECEIVE_LENGTH max length of receive packet length
const MAX_RECEIVE_LENGTH  =  1024

type DiyProtocol struct {

}
// ReadPacket read packet from connection
func (d DiyProtocol) ReadPacket(conn *net.TCPConn) (tcpskeleton.Packet, error)  {
	var (
		buf []byte
		bodyLen uint32
		emptyPacket DiyPacket = DiyPacket{}
	)
	buf = make([]byte, 4)
	if _, err := conn.Read(buf); err != nil {
		return emptyPacket, err
	}
	if bodyLen = binary.LittleEndian.Uint32(buf); bodyLen > MAX_RECEIVE_LENGTH {
		return emptyPacket, errors.New("body length max than 1024")
	}
	bodyByte := make([]byte, bodyLen)
	if _, err := io.ReadFull(conn, bodyByte); err != nil {
		return emptyPacket, err
	}

	return DiyPacket{
		HeadLen: bodyLen,
		Body: bodyByte,
	}, nil
}

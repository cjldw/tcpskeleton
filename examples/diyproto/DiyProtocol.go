package diyproto

import (
	"net"
	"encoding/binary"
	"errors"
	"io"
)

// MAX_RECEIVE_LENGTH max length of receive packet length
const MAX_RECEIVE_LENGTH  =  1024

type DiyProtocol struct {

}
// ReadPacket read packet from connection
func (d *DiyProtocol) ReadPacket(conn *net.TCPConn) (DiyPacket, error)  {
	var (
		buf []byte
		bodyLen uint32
	)
	buf = make([]byte, 4)
	if _, err := conn.Read(buf); err != nil {
		return nil, err
	}
	if bodyLen = binary.LittleEndian.Uint32(buf); bodyLen > MAX_RECEIVE_LENGTH {
		return nil, errors.New("body length max than 1024")
	}
	bodyByte := make([]byte, bodyLen)
	if _, err := io.ReadFull(conn, bodyByte); err != nil {
		return nil, err
	}

	return DiyPacket{
		HeadLen: bodyLen,
		Body: bodyByte,
	}
}

package diyproto

import "encoding/binary"

type DiyPacket struct {
	HeadLen uint32
	Body []byte
}

func (dp DiyPacket) Serialize() []byte {
	buf := make([]byte, 4 + uint32(len(dp.Body)))
	binary.LittleEndian.PutUint32(buf[:4], dp.HeadLen)
	copy(buf[4:], dp.Body)
	return buf
}

func NewDiyPacket(packet string) DiyPacket {
	return DiyPacket{
		HeadLen: uint32(len(packet)),
		Body: []byte(packet),
	}
}

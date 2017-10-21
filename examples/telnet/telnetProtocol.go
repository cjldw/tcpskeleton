package telnet

import (
	"bytes"
	"net"
	"strings"
	"github.com/vvotm/tcpskeleton"
)

type TelnetProtocol struct {
}

func (this *TelnetProtocol) ReadPacket(conn *net.TCPConn) (tcpskeleton.Packet, error) {
	fullBuf := bytes.NewBuffer([]byte{})
	for {
		data := make([]byte, 1024)

		readLengh, err := conn.Read(data)

		if err != nil { //EOF, or worse
			return nil, err
		}

		if readLengh == 0 { // Connection maybe closed by the client
			return nil, tcpskeleton.ErrConnClosing
		} else {
			fullBuf.Write(data[:readLengh])
			index := bytes.Index(fullBuf.Bytes(), endTag)
			if index > -1 {
				command := fullBuf.Next(index)
				fullBuf.Next(2)
				//fmt.Println(string(command))

				commandList := strings.Split(string(command), " ")
				if len(commandList) > 1 {
					return NewTelnetPacket(commandList[0], []byte(commandList[1])), nil
				} else {
					if commandList[0] == "quit" {
						return NewTelnetPacket("quit", command), nil
					} else {
						return NewTelnetPacket("unknow", command), nil
					}
				}
			}
		}
	}
}


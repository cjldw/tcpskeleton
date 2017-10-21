package telnet

var (
	endTag = []byte("\r\n") // Telnet command's end tag
)

// Packet
type TelnetPacket struct {
	pType string
	pData []byte
}

func (p *TelnetPacket) Serialize() []byte {
	buf := p.pData
	buf = append(buf, endTag...)
	return buf
}

func (p *TelnetPacket) GetType() string {
	return p.pType
}

func (p *TelnetPacket) GetData() []byte {
	return p.pData
}

func NewTelnetPacket(pType string, pData []byte) *TelnetPacket {
	return &TelnetPacket{
		pType: pType,
		pData: pData,
	}
}


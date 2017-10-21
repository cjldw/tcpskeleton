package tcpskeleton

// Packet
type Packet interface {
	Serialize() []byte
}


package pfcpie

type PacketRate struct {
	PacketRatedata []byte `json:"packetRatedata,omitempty"`
}

func (p *PacketRate) MarshalBinary() (data []byte, err error) {
	return p.PacketRatedata, nil
}

func (p *PacketRate) UnmarshalBinary(data []byte) error {
	p.PacketRatedata = make([]byte, len(data))
	copy(p.PacketRatedata, data)
	return nil
}

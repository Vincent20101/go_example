package pfcpie

import (
	"fmt"
)

type PFCPSRRspFlags struct {
	Drobu bool `json:"drobu,omitempty"`
}

func (p *PFCPSRRspFlags) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(p.Drobu)
	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (p *PFCPSRRspFlags) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE PFCPSRRspFlags Inadequate TLV length: %d", length)
	}
	p.Drobu = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

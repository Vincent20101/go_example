package pfcpie

import (
	"fmt"
)

type MeasurementInformation struct {
	Radi bool `json:"radi,omitempty"`
	Inam bool `json:"inam,omitempty"`
	Mbqe bool `json:"mbqe,omitempty"`
}

func (m *MeasurementInformation) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(m.Radi)<<2 |
		btou(m.Inam)<<1 |
		btou(m.Mbqe)
	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (m *MeasurementInformation) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE MeasurementInformation Inadequate TLV length: %d", length)
	}
	m.Radi = utob(uint8(data[idx]) & BitMask3)
	m.Inam = utob(uint8(data[idx]) & BitMask2)
	m.Mbqe = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

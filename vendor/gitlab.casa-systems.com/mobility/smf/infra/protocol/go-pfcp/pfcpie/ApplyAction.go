package pfcpie

import (
	"fmt"
)

type ApplyAction struct {
	DFRT bool `json:"dFRT,omitempty"`
	Dupl bool `json:"dupl,omitempty"`
	Nocp bool `json:"nocp,omitempty"`
	Buff bool `json:"buff,omitempty"`
	Forw bool `json:"forw,omitempty"`
	Drop bool `json:"drop,omitempty"`
	EDRT bool `json:"eDRT,omitempty"`
}

func (a *ApplyAction) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(a.DFRT)<<7 |
		btou(a.Dupl)<<4 |
		btou(a.Nocp)<<3 |
		btou(a.Buff)<<2 |
		btou(a.Forw)<<1 |
		btou(a.Drop)
	data = append([]byte(""), byte(tmpUint8))

	// Octet 6
	tmpUint8 = btou(a.EDRT)
	data = append(data, byte(tmpUint8))

	return data, nil
}

func (a *ApplyAction) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE ApplyAction Inadequate TLV length: %d", length)
	}
	a.DFRT = utob(uint8(data[idx]) & BitMask8)
	a.Dupl = utob(uint8(data[idx]) & BitMask5)
	a.Nocp = utob(uint8(data[idx]) & BitMask4)
	a.Buff = utob(uint8(data[idx]) & BitMask3)
	a.Forw = utob(uint8(data[idx]) & BitMask2)
	a.Drop = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet 6
	if length < idx+1 {
		return nil
	}
	a.EDRT = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

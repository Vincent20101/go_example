package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type Precedence struct {
	PrecedenceValue uint32 `json:"precedenceValue,omitempty" cmp:"ignore"`
}

func (p *Precedence) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, p.PrecedenceValue)

	return data, nil
}

func (p *Precedence) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE Precedence Inadequate TLV length: %d", length)
	}
	p.PrecedenceValue = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

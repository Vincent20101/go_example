package pfcpie

import (
	"encoding/binary"
	fmt "fmt"
)

type Multiplier struct {
	ValueDigits uint64 `json:"valueDigits,omitempty"`
	Exponent    uint32 `json:"exponent,omitempty"`
}

func (m *Multiplier) MarshalBinary() ([]byte, error) {
	var data []byte
	var buffer1 [8]byte
	binary.BigEndian.PutUint64(buffer1[:], m.ValueDigits)
	data = append(data, buffer1[:]...)
	var buffer2 [4]byte
	binary.BigEndian.PutUint32(buffer2[:], m.Exponent)
	data = append(data, buffer2[:]...)
	return data, nil
}

func (m *Multiplier) UnmarshalBinary(data []byte) error {
	if len(data) < 8 {
		return fmt.Errorf("ie Multiplier.ValueDigits Inadequate TLV length: %d", len(data))
	}
	m.ValueDigits = binary.BigEndian.Uint64(data)
	data = data[8:]
	if len(data) < 4 {
		return fmt.Errorf("ie Multiplier.ValueDigits Exponent TLV length: %d", len(data))
	}
	m.Exponent = binary.BigEndian.Uint32(data)
	data = data[4:]
	return nil
}

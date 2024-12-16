package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type NumberOfReport struct {
	NumberOfReport uint16 `json:"numberOfReport,omitempty"`
}

func (t *NumberOfReport) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 6
	data = make([]byte, 2)
	binary.BigEndian.PutUint16(data[idx:], t.NumberOfReport)

	return data, nil
}

func (t *NumberOfReport) UnmarshalBinary(data []byte) error {

	length := uint16(len(data))
	var idx uint16 = 0
	// Octet 5 to 6
	if length < idx+2 {
		return fmt.Errorf("IE NumberOfReport Inadequate TLV length: %d", length)
	}
	t.NumberOfReport = binary.BigEndian.Uint16(data[idx:])
	idx = idx + 2

	return nil
}

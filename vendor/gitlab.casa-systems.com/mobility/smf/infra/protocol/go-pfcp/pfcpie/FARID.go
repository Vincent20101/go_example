package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type FARID struct {
	FarIdValue uint32 `json:"farIdValue,omitempty" cmp:"ignore"`
}

func (f *FARID) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data[idx:], f.FarIdValue)

	return data, nil
}

func (f *FARID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE FARID Inadequate TLV length: %d", length)
	}
	f.FarIdValue = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

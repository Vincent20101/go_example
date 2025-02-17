package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type URRID struct {
	UrrIdValue uint32 `json:"urrIdValue,omitempty" cmp:"ignore"`
}

func (u *URRID) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data[idx:], u.UrrIdValue)

	return data, nil
}

func (u *URRID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie URRID Inadequate TLV length: %d", length)
	}
	u.UrrIdValue = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

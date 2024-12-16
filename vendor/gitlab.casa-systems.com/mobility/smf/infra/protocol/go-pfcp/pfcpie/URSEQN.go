package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type URSEQN struct {
	UrseqnValue uint32 `json:"urseqnValue,omitempty"`
}

func (u *URSEQN) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data[idx:], u.UrseqnValue)

	return data, nil
}

func (u *URSEQN) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie URSEQN Inadequate TLV length: %d", length)
	}
	u.UrseqnValue = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

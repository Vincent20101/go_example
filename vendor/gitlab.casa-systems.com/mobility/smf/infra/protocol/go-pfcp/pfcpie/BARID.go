package pfcpie

import (
	"fmt"
)

type BARID struct {
	BarIdValue uint8 `json:"barIdValue,omitempty" cmp:"ignore"`
}

func (b *BARID) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(b.BarIdValue))

	return data, nil
}

func (b *BARID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE BARID Inadequate TLV length: %d", length)
	}
	b.BarIdValue = uint8(data[idx])
	idx = idx + 1

	return nil
}

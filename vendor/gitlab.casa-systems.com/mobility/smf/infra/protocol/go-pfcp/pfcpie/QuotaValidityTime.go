package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type QuotaValidityTime struct {
	QuotaValidityTime uint32 `json:"quotaValidityTime,omitempty"`
}

func (t *QuotaValidityTime) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, t.QuotaValidityTime)

	return data, nil
}

func (t *QuotaValidityTime) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie QuotaValidityTime Inadequate TLV length: %d", length)
	}
	t.QuotaValidityTime = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

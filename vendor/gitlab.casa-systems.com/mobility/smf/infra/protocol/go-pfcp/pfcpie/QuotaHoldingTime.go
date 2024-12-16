package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type QuotaHoldingTime struct {
	QuotaHoldingTimeValue uint32 `json:"quotaHoldingTimeValue,omitempty"`
}

func (q *QuotaHoldingTime) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, q.QuotaHoldingTimeValue)

	return data, nil
}

func (q *QuotaHoldingTime) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE QuotaHoldingTime Inadequate TLV length: %d", length)
	}
	q.QuotaHoldingTimeValue = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

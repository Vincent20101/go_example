package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type QERCorrelationID struct {
	QerCorrelationIdValue uint32 `json:"qerCorrelationIdValue,omitempty"`
}

func (q *QERCorrelationID) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, q.QerCorrelationIdValue)

	return data, nil
}

func (q *QERCorrelationID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE QERCorrelationID Inadequate TLV length: %d", length)
	}
	q.QerCorrelationIdValue = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

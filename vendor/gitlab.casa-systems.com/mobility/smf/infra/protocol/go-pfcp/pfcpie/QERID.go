package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type QERID struct {
	QerIdValue uint32 `json:"qerIdValue,omitempty" cmp:"ignore"`
}

func (q *QERID) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data[idx:], q.QerIdValue)

	return data, nil
}

func (q *QERID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE QERID Inadequate TLV length: %d", length)
	}
	q.QerIdValue = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

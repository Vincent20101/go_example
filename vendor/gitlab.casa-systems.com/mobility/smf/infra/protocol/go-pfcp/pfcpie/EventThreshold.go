package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type EventThreshold struct {
	EventThreshold uint32 `json:"eventThreshold,omitempty"`
}

func (e *EventThreshold) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, e.EventThreshold)

	return data, nil
}

func (e *EventThreshold) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE EventThreshold Inadequate TLV length: %d", length)
	}
	e.EventThreshold = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

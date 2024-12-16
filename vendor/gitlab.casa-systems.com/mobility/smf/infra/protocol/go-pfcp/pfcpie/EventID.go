package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type EventID struct {
	EventId uint32 `json:"eventId,omitempty"`
}

func (e *EventID) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, e.EventId)

	return data, nil
}

func (e *EventID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE EventID Inadequate TLV length: %d", length)
	}
	e.EventId = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

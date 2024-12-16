package pfcpie

import (
	"encoding/binary"
	"fmt"
)

const (
	DefaultEthernetInactivityTimer = 30
)

type EthernetInactivityTimer struct {
	EthernetInactivityTimer uint32 `json:"ethernetInactivityTimer,omitempty"`
}

func (e *EthernetInactivityTimer) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, e.EthernetInactivityTimer)

	return data, nil
}

func (e *EthernetInactivityTimer) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE EthernetInactivityTimer Inadequate TLV length: %d", length)
	}
	e.EthernetInactivityTimer = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

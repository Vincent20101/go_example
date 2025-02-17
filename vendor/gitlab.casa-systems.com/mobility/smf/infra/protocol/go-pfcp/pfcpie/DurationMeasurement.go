package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type DurationMeasurement struct {
	DurationValue uint32 `json:"durationValue,omitempty"`
}

func (d *DurationMeasurement) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, d.DurationValue)

	return data, nil
}

func (d *DurationMeasurement) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE DurationMeasurement Inadequate TLV length: %d", length)
	}
	d.DurationValue = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

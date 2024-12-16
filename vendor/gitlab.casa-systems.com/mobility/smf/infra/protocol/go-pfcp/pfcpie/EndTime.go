package pfcpie

import (
	"encoding/binary"
	"fmt"
	"time"
)

type EndTime struct {
	EndTime time.Time `json:"endTime,omitempty"`
}

func (e *EndTime) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	ts := TimeToTimestampValue(e.EndTime)
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, ts)

	return data, nil
}

func (e *EndTime) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE EndTime Inadequate TLV length: %d", length)
	}
	ts := binary.BigEndian.Uint32(data[idx:])
	e.EndTime = TimestampValueToTime(ts)
	idx = idx + 4

	return nil
}

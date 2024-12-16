package pfcpie

import (
	"encoding/binary"
	"fmt"
	"time"
)

type TimeOfFirstPacket struct {
	TimeOfFirstPacket time.Time `json:"timeOfFirstPacket,omitempty"`
}

func (t *TimeOfFirstPacket) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	ts := TimeToTimestampValue(t.TimeOfFirstPacket)
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, ts)

	return data, nil
}

func (t *TimeOfFirstPacket) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie TimeOfFirstPacket Inadequate TLV length: %d", length)
	}
	ts := binary.BigEndian.Uint32(data[idx:])
	t.TimeOfFirstPacket = TimestampValueToTime(ts)
	idx = idx + 4

	return nil
}

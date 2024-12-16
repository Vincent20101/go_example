package pfcpie

import (
	"encoding/binary"
	"fmt"
	"time"
)

type TimeOfLastPacket struct {
	TimeOfLastPacket time.Time `json:"timeOfLastPacket,omitempty"`
}

func (t *TimeOfLastPacket) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	ts := TimeToTimestampValue(t.TimeOfLastPacket)
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, ts)

	return data, nil
}

func (t *TimeOfLastPacket) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie TimeOfLastPacket Inadequate TLV length: %d", length)
	}
	ts := binary.BigEndian.Uint32(data[idx:])
	t.TimeOfLastPacket = TimestampValueToTime(ts)
	idx = idx + 4

	return nil
}

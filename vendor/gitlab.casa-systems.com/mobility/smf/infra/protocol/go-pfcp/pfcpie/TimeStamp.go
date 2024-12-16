package pfcpie

import (
	"encoding/binary"
	"fmt"
	"time"
)

type TimeStamp struct {
	TimeStamp time.Time `json:"timeStamp,omitempty"`
}

func (t *TimeStamp) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	ts := TimeToTimestampValue(t.TimeStamp)
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, ts)

	return data, nil
}

func (t *TimeStamp) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie TimeStamp Inadequate TLV length: %d", length)
	}
	ts := binary.BigEndian.Uint32(data[idx:])
	t.TimeStamp = TimestampValueToTime(ts)
	idx = idx + 4

	return nil
}

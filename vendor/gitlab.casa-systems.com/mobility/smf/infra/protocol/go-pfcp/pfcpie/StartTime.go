package pfcpie

import (
	"encoding/binary"
	"fmt"
	"time"
)

type StartTime struct {
	StartTime time.Time `json:"startTime,omitempty"`
}

func (s *StartTime) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	ts := TimeToTimestampValue(s.StartTime)
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, ts)

	return data, nil
}

func (s *StartTime) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie StartTime Inadequate TLV length: %d", length)
	}
	ts := binary.BigEndian.Uint32(data[idx:])
	s.StartTime = TimestampValueToTime(ts)
	idx = idx + 4

	return nil
}

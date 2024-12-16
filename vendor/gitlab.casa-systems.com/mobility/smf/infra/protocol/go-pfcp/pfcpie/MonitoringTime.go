package pfcpie

import (
	"encoding/binary"
	"fmt"
	"time"
)

type MonitoringTime struct {
	MonitoringTime time.Time `json:"monitoringTime,omitempty"`
}

func (m *MonitoringTime) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	ts := TimeToTimestampValue(m.MonitoringTime)
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, ts)

	return data, nil
}

func (m *MonitoringTime) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE MonitoringTime Inadequate TLV length: %d", length)
	}
	ts := binary.BigEndian.Uint32(data[idx:])
	m.MonitoringTime = TimestampValueToTime(ts)
	idx = idx + 4

	return nil
}

package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// 8.2.149	Time Offset Measurement
// The Time Offset Measurement IE type shall be encoded as shown in Figure 8.2.149-1.
//                                         Bits
//  Octets        8       7       6       5       4       3       2       1
//            +----------------------------------------------------------------+
//  1 to 2    |                    Type = 209(decimal)                         |
//            |----------------------------------------------------------------|
//  3 to 4    |                      Length = n                                |
//            |----------------------------------------------------------------|
//  5 to 12   |                   Time Offset Measurement                      |
//            |----------------------------------------------------------------|
// 13 to (n+4)|   These octet(s) is/are present only if explicitly specified   |
//            +----------------------------------------------------------------+
//
// The Time Offset Measurement field shall be encoded as a signed64 binary integer value.
// It shall contain the Time Offset Measurement in nanoseconds. It shall contain the time offset
// between the 5GS clock and the clock of the TSN time domain.

type TimeOffsetMeasurement struct {
	TimeOffsetMeasurement uint64 `json:"timeOffsetMeasurement,omitempty"`
}

func (s *TimeOffsetMeasurement) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 12
	data = make([]byte, 8)
	binary.BigEndian.PutUint64(data, s.TimeOffsetMeasurement)

	return data, nil
}

func (s *TimeOffsetMeasurement) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 12
	if length < idx+8 {
		return fmt.Errorf("ie TimeOffsetMeasurement Inadequate TLV length: %d", length)
	}
	s.TimeOffsetMeasurement = binary.BigEndian.Uint64(data[idx:])

	return nil
}

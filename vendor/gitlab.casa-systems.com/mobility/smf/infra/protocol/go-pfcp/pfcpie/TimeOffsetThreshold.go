package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// 8.2.147	Time Offset Threshold
// The Time Offset Threshold IE type shall be encoded as shown in Figure 8.2.147-1.
//                                         Bits
//  Octets        8       7       6       5       4       3       2       1
//            +----------------------------------------------------------------+
//  1 to 2    |                    Type = 207(decimal)                         |
//            |----------------------------------------------------------------|
//  3 to 4    |                      Length = n                                |
//            |----------------------------------------------------------------|
//  5 to 12   |                    Time Offset Threshold                       |
//            |----------------------------------------------------------------|
// 13 to (n+4)|   These octet(s) is/are present only if explicitly specified   |
//            +----------------------------------------------------------------+

// The Time Offset Threshold field shall be encoded as a signed64 binary integer value.
// It shall contain the Time Offset Threshold in nanoseconds.

type TimeOffsetThreshold struct {
	TimeOffsetThreshold uint64 `json:"timeOffsetThreshold,omitempty"`
}

func (s *TimeOffsetThreshold) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 12
	data = make([]byte, 8)
	binary.BigEndian.PutUint64(data, s.TimeOffsetThreshold)

	return data, nil
}

func (s *TimeOffsetThreshold) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 12
	if length < idx+8 {
		return fmt.Errorf("ie TimeOffsetThreshold Inadequate TLV length: %d", length)
	}
	s.TimeOffsetThreshold = binary.BigEndian.Uint64(data[idx:])

	return nil
}

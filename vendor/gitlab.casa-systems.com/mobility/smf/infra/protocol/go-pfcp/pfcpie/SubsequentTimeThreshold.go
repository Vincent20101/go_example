package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// SubsequentTimeThreshold IE. Refer to [TS29244 8.2.17  Subsequent Time Threshold].
//
// The Subsequent Time Threshold IE contains the subsequent traffic duration
// threshold to be monitored by the UP function after the Monitoring Time. It
// shall be encoded as shown in Figure 8.2.17-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                    Type = 35(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	5 to 8   |                  Subsequent Time Threshold                     |
//	         |----------------------------------------------------------------|
//
// 9 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The Subsequent Time Threshold field shall be encoded as an Unsigned32 binary
// integer value. It shall contain the duration in seconds.
type SubsequentTimeThreshold struct {
	SubsequentTimeThreshold uint32 `json:"subsequentTimeThreshold,omitempty"`
}

func (s *SubsequentTimeThreshold) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, s.SubsequentTimeThreshold)

	return data, nil
}

func (s *SubsequentTimeThreshold) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie SubsequentTimeThreshold Inadequate TLV length: %d", length)
	}
	s.SubsequentTimeThreshold = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

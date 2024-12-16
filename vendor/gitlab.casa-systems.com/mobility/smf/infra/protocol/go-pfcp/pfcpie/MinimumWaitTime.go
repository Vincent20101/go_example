package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// 8.2.170	Minimum Wait Time
// The Minimum Wait Time IE contains the minimum waiting time between two consecutive reports for event triggered QoS monitoring reporting. It shall be encoded as shown in Figure 8.2.170-1.
//
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                    Type = 246(decimal)                          |
//           |----------------------------------------------------------------|
//  3 to 4   |                      Length = n                                |
//           |----------------------------------------------------------------|
//  5 to 8   |                     Minimum Wait Time                          |
//           |----------------------------------------------------------------|
// 9 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//           +----------------------------------------------------------------+
//
// The Minimum Wait Time field shall be encoded as an Unsigned32 binary integer value. It shall contain the duration in seconds.
//

type MinimumWaitTime struct {
	Time uint32 `json:"time,omitempty"`
}

func (m *MinimumWaitTime) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, m.Time)

	return data, nil
}

func (m *MinimumWaitTime) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE MinimumWaitTime Inadequate TLV length: %d", length)
	}
	m.Time = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

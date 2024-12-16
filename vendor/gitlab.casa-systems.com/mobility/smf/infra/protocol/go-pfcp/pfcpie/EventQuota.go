package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// EventQuota IE. Refer to [TS29244 8.2.112 Event Quota].
//
// The Event Quota IE type shall be encoded as shown in Figure 8.2.112-1. It
// contains the event quota to be monitored by the UP function.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 148(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	5 to 8   |                     Event Quota                                |
//	         |----------------------------------------------------------------|
//
// 8 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The Event Quota field shall be encoded as an Unsigned32 binary integer
// value.
type EventQuota struct {
	EventQuota uint32 `json:"eventQuota,omitempty"`
}

func (e *EventQuota) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, e.EventQuota)

	return data, nil
}

func (e *EventQuota) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE EventQuota Inadequate TLV length: %d", length)
	}
	e.EventQuota = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// UserPlaneInactivityTimer IE.
// Refer to [TS29244 8.2.83 User Plane Inactivity Timer]
//
// The User Plane Inactivity Timer IE contains the inactivity time period, in
// seconds, to be monitored by the UP function.  It shall be encoded as shown
// in Figure 8.2.83-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 117(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	5 to 8   |                User Plane Inactivity Timer                     |
//	         |----------------------------------------------------------------|
//	9 to n+4 |   These octets(s) is/are present only if explicity specified   |
//	         +----------------------------------------------------------------+
//
// The User Plane Inactivity Timer field shall be encoded as an Unsigned32
// binary integer value. The timer value "0" shall be interpreted as an
// indication that user plane inactivity detection and reporting is stopped.
type UserPlaneInactivityTimer struct {
	UserPlaneInactivityTimerValue uint32 `json:"userPlaneInactivityTimerValue,omitempty"`
}

func (u *UserPlaneInactivityTimer) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, u.UserPlaneInactivityTimerValue)
	return data, nil
}

func (u *UserPlaneInactivityTimer) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie UserPlaneInactivityTimer Inadequate TLV length: %d", length)
	}
	sli := make([]byte, 4)
	copy(sli, data)
	u.UserPlaneInactivityTimerValue = binary.BigEndian.Uint32(sli)
	idx = idx + 4

	return nil
}

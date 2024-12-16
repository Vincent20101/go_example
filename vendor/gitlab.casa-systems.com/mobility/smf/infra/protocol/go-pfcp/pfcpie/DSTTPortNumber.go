package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// 8.2.141	DS-TT Port Number
// The DS-TT Port Number IE shall be encoded as shown in Figure 8.2.141-1.
// It shall contain one Port Number value.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 196 (decimal)                         |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = 4                                |
//	         |----------------------------------------------------------------|
//	5 to 8   |                    Port Number value                           |
//	         +----------------------------------------------------------------+
type DSTTPortNumber struct {
	DSPortNumbervalue uint32 `json:"dSPortNumbervalue,omitempty"`
}

func (f *DSTTPortNumber) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, f.DSPortNumbervalue)
	return data, nil
}

func (f *DSTTPortNumber) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie FramedRouting Inadequate TLV length: %d", length)
	}
	f.DSPortNumbervalue = binary.BigEndian.Uint32(data[idx:])

	return nil
}

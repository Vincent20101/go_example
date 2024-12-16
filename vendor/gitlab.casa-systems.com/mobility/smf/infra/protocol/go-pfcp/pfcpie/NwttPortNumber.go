package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// 8.2.142	NW-TT Port Number
// The NW-TT Port Number IE shall be encoded as shown in Figure 8.2.142-1.
// It shall contain one Port Number value.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 197 (decimal)                         |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = 4                                |
//	         |----------------------------------------------------------------|
//	5 to 8   |                    Port Number value                           |
//	         +----------------------------------------------------------------+
type NWTTPortNumber struct {
	NWPortNumbervalue uint32 `json:"nWPortNumbervalue,omitempty"`
}

func (f *NWTTPortNumber) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, f.NWPortNumbervalue)
	return data, nil
}

func (f *NWTTPortNumber) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie FramedRouting Inadequate TLV length: %d", length)
	}
	f.NWPortNumbervalue = binary.BigEndian.Uint32(data[idx:])

	return nil
}

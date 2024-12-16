package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// TS 2924 8.2.98 Ethernet Filter ID
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 138(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	5 to 8   |                 	 Ethernet Filter ID Value                     |
//	         |----------------------------------------------------------------|
//
// 10 to n+4 |   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The Ethernet Filter ID value shall be encoded as an Unsigned32 binary
// integer value.
type EthernetFilterID struct {
	EthernetFilterIDValue uint32 `json:"ethernetFilterIDValue,omitempty"`
}

func (e *EthernetFilterID) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, e.EthernetFilterIDValue)
	return data, nil
}

func (e *EthernetFilterID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE EthernetFilterID Inadequate TLV length: %d", length)
	}
	sli := make([]byte, 4)
	copy(sli, data)
	e.EthernetFilterIDValue = binary.BigEndian.Uint32(sli)
	idx = idx + 4

	return nil
}

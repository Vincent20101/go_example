package pfcpie

import "fmt"

// TS29244 8.2.99 Ethernet Filter Properties
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 139(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      |             Spare                                     | BIDE   |
//	         |----------------------------------------------------------------|
//
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//
//   - Bit 1 – BIDE (Bidirectional Ethernet Filter): If this bit is set to
//     "1", then the Ethernet Filter identified by the  Ethernet Filter ID is
//     bidirectional.
//
//   - Bit 2 to 8 – spare and reserved for future use.
type EthernetFilterProperties struct {
	Bide bool `json:"bide,omitempty"`
}

func (e *EthernetFilterProperties) MarshalBinary() (data []byte, err error) {
	data = append([]byte(""), byte(btou(e.Bide)))
	return data, nil
}

func (e *EthernetFilterProperties) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))
	if length < 1 {
		return fmt.Errorf("IE EthernetFilterProperties Inadequate TLV length: %d", length)
	}
	e.Bide = utob(uint8(data[0]) & BitMask1)
	return nil
}

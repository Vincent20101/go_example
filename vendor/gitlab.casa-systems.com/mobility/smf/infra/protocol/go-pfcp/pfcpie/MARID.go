package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// TS29.244 8.2.123	MAR ID
//
// The MAR ID IE type shall be encoded as shown in Figure 8.2.123-1. It shall contain a Multi-Access Rule ID.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 170(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	5 to 6   |                     MAR ID value                               |
//	         |----------------------------------------------------------------|
//
// 7 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The MAR ID value field shall be encoded as an Unsigned16 binary integer value.
type MARID struct {
	MarIdValue uint16 `json:"marIdValue,omitempty" cmp:"ignore"`
}

func (m *MARID) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 6
	data = make([]byte, 2)
	binary.BigEndian.PutUint16(data[idx:], m.MarIdValue)

	return data, nil
}

func (m *MARID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 6
	if length < idx+2 {
		return fmt.Errorf("IE MARID Inadequate TLV length: %d", length)
	}
	m.MarIdValue = binary.BigEndian.Uint16(data[idx:])
	idx = idx + 2

	return nil
}

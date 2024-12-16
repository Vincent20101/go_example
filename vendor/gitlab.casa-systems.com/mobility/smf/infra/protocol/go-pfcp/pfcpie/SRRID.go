package pfcpie

import (
	"fmt"
)

//8.2.151	SRR ID
//  The SRR ID IE type shall be encoded as shown in Figure 8.2.151-1. It contains a Session Reporting Rule ID.
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                   Type = 215(decimal)                          |
//           |----------------------------------------------------------------|
//  3 to 4   |                      Length = n                                |
//           |----------------------------------------------------------------|
//  5        |                     SRR ID value                               |
//           |----------------------------------------------------------------|
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//           +----------------------------------------------------------------+
//
// The SRR ID value shall be encoded as a binary integer value.

type SRRID struct {
	SrrIdValue uint8 `json:"srrIdValue,omitempty" cmp:"ignore"`
}

func (s *SRRID) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(s.SrrIdValue))

	return data, nil
}

func (s *SRRID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE SRRID Inadequate TLV length: %d", length)
	}
	s.SrrIdValue = uint8(data[idx])
	idx = idx + 1

	return nil
}

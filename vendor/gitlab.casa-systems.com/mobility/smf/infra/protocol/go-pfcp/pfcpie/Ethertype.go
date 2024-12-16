package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// [TS 29244 8.2.96 EtherType]
// The Ethertype IE type shall be encoded as shown in Figure 8.2.96-1. It
// contains an Ethertype as defined in  IEEE 802.3 [31].
//
//	Octets
//	          +-------------------------------------------------------------+
//	1 to 2    |                   Type = 136(decimal)                       |
//	          |-------------------------------------------------------------|
//	3 to 4    |                      Length = n                             |
//	          |-------------------------------------------------------------|
//	5 to 6    |                      EtherType                              |
//	          |-------------------------------------------------------------|
//	7 to (n+4)| These octets(s) is/are present only if explicity specified  |
//	          +-------------------------------------------------------------+
type Ethertype struct {
	Ethertypedata uint16 `json:"ethertypedata,omitempty"`
}

func (e *Ethertype) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 2)
	binary.BigEndian.PutUint16(data, e.Ethertypedata)

	return data, nil
}

func (e *Ethertype) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 6
	if length < idx+2 {
		return fmt.Errorf("IE Ethertype Inadequate TLV length: %d", length)
	}
	e.Ethertypedata = binary.BigEndian.Uint16(data[idx:])
	idx = idx + 2

	return nil
}

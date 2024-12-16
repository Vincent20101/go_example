package pfcpie

import (
	"fmt"
)

// TS29244 8.2.140 Create Bridge Info for TSC IE
//
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                   Type = 194 (decimal)                         |
//           |----------------------------------------------------------------|
//  3 to 4   |                      Length = n                                |
//           |----------------------------------------------------------------|
//    5      |             Spare                                     |  BII   |
//           |----------------------------------------------------------------|
// m to (n+4)|   These octet(s) is/are present only if explicitly specified   |
//           +----------------------------------------------------------------+
// The Create Bridge Info for TSC IE shall be encoded as shown in Figure 8.2.140-1.
//
// The following flags are coded within Octet 5:
// -	Bit 1 – BII (Bridge Information Indication): If this bit is set to "1",
// then the Bridge Information comprising a DS-TT port number and
// the related TSN Bridge ID is requested to be provided.

// -	Bit 2 to 8 Spare, for future use and set to "0".
//

type CreateBridgeInfoForTSC struct {
	BII bool `json:"bII,omitempty"`
}

func (p *CreateBridgeInfoForTSC) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(p.BII)
	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (p *CreateBridgeInfoForTSC) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE CreateBridgeInfoForTSC Inadequate TLV length: %d", length)
	}
	p.BII = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

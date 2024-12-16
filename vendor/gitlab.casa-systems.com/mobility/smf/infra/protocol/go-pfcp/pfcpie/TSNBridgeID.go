package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// 8.2.143	TSN Bridge ID
// The TSN Bridge ID IE shall be encoded as shown in Figure 8.2.143-1.
//
//                                         Bits
//  Octets        8       7       6       5       4       3       2       1
//            +----------------------------------------------------------------+
//  1 to 2    |                   Type = 198 (decimal)                         |
//            |----------------------------------------------------------------|
//  3 to 4    |                      Length = n                                |
//            |----------------------------------------------------------------|
//  5         |                        Spare                          |  BID   |
//            |----------------------------------------------------------------|
// m to (m+7) |                     Bridge ID value                            |
//            |----------------------------------------------------------------|
// s to (n+4) |   These octet(s) is/are present only if explicitly specified   |
//            +----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
// -	Bit 1 – BID: If this bit is set to "1", then the Bridge ID value field shall be present.
//
// -	Bit 2 to 8: Spare, for future use and set to "0".
// The Bridge ID is defined in IEEE 802.1Q [30] clause 14.2.5 and value shall be encoded as
// an Unsigned64 binary integer.
//

type TSNBridgeID struct {
	BID           bool   `json:"bID,omitempty"`
	BridgeIdValue uint64 `json:"bridgeIdValue,omitempty"`
}

func (t *TSNBridgeID) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 13
	data = append([]byte(""), byte(btou(t.BID)))
	temp := make([]byte, 8)
	binary.BigEndian.PutUint64(temp, t.BridgeIdValue)
	data = append(data, temp...)

	return data, nil
}

func (t *TSNBridgeID) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))

	var idx uint8 = 0
	// Octet 5 to 13
	if length < idx+9 {
		return fmt.Errorf("ie TransportLevelMarking Inadequate TLV length: %d", length)
	}
	t.BID = utob(uint8(data[0]) & BitMask1)
	idx = idx + 1
	t.BridgeIdValue = binary.BigEndian.Uint64(data[idx:])
	return nil
}

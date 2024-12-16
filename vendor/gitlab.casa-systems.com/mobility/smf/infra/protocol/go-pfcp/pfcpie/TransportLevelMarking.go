package pfcpie

import (
	"fmt"
)

// TransportLevelMarking IE, Refer to [TS29244 8.2.12 Transport Level Marking]
//
// The Transport Level Marking IE type shall be encoded as shown in Figure
// 8.2.12-1. It indicates the DSCP to be used for UL/DL transport level
// marking.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                       Type = 30(decimal)                      |
//	         |---------------------------------------------------------------|
//	3 to 4   |                           Length = n                          |
//	         |-------------------------------+-------------------------------|
//	5 to 6   |                        ToS/Traffic Class                      |
//	         |-------------------------------+-------------------------------|
//	7 to n+4 |                         Node ID Value                         |
//	         |---------------------------------------------------------------|
//	m to n+4 |   These octets(s) is/are present only if explicity specified  |
//	         +---------------------------------------------------------------+
//
// The ToS/Traffic Class shall be encoded on two octets as an OctetString. The
// first octet shall contain the DSCP value in the IPv4 Type-of-Service or the
// IPv6 Traffic-Class field and the second octet shall contain the ToS/Traffic
// Class mask field, which shall be set to "0xFC". See clause 5.3.15 of
// 3GPP TS 29.212 [8].
//
// NOTE:  The original ECN bits in the IP header of user plane packets are not
// changed after applying transport level marking.
//
// ==> [Wang Chenglong]: We separate the field `Tos/Traffic Class` into two
// uint8 rather than a whole byte slice, it makes more convenient for us to
// code and config.
type TransportLevelMarking struct {
	Tos  uint8 `json:"tos,omitempty"`
	Mask uint8 `json:"mask,omitempty"`
}

func (t *TransportLevelMarking) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 6
	data = []byte{t.Tos, t.Mask}
	return data, nil
}

func (t *TransportLevelMarking) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))

	// Octet 5 to 6
	if length < 2 {
		return fmt.Errorf("ie TransportLevelMarking Inadequate TLV length: %d", length)
	}
	t.Tos = uint8(data[0])
	t.Mask = uint8(data[1])
	return nil
}

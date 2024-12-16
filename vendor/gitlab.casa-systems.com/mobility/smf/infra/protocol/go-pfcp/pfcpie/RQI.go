package pfcpie

import "fmt"

// RQI IE: Refer to
//
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                       Type = 124(decimal)                      |
//           |----------------------------------------------------------------|
//  3 to 4   |                          Length = n                            |
//           |----------------------------------------------------------------|
//    5      |     Spare     |                  QFI Value                     |
//           |----------------------------------------------------------------|
// p to (n+4)|   These octets(s) is/are present only if explicity specified   |
//           +----------------------------------------------------------------+
//
// The QFI value shall be encoded as binary integer value, as specified in
// clause 5.5.3.3 of 3GPP TS 38.415 [34].
//

type RQI struct {
	Rqi bool `json:"rqi,omitempty"`
}

func (r *RQI) MarshalBinary() (data []byte, err error) {
	return []byte{btou(r.Rqi)}, nil
}

func (r *RQI) UnmarshalBinary(data []byte) error {
	length := len(data)
	if length < 1 {
		return fmt.Errorf("ie RQI empty body")
	}
	r.Rqi = utob(uint8(data[0]) & 0x01)
	return nil
}

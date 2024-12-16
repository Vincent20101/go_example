package pfcpie

import "fmt"

// QFI IE: Refer to [TS29244 8.2.89 QFI]
//
//	The QFI IE type shall be encoded as shown in Figure 8.2.89-1. It contains
//	an QoS flow identifier identifying a QoS  flow in a 5G system filter.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                       Type = 124(decimal)                      |
//	         |----------------------------------------------------------------|
//	3 to 4   |                          Length = n                            |
//	         |----------------------------------------------------------------|
//	  5      |     Spare     |                  QFI Value                     |
//	         |----------------------------------------------------------------|
//
// p to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The QFI value shall be encoded as binary integer value, as specified in
// clause 5.5.3.3 of 3GPP TS 38.415 [34].
type QFI struct {
	Qfi uint8 `json:"qfi,omitempty"`
}

func (q *QFI) MarshalBinary() (data []byte, err error) {
	data = append(data, byte(q.Qfi&0x3F))
	return data, nil
}

func (q *QFI) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))
	if length < 1 {
		return fmt.Errorf("IE QFI empty body")
	}
	q.Qfi = uint8(data[0] & 0x3f)
	return nil
}

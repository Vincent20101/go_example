package pfcpie

import (
	"fmt"
)

// MeasurementMethod IE. Refer to [TS29244 8.2.40  Measurement Method].
//
// The Measurement Method IE shall be encoded as shown in Figure 8.2.40-1. It
// indicates the method for measuring the usage of network resources.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                    Type = 62(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      | Spare | Spare | Spare | Spare | Spare | EVENT | VOLUM | DURAT  |
//	         |----------------------------------------------------------------|
//
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// Octet 5 shall be encoded as follows:
//
//   - Bit 1 – DURAT (Duration): when set to "1", this indicates a request
//     for measuring the duration of the traffic.
//
//   - Bit 2 – VOLUM (Volume): when set to "1", this indicates a request for
//     measuring the volume of the traffic.
//
//   - Bit 3 – EVENT (Event): when set to "1", this indicates a request for
//     measuring the events.
//
//   - Bit 4 to 8: Spare, for future use and set to "0".
//
// At least one bit shall be set to "1". Several bits may be set to "1".
type MeasurementMethod struct {
	Event bool `json:"event,omitempty"`
	Volum bool `json:"volum,omitempty"`
	Durat bool `json:"durat,omitempty"`
}

func (m *MeasurementMethod) MarshalBinary() (data []byte, err error) {
	// Octet 5
	if !m.Event && !m.Volum && !m.Durat {
		return []byte(""), fmt.Errorf("At least one of EVENT, VOLUM and DURAT shall be set")
	}
	tmpUint8 := btou(m.Event)<<2 |
		btou(m.Volum)<<1 |
		btou(m.Durat)
	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (m *MeasurementMethod) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE MeasurementMethod Inadequate TLV length: %d", length)
	}
	m.Event = utob(uint8(data[idx]) & BitMask3)
	m.Volum = utob(uint8(data[idx]) & BitMask2)
	m.Durat = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

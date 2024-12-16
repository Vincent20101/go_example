package pfcpie

import "fmt"

// TS29.244 8.2.181	MPTCP Applicable Indication
//
// MPTCP Applicable Indication shall be coded as depicted in Figure 8.2.181-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 265(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |-------------------------------------------------------+--------|
//	  5      |             Spare                                     |  MAI   |
//	         |-------------------------------------------------------+--------|
//
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//   - Bit 1 – MAI (MPTCP Applicable Indication): When this bit is set to "1",
//     it indicates that the PDR is used to detect user plane traffic for which MPTCP
//     is applicable (see clause 5.20.2.4).
//   - Bit 2 to 8 are spare and reserved for future use.
type MPTCPApplicableIndication struct {
	Mai bool `json:"mai,omitempty"`
}

func (m *MPTCPApplicableIndication) MarshalBinary() (data []byte, err error) {
	data = append(data, byte(btou(m.Mai)))
	return data, nil
}

func (m *MPTCPApplicableIndication) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))
	if length < 1 {
		return fmt.Errorf("IE MPTCPApplicableIndication Inadequate TLV length: %d", length)
	}
	m.Mai = utob(uint8(data[0]))
	return nil
}

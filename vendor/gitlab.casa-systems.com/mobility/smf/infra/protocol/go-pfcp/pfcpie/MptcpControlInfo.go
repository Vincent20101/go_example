package pfcpie

import "fmt"

// TS29.244 8.2.154	MPTCP Control Information
//
// The MPTCP Control Information IE shall provide details of the required MPTCP functionality.
// It shall be encoded as shown in Figure 8.2.154-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 222(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |-------------------------------------------------------+--------|
//	  5      |             Spare                                     |  TCI   |
//	         |-------------------------------------------------------+--------|
//
// m to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//   - Bit 1 – TCI (Transport Converter Indication): If this bit is set to "1",
//     it indicates that the required MPTCP steering functionality is of type Transport
//     Converter (see IETF RFC 8803 [60]) and the UPF shall allocate relevant resource as
//     specified in clause 5.20.2.1.
//   - Bit 2 to 8 Spare, for future use and set to "0".
type MptcpControlInfo struct {
	Tci bool `json:"tci,omitempty"`
}

func (m *MptcpControlInfo) MarshalBinary() (data []byte, err error) {
	data = append(data, byte(btou(m.Tci)))
	return data, nil
}

func (m *MptcpControlInfo) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))
	if length < 1 {
		return fmt.Errorf("IE MptcpControlInfo Inadequate TLV length: %d", length)
	}
	m.Tci = utob(uint8(data[0]))
	return nil
}

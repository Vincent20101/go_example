package pfcpie

import "fmt"

// TS29.244 8.2.154	PMF Control Information
//
// The PMF Control Information IE shall provide details of the required PMF functionality.
// It shall be encoded as shown in Figure 8.2.156-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 224(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |------------------------------------------------+------+--------|
//	  5      |             Spare                              |DRTTI |  PMFI  |
//	         |------------------------------------------------+------+--------|
//
// m to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//		+----------------------------------------------------------------+
//
//	  - Bit 1 – PMFI: If this bit is set to "1", it indicates that the PMF functionality is required
//	    and the UPF shall allocate relevant resource as specified in 5.20.3.1.
//	  - Bit 2 – DRTTI (Disallow PMF RTT Indication): If this bit is set to "1", it indicates that
//	    PMF RTT measurements are not allowed. This bit shall be set to "1" if the UE does not support
//	    PMF RTT measurements (i.e. if the UE supports MPTCP with any steering mode and ATSSS-LL with only
//	    the Active-Standby steering mode, see clauses 5.32.2 and 5.32.5.1 of 3GPP TS 23.501 [28]).
//	  - Bit 3 to 8 Spare, for future use and set to "0".
type PmfControlInfo struct {
	Pmfi  bool `json:"pmfi,omitempty"`
	Drtti bool `json:"drtti,omitempty"`
}

func (p *PmfControlInfo) MarshalBinary() (data []byte, err error) {
	// Octet 5
	octet5 := btou(p.Pmfi) | btou(p.Drtti)<<1
	data = append([]byte(""), byte(octet5))

	return data, nil
}

func (a *PmfControlInfo) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))
	if length < 1 {
		return fmt.Errorf("IE PmfControlInfo Inadequate TLV length: %d", length)
	}

	a.Pmfi = utob(uint8(data[0]) & BitMask1)
	a.Drtti = utob(uint8(data[0]) & BitMask2)

	return nil
}

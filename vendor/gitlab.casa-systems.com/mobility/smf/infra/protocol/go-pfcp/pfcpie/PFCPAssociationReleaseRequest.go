package pfcpie

import "fmt"

// TS29.244-g40 8.2.77	PFCP Association Release Request
// The PFCP Association Release Request IE shall be encoded as shown in Figure 8.2.77-1.

//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                      Type = 111(decimal)                       |
//           |----------------------------------------------------------------|
//  3 to 4   |                          Length = n                            |
//           |-----------------------------------------------+-------+--------|
//    5      |             Spare                             | URSS  |  SARR  |
//           |-----------------------------------------------+-------+--------|
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//           +----------------------------------------------------------------+
// Figure 8.2.77-1: PFCP Association Release Request
// The following flags are coded within Octet 5:
// -	Bit 1 – SARR (PFCP Association Release Request): If this bit is set to "1", then the UP function requests the release of the PFCP association.
// -	Bit 2 – URSS (non-zero Usage Reports for the affected PFCP Sessions Sent): If this bit is set to "1", it indicates that the UP function has sent all the non-zero usage reports to the CP function for all PFCP essions affected by the PFCP Association Release.
// -	Bit 3 to 8: Spare, for future use and set to "0".

type PFCPAssociationReleaseRequest struct {
	URSS bool `json:"uRSS,omitempty"`
	SARR bool `json:"sARR,omitempty"`
}

func (p *PFCPAssociationReleaseRequest) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(p.URSS)<<1 |
		btou(p.SARR)
	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (p *PFCPAssociationReleaseRequest) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE PFCPAssociationReleaseRequest Inadequate TLV length: %d", length)
	}

	p.URSS = utob(uint8(data[idx]) & BitMask2)
	p.SARR = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

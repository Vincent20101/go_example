package pfcpie

import "fmt"

// Proxying IE, Refer to [TS29244 8.2.97 Proxying]
//
// The Proxying IE shall be encoded as shown in Figure 8.2.97-1. It specifies
// if responding to Address Resolution Protocol (ARP) (see IETF RFC 826 [32])
// and / or IPv6 Neighbour Solicitation (see IETF RFC 4861 [33]) as specified
// in  clause 5.6.10.2 of 3GPP TS 23.501 [28], functionality for the Ethernet
// PDUs is performed in UPF.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 137(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      |             Spare                             |  INS  |  ARP   |
//	         |----------------------------------------------------------------|
//
// 6 to (n+4)|   These octets(s) is/are present only if explicitly specified  |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//
//   - Bit 1 – ARP: If this bit is set to "1", then responding ARP is performed
//     in UPF based on local cache information.
//
//   - Bit 2 – INS: If this bit is set to "1", then responding to IPv6
//     Neighbour Solicitation is performed in UPF based on
//
// - Bit 3 to 8 – spare and reserved for future use. 8.2.98 Ethernet Filter ID
type Proxying struct {
	Arp bool `json:"arp,omitempty"`
	Ins bool `json:"ins,omitempty"`
}

// MarshalBinary marshal Proxying IE to bytes.
func (p *Proxying) MarshalBinary() (data []byte, err error) {
	tmp := btou(p.Arp) | btou(p.Ins)<<1
	data = append(data, byte(tmp))
	return data, nil
}

// MarshalBinary unmarshal Proxying IE from bytes.
func (p *Proxying) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))
	if length < 1 {
		return fmt.Errorf("IE Proxying Inadequate TLV length: %d", length)
	}
	tmp := uint8(data[0])
	p.Arp = utob(tmp & BitMask1)
	p.Ins = utob((tmp & BitMask2) >> 1)
	return nil
}

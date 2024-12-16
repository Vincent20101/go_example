package pfcpie

import "fmt"

// TS29.244-g40 8.2.131	SMF Set ID
// The SMF Set ID IE shall contain an FQDN representing the SMF Set. It shall be encoded as shown in Figure 8.2.131-1.

//	                                               Bits
//	 Octets             8       7       6       5       4       3       2       1
//	                +----------------------------------------------------------------+
//	 1 to 2         |                   Type = 180(decimal)                          |
//	                |----------------------------------------------------------------|
//	 3 to 4         |                      Length = n                                |
//	                |----------------------------------------------------------------|
//	   5            |                           Spare                                |
//	                |----------------------------------------------------------------|
//		6 to m	       |                           FQDN                                 |
//	                |----------------------------------------------------------------|
//		(m+1) to (n+4) |   These octet(s) is/are present only if explicitly specified   |
//	                +----------------------------------------------------------------+
//
// Figure 8.2.131-1: SMF Set ID
// FQDN encoding shall be identical to the encoding of a FQDN within a DNS message of clause 3.1 of IETF RFC 1035 [27] but excluding the trailing zero byte.
// NOTE:	The FQDN field in the IE is not encoded as a dotted string as commonly used in DNS master zone files.
type SMFSetID struct {
	FQDN string `json:"fQDN,omitempty"`
}

func (s *SMFSetID) MarshalBinary() (data []byte, err error) {
	spare := uint8(0)
	data = append([]byte(""), spare)
	data = append(data, []byte(s.FQDN)...)
	return data, nil
}

func (s *SMFSetID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie SMFSetID Inadequate TLV length: %d", length)
	}
	idx = idx + 1
	s.FQDN = string(data[idx:])
	return nil
}

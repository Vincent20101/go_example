package pfcpie

import "fmt"

// TS29.244 8.2.154	ATSSS-LL Control Information
//
// The ATSSS-LL Control Information IE shall provide details of the required ATSSS-LL functionality.
// It shall be encoded as shown in Figure 8.2.155-1
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 223(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |-------------------------------------------------------+--------|
//	  5      |             Spare                                     |  LLI   |
//	         |-------------------------------------------------------+--------|
//
// m to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
// -	Bit 1 – LLI: If this bit is set to "1", it indicates that the ATSSS-LL steering functionality is required.
// -	Bit 2 to 8 Spare, for future use and set to "0".
type AtsssLLControlInfo struct {
	Lli bool `json:"lli,omitempty"`
}

func (a *AtsssLLControlInfo) MarshalBinary() (data []byte, err error) {
	data = append(data, byte(btou(a.Lli)))
	return data, nil
}

func (a *AtsssLLControlInfo) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))
	if length < 1 {
		return fmt.Errorf("IE AtsssLLControlInfo Inadequate TLV length: %d", length)
	}
	a.Lli = utob(uint8(data[0]))
	return nil
}

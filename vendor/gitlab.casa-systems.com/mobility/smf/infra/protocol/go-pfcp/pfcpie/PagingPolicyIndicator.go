package pfcpie

import (
	"fmt"
)

// [TS29244 8.2.116 Paging Policy Indicator (PPI) ]
// The Paging Policy Indicator (PPI) IE indicates a PPI value for paging
// policy differentation. It shall be encoded as shown in Figure 8.2.116-1.
//
//	Octets        8       7       6       5       4       3       2       1
//	          +-------------------------------------------------------------+
//	1 to 2    |                   Type = 158(decimal)                       |
//	          |-------------------------------------------------------------|
//	3 to 4    |                      Length = n                             |
//	          |-------------------------------------------------------------|
//	     5    |	Spare	                               | PPI value(bit1,2,3)|
//	          |-------------------------------------------------------------|
//
// 6 to (n+4) |	These octet(s) is/are present only if explicitly specified  |
//
//	+-------------------------------------------------------------+
//
// The definition of Framed-Routing could be found in
// gitlab.casa-systems.com/mobility/smf/infra/protocol/radius/rfc2865, which is uint32 type.
type PagingPolicyIndicator struct {
	PPIValue uint8 `json:"pPIValue,omitempty"`
}

func (f *PagingPolicyIndicator) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := ((f.PPIValue) & 0x7)
	data = append([]byte(""), byte(tmpUint8))
	return data, nil
}

func (f *PagingPolicyIndicator) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie PagingPolicyIndicator Inadequate TLV length: %d", length)
	}

	f.PPIValue = (data[idx] & 0x7)
	idx = idx + 1

	return nil
}

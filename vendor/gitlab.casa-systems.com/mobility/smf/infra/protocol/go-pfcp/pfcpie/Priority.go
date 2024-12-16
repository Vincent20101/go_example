package pfcpie

import "fmt"

const (
	Active uint8 = iota
	Standby
	NoStandby
	High
	Low
)

// TS29.244 8.2.127	Priority
//
// The Priority IE type shall be encoded as shown in Figure 8.2.127-1.
// It indicates whether it is active or standby or no standby for a given
// access when the Steering Mode is set to Active-Standby, or whether it is
// high or low priority for a given access when the Steering Mode is set to Priority-based.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 174(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |-------------------------------+-------+-------+-------+--------|
//	  5      |             Spare             |      Priority Value            |
//	         |-------------------------------+-------+-------+-------+--------|
//
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// Priority Value Values (Decimal)
// - Active:     0
// - Standby:    1
// - No Standby: 2
// - High:       3
// - Low:        4
// - Spare:      5 to 15
type Priority struct {
	PriorityValue uint8 `json:"priorityValue,omitempty"`
}

func (p *Priority) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(p.PriorityValue&0xF))

	return data, nil
}

func (p *Priority) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE Priority Inadequate TLV length: %d", length)
	}
	p.PriorityValue = (uint8(data[idx]) & 0xF)
	idx = idx + 1

	return nil
}

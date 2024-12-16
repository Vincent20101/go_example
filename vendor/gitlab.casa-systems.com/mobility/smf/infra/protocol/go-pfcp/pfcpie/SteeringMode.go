package pfcpie

import "fmt"

const (
	ActiveStandby uint8 = iota
	SmallestDelay
	LoadBalancing
	PriorityBased
)

// TS29.244 8.2.125	Steering Mode
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 171(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |-------------------------------+-------+-------+-------+--------|
//	  5      |             Spare             |     Steering Mode Value        |
//	         |-------------------------------+-------+-------+-------+--------|
//
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// - Steering Mode Value  Values (Decimal)
// - Active-Standby:        0
// - Smallest Delay:        1
// - Load Balancing:        2
// - Priority-based:        3
// - Spare:                 4 to 15
type SteeringMode struct {
	SteeringModeValue uint8 `json:"steeringModeValue,omitempty"`
}

func (s *SteeringMode) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(s.SteeringModeValue&0xF))

	return data, nil
}

func (s *SteeringMode) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie SteeringMode Inadequate TLV length: %d", length)
	}
	s.SteeringModeValue = (uint8(data[idx]) & 0xF)
	idx = idx + 1

	return nil
}

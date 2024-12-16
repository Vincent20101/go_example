package pfcpie

import "fmt"

const (
	SteerFunAtsssLL uint8 = iota
	SteerFunMptcp
)

// TS29.244 8.2.124	Steering Functionality
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 171(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |-------------------------------+-------+-------+-------+--------|
//	  5      |             Spare             | Steering Functionality Value   |
//	         |-------------------------------+-------+-------+-------+--------|
//
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	         +----------------------------------------------------------------+
//
//	Steering Functionality value
//	- ATSSS-LL: 0
//	- MPTCP:    1
//	- Spare:    2 to 15
type SteeringFunctionality struct {
	// octet 5
	SteeringFunctionalityValue uint8 `json:"steeringFunctionalityValue,omitempty"`
}

func (p *SteeringFunctionality) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(p.SteeringFunctionalityValue&0xF))

	return data, nil
}

func (p *SteeringFunctionality) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie SteeringFunctionality Inadequate TLV length: %d", length)
	}
	p.SteeringFunctionalityValue = (uint8(data[idx]) & 0xF)
	idx = idx + 1

	return nil
}

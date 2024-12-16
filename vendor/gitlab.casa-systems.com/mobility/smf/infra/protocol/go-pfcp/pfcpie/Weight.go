package pfcpie

import "fmt"

// TS29.244 8.2.126	Weight
//
// The Weight IE shall be encoded as shown in Figure 8.2.126-1.
// It indicates the percentage of the traffic to be transferred over one access.
//
// The Weight Value field shall take binary coded integer values from 0 up to 100.
// Other values shall be considered as 0.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 173(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = 1                                |
//	         |-------------------------------+-------+-------+-------+--------|
//	  5      |                      Weight Value                              |
//	         +-------------------------------+-------+-------+-------+--------+
type Weight struct {
	WeightValue uint8 `json:"weightValue,omitempty"`
}

func (w *Weight) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(w.WeightValue))

	return data, nil
}

func (w *Weight) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie Weight Inadequate TLV length: %d", length)
	}
	w.WeightValue = uint8(data[idx])
	idx = idx + 1

	return nil
}

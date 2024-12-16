package pfcpie

import (
	"fmt"
)

type TrafficEndpointID struct {
	TrafficEndpointIdValue uint8 `json:"trafficEndpointIdValue,omitempty"`
}

func (t *TrafficEndpointID) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(t.TrafficEndpointIdValue))

	return data, nil
}

func (t *TrafficEndpointID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie TrafficEndpointID Inadequate TLV length: %d", length)
	}
	t.TrafficEndpointIdValue = uint8(data[idx])
	idx = idx + 1

	return nil
}

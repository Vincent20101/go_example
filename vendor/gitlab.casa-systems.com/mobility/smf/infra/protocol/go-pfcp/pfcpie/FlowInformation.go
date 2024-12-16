package pfcpie

import (
	"encoding/binary"
	fmt "fmt"
	math "math"
)

type FlowInformation struct {
	FlowDirection   uint8  `json:"flowDirection,omitempty"`
	FlowDescription []byte `json:"flowDescription,omitempty"`
}

func (f *FlowInformation) MarshalBinary() ([]byte, error) {
	var data []byte
	// Octet 5
	data = append(data, byte(f.FlowDirection))
	// Octet 6...7
	lengthOfFlowDescription := len(f.FlowDescription)
	if lengthOfFlowDescription > math.MaxUint16 {
		return nil, fmt.Errorf("length of flow description overflowed: %d", lengthOfFlowDescription)
	}
	data = append(data, 0, 0)
	binary.BigEndian.PutUint16(data[len(data)-2:], uint16(len(f.FlowDescription)))
	// Octet 8...
	data = append(data, f.FlowDescription...)
	return data, nil
}

func (f *FlowInformation) UnmarshalBinary(data []byte) error {
	length := len(data)

	if len(data) < 1 {
		return fmt.Errorf("IE FlowInformation Inadequate TLV length: %d", length)
	}
	f.FlowDirection = data[0]
	data = data[1:]

	if len(data) < 2 {
		return fmt.Errorf("IE FlowInformation Field FlowDescription Inadequate TLV length: %d", length)
	}
	lengthOfFlowDescription := int(binary.BigEndian.Uint16(data))
	data = data[2:]

	if len(data) < lengthOfFlowDescription {
		return fmt.Errorf("IE FlowInformation Field FlowDescription Inadequate TLV length: %d", length)
	}
	f.FlowDescription = data[:lengthOfFlowDescription]
	return nil
}

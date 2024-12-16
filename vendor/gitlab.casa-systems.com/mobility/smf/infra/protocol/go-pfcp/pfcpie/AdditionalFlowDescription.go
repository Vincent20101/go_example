package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type AdditionalFlowDescriptions struct {
	FlowDescription []byte `json:"flowDescription,omitempty"`
}

func (a *AdditionalFlowDescriptions) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	data = append(data, make([]byte, 2)...)
	binary.BigEndian.PutUint16(data[idx:], uint16(len(a.FlowDescription)))
	data = append(data, a.FlowDescription...)

	return
}

func (a *AdditionalFlowDescriptions) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	if length < idx+2 {
		return fmt.Errorf("IE AdditionalFlowDescriptions Field Fd Inadequate TLV length: %d", length)
	}
	afdLen := binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2

	if length < idx+afdLen {
		return fmt.Errorf("IE AdditionalFlowDescriptions Field FlowDescription Inadequate TLV length: %d, idx: %d, afdLen: %d", length, idx, afdLen)
	}
	a.FlowDescription = append(a.FlowDescription, data[idx:idx+afdLen]...)

	return nil
}

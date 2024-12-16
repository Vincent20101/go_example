package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type PacketDetectionRuleID struct {
	RuleId uint16 `json:"ruleId,omitempty" cmp:"ignore"`
}

func (p *PacketDetectionRuleID) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 6
	data = make([]byte, 2)
	binary.BigEndian.PutUint16(data[idx:], p.RuleId)

	return data, nil
}

func (p *PacketDetectionRuleID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 6
	if length < idx+2 {
		return fmt.Errorf("IE PacketDetectionRuleID Inadequate TLV fail,length: %d", length)
	}
	p.RuleId = binary.BigEndian.Uint16(data[idx:])
	idx = idx + 2

	return nil
}

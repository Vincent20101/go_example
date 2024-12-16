package pfcpie

import (
	"fmt"
	"math/bits"
)

type FailedRuleID struct {
	RuleIdType  uint8  `json:"ruleIdType,omitempty"` // 0x00001111
	RuleIdValue []byte `json:"ruleIdValue,omitempty"`
}

func (f *FailedRuleID) MarshalBinary() (data []byte, err error) {
	// Octet 5
	if bits.Len8(f.RuleIdType) > 4 {
		return []byte(""), fmt.Errorf("Rule ID type shall not be greater than 4 bits binary integer")
	}
	data = append([]byte(""), byte(f.RuleIdType))

	// Octet 6 to p
	data = append(data, f.RuleIdValue...)

	return data, nil
}

func (f *FailedRuleID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE FailedRuleID Inadequate TLV length: %d", length)
	}
	f.RuleIdType = uint8(data[idx]) & Mask4
	idx = idx + 1

	// Octet 6 to p
	f.RuleIdValue = data[idx:]

	return nil
}

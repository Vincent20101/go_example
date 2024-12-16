package pfcpie

import (
	"encoding/binary"
	"errors"
	"fmt"
)

const (
	CasaEdrFieldsPdrId            = 1
	CasaEdrFieldsSnRuleBase       = 2
	CasaEdrFieldsSnRuleDefName    = 3
	CasaEdrFieldsSnChargingAction = 4
	CasaEdrFieldsSnRatingGroup    = 5
)

// CasaEdrFields <FGC-3472> smf should enhance casa private IE to support Claro EDR fields
type CasaEdrFields struct {
	EnterpriseId   uint16           `json:"enterpriseId,omitempty"`
	PdrFilterRules []*PdrFilterRule `json:"pdrFilterRules,omitempty"`
}

func (c *CasaEdrFields) MarshalBinary() ([]byte, error) {
	// Write enterpriseId and numberOfPdr
	var data []byte
	enterpriseId := uint16(CasaEnterpriseID)
	data = AppendUint16(data, enterpriseId)

	num := uint8(len(c.PdrFilterRules))
	if num == 0 {
		return nil, fmt.Errorf("PdrFilterRules list is empty")
	}
	data = append(data, num)

	// Write PdrFilterRules
	for _, rule := range c.PdrFilterRules {
		ruleData, err := rule.MarshalBinary()
		if err != nil {
			return nil, err
		}
		data = append(data, ruleData...)
	}
	return data, nil
}
func (c *CasaEdrFields) UnmarshalBinary(data []byte) error {
	// Read enterpriseId and numberOfPdr
	c.EnterpriseId = binary.BigEndian.Uint16(data[0:2])
	numOfRules := int(data[2])
	offset := 3

	for i := 0; i < numOfRules; i++ {
		rule := &PdrFilterRule{}

		// Read PdrFilterRule data
		if offset >= len(data) {
			return errors.New("data is too short")
		}
		if err := rule.UnmarshalBinary(data[offset:]); err != nil {
			return err
		}
		c.PdrFilterRules = append(c.PdrFilterRules, rule)

		// Update offset
		offset += rule.Length()
		if offset > len(data) {
			return errors.New("data is too short")
		}
	}

	return nil
}

func (c *CasaEdrFields) Length() int {
	length := 0

	for _, rule := range c.PdrFilterRules {
		length += rule.Length()
	}

	return length
}
func (c *CasaEdrFields) New() *CasaEdrFields {
	return &CasaEdrFields{
		EnterpriseId: CasaEnterpriseID,
	}
}

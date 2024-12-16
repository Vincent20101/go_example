package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// PdrFilterRule
// CasaEdrFields use
/*
   //	                                Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	 1 to 3  |                      PdrId                                    |
//	         |---------------------------------------------------------------|
//	 4 to n  |                      SnRuleBase                               |
//	         |---------------------------------------------------------------|
//  n+1 to z |                      SnRuleDefName                            |
//	         |---------------------------------------------------------------|
//  z+1 to x |                      SnChargingAction                         |
//	         |---------------------------------------------------------------|
//x+1 to x+5 |                      SnRatingGroup                            |
//	         |---------------------------------------------------------------|
*/

type PdrFilterRule struct {
	PdrId uint16 `json:"pdrId,mandatory"`
	//PCC RULE Name and belongs rulebases
	SnRuleBase    string `json:"snRuleBase,omitempty"`
	SnRuleDefName string `json:"snRuleDefName,omitempty"`
	//Charging Data Name and ratingGroup
	SnChargingAction string `json:"snChargingAction,omitempty"`
	SnRatingGroup    uint32 `json:"snRatingGroup,omitempty"`
}

func (p *PdrFilterRule) MarshalBinary() ([]byte, error) {
	var data []byte
	// PdrId
	tmp := make([]byte, 3)
	tmp[0] = CasaEdrFieldsPdrId
	binary.BigEndian.PutUint16(tmp[1:], p.PdrId)
	data = append(data, tmp...)

	// SnRuleBase
	if p.SnRuleBase != "" {
		snRuleBaseBuf := []byte(p.SnRuleBase)
		snRuleBaseLen := byte(len(snRuleBaseBuf))
		if snRuleBaseLen > 255 {
			return nil, fmt.Errorf("snRuleBaseLen is too large: %d, should be less than 256", snRuleBaseLen)
		}
		tmp := make([]byte, 2)
		tmp[0] = CasaEdrFieldsSnRuleBase
		tmp[1] = snRuleBaseLen
		data = append(data, tmp...)
		data = append(data, snRuleBaseBuf...)
	}

	// SnRuleDefName
	if p.SnRuleDefName != "" {
		snRuleDefNameBuf := []byte(p.SnRuleDefName)
		snRuleDefNameLen := byte(len(snRuleDefNameBuf))
		if snRuleDefNameLen > 255 {
			return nil, fmt.Errorf("snRuleDefNameLen is too large: %d, should be less than 256", snRuleDefNameLen)
		}
		tmp := make([]byte, 2)
		tmp[0] = CasaEdrFieldsSnRuleDefName
		tmp[1] = snRuleDefNameLen
		data = append(data, tmp...)
		data = append(data, snRuleDefNameBuf...)
	}

	// SnChargingAction
	if p.SnChargingAction != "" {
		snChargingActionBuf := []byte(p.SnChargingAction)
		snChargingActionLen := byte(len(snChargingActionBuf))
		if snChargingActionLen > 255 {
			return nil, fmt.Errorf("snChargingActionLen is too large: %d, should be less than 256", snChargingActionLen)
		}
		tmp := make([]byte, 2)
		tmp[0] = CasaEdrFieldsSnChargingAction
		tmp[1] = snChargingActionLen
		data = append(data, tmp...)
		data = append(data, snChargingActionBuf...)
	}

	// SnRatingGroup
	tmp2 := make([]byte, 5)
	tmp2[0] = CasaEdrFieldsSnRatingGroup
	binary.BigEndian.PutUint32(tmp2[1:], p.SnRatingGroup)
	data = append(data, tmp2...)

	return data, nil

}

func (p *PdrFilterRule) UnmarshalBinary(data []byte) error {
	for len(data) >= 2 {
		fieldType := data[0]
		var fieldLen uint16

		switch fieldType {
		case CasaEdrFieldsPdrId:
			fieldLen = 2
			if len(data) < int(fieldLen+1) {
				return fmt.Errorf("not enough data for PDR ID")
			}
			p.PdrId = binary.BigEndian.Uint16(data[1:3])
			data = data[fieldLen+1:]

		case CasaEdrFieldsSnRuleBase, CasaEdrFieldsSnRuleDefName, CasaEdrFieldsSnChargingAction:
			fieldLen = uint16(data[1])
			if len(data) < int(fieldLen+2) {
				return fmt.Errorf("not enough data for field type %d", fieldType)
			}
			fieldData := data[2 : 2+fieldLen]
			fieldStr := string(fieldData)
			switch fieldType {
			case CasaEdrFieldsSnRuleBase:
				p.SnRuleBase = fieldStr
			case CasaEdrFieldsSnRuleDefName:
				p.SnRuleDefName = fieldStr
			case CasaEdrFieldsSnChargingAction:
				p.SnChargingAction = fieldStr
			}
			data = data[fieldLen+2:]

		case CasaEdrFieldsSnRatingGroup:
			fieldLen = 4
			if len(data) < int(fieldLen+1) {
				return fmt.Errorf("not enough data for SN rating group")
			}
			p.SnRatingGroup = binary.BigEndian.Uint32(data[1:5])
			data = data[fieldLen+1:]
			return nil

		default:
			return fmt.Errorf("unknown field type %d", fieldType)
		}
	}

	return nil
}

func (p *PdrFilterRule) Length() int {
	length := 0

	// p.PdrId
	length += 3

	if p.SnRuleBase != "" {
		length += 2 + len(p.SnRuleBase)
	}

	if p.SnRuleDefName != "" {
		length += 2 + len(p.SnRuleDefName)
	}

	if p.SnChargingAction != "" {
		length += 2 + len(p.SnChargingAction)
	}

	// p.SnRatingGroup
	length += 5

	return length
}

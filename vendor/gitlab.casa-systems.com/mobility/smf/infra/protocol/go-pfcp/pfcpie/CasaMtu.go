package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// GCS-25579: smf should support the Private IE about MTU

type CasaMtu struct {
	EnterpriseId uint16 `json:"enterpriseId,omitempty"`
	MtuValue     uint16 `json:"mtuValue,omitempty"`
}

func NewCasaMtu() *CasaMtu {
	return &CasaMtu{
		EnterpriseId: CasaEnterpriseID,
	}
}

func (c *CasaMtu) MarshalBinary() ([]byte, error) {
	var data []byte
	// enterpriseId
	data = AppendUint16(data, c.EnterpriseId)
	// mtu value
	data = AppendUint16(data, c.MtuValue)
	return data, nil
}

func (c *CasaMtu) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	if length < idx+2 {
		return fmt.Errorf("IE CasaMtu Inadequate TLV length: %d", length)
	}
	c.EnterpriseId = binary.BigEndian.Uint16(data[idx : idx+2])

	idx = idx + 2
	if length < idx+2 {
		return fmt.Errorf("IE CasaMtu Inadequate TLV length: %d", length)
	}
	c.MtuValue = binary.BigEndian.Uint16(data[idx : idx+2])

	return nil
}

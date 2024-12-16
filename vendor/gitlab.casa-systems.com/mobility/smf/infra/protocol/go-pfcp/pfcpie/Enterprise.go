package pfcpie

import (
	"fmt"
)

// Enterprise
// Casa private IE
//
//      Octets              8       7       6       5       4       3       2       1
//                      +---------------------------------------------------------------+
//      1 to 2          |                    Type = 32777(decimal)                      |
//                      |---------------------------------------------------------------|
//      3 to 4          |                        Length = n                             |
//                      |---------------------------------------------------------------|
//      5 to 6          |                    Enterprise ID = 20858                      |
//                      |---------------------------------------------------------------|
//      7 to 7+n-3      |                     Enterprise Ie Data                        |
//                      |---------------------------------------------------------------|
//
//  Enterprise Ie Data:
//      Different scenarios have different values, for example:
//          1. jira GCS-8626: SMF+PWG-C need to provide EBI to UPF/PGW-U when pfcp establishment request msg for 4G session.
//

// Enterprise is a self-defined IE, not shown in reference.  The tlv tag is
// from OID(Object Identifier Descriptors) of Casa Systerm Inc. which is
// 1.3.6.1.4.1.20858.
type Enterprise struct {
	EnterpriseId     uint16 `json:"enterpriseId,omitempty"`
	EnterpriseIeData []byte `json:"enterpriseIeData,omitempty"`
}

func NewEnterprise() *Enterprise {
	return &Enterprise{
		EnterpriseId: CasaEnterpriseID,
	}
}

func (e *Enterprise) MarshalBinary() (data []byte, err error) {
	// oct 5 to 6
	data = AppendUint16(data, e.EnterpriseId)
	// oct (7) to (7+n-3)
	if len(e.EnterpriseIeData) != 0 {
		EnterpriseIeData := make([]byte, len(e.EnterpriseIeData))
		copy(EnterpriseIeData, e.EnterpriseIeData)
		data = append(data, EnterpriseIeData...)
	}

	return data, nil
}

func (e *Enterprise) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)

	// EnterpriseId
	enterpriseId, err := byteReader.ReadUint16()
	if err != nil {
		return fmt.Errorf("Read EnterpriseId failed, err: %v", e)
	}
	e.EnterpriseId = enterpriseId

	// EnterpriseIeData
	if len(data[2:]) > 0 {
		e.EnterpriseIeData = make([]byte, len(data[2:]))
		copy(e.EnterpriseIeData, data[2:])
	}

	return nil
}

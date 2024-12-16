package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// FlowInactivityTimer IE.
//
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                        Type = 32794(decimal)                   |
//           |----------------------------------------------------------------|
//  3 to 4   |                          Length = n                            |
//           |----------------------------------------------------------------|
//  5 to 6   |                        Enterprise ID                           |
//           |----------------------------------------------------------------|
// 7 to 10   |                       Flow Inactivity Timer Value              |
//           |----------------------------------------------------------------|
//

type FlowInactivityTimer struct {
	EnterpriseID   uint16 `json:"enterpriseID,omitempty"`
	FlowTimerValue uint32 `json:"flowTimerValue,omitempty"`
}

func NewFlowInactivityTimer() *FlowInactivityTimer {
	return &FlowInactivityTimer{
		EnterpriseID:   CasaEnterpriseID,
		FlowTimerValue: 0,
	}
}

func (u *FlowInactivityTimer) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 6)
	// Enterprise ID
	binary.BigEndian.PutUint16(data, u.EnterpriseID)
	// Octet 7 to 10
	var idx uint16 = 2
	binary.BigEndian.PutUint32(data[idx:], u.FlowTimerValue)

	return data, nil
}

func (u *FlowInactivityTimer) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)

	// EnterpriseID
	eid, e := byteReader.ReadUint16()
	if e != nil {
		return fmt.Errorf("Inadequate length for EnterpriseId, %v", e)
	}
	u.EnterpriseID = eid

	// Pool Name
	timerVal, e := byteReader.ReadUint32()
	if e != nil {
		return fmt.Errorf("Inadequate length for Flow Timer value,  %v", e)
	}
	u.FlowTimerValue = timerVal

	return nil
}

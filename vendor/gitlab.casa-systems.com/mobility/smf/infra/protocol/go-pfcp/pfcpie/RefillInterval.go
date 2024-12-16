package pfcpie

import "fmt"

// RefillInterval implements Refill Interval IE.
// IE Type: 0x8017
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                        Type = 32798(decimal)                   |
//	         |----------------------------------------------------------------|
//	3 to 4   |                        Length = 2+4=6                          |
//	         |----------------------------------------------------------------|
//	5 to 6   |                        Enterprise ID                           |
//	         |----------------------------------------------------------------|
//	7 to 10  |                        Refill interval in milliseconds         |
//	         +----------------------------------------------------------------+
type RefillInterval struct {
	EnterpriseID   uint16 `json:"enterpriseID,omitempty"`
	RefillInterval uint32 `json:"refillInterval,omitempty"`
}

// NewRefillInterval return a new RefillInterval.
func NewRefillInterval() *RefillInterval {
	return &RefillInterval{
		EnterpriseID:   EricssonEnterpriseID,
		RefillInterval: 0,
	}
}

// MarshalBinary marshals RefillInterval into octets.
func (v *RefillInterval) MarshalBinary() (data []byte, err error) {
	data = AppendUint16([]byte(""), v.EnterpriseID)
	data = AppendUint32(data, v.RefillInterval)
	return data, nil
}

// UnmarshalBinary decodes octets to RefillInterval.
func (v *RefillInterval) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)

	// EnterpriseID
	eid, e := byteReader.ReadUint16()
	if e != nil {
		return fmt.Errorf("inadequate length for Enterprise ID: %v", e)
	}
	v.EnterpriseID = eid

	// Action
	refillInterval, e := byteReader.ReadUint32()
	if e != nil {
		return fmt.Errorf("inadequate length for RefillInterval: %v", e)
	}
	v.RefillInterval = refillInterval
	return nil
}

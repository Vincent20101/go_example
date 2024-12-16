package pfcpie

import "fmt"

// TS29.244 8.2.153	Access Availability Information
//
// The Access Availability Information IE shall indicate an access type and whether the access type
// has become available or not available. It shall be encoded as shown in Figure 8.2.153-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 219(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      |             Spare             |Availability Status| Access Type|
//	         |----------------------------------------------------------------|
//
// m to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The Access Type shall be encoded as a 2 bits binary integer as specified in Table 8.2.153-1.
// Access Type       Values (Decimal)
// 3GPP access type:     0
// Non-3GPP access type: 1
// Spare:                2 to  3
//
// The Availability Status shall be encoded as a 2 bits binary integer as specified in Table 8.2.153-2.
//
//	Availability Status					Values (Decimal)
//	Access has become unavailable			0
//	Access has become available				1
//	Spare									2 to 3
type AccessAvailabilityInfo struct {
	AccessType         uint8 `json:"accessType,omitempty"`
	AvailabilityStatus uint8 `json:"availabilityStatus,omitempty"`
}

func (a *AccessAvailabilityInfo) MarshalBinary() (data []byte, err error) {

	octet5 := a.AccessType | a.AvailabilityStatus<<2
	data = append(data, byte(octet5))

	return data, nil
}

func (a *AccessAvailabilityInfo) UnmarshalBinary(data []byte) error {
	length := uint8(len(data))
	if length < 1 {
		return fmt.Errorf("IE AccessAvailabilityInfo Inadequate TLV length: %d", length)
	}

	a.AccessType = uint8(data[0] & 0x03)
	a.AvailabilityStatus = uint8((data[0] & 0x0c) >> 2)

	return nil
}

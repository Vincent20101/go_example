package pfcpie

import (
	"encoding/binary"
	"fmt"
	"math"
)

const (
	// Octet 5
	AdditionalOuterHeaderCreationStag1 uint8 = 1
	AdditionalOuterHeaderCreationStag2 uint8 = 1 << 1
)

// AdditionalOuterHeaderCreation IE: verizon custom IE
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                       Type = 32795(decimal)                    |
//	         |----------------------------------------------------------------|
//	3 to 4   |                          Length = n                            |
//	         |----------------------------------------------------------------|
//	5 to 6   |                			Enterprise ID           		      |
//	         |----------------------------------------------------------------|
//	7 to 8	 |          Additional Outer Header Creation Description          |
//	         |----------------------------------------------------------------|
//
// t to t+2  |                           S-TAG1                               |
//
//	|----------------------------------------------------------------|
//
// u to u+2  |                           S-TAG2                               |
//
//	|----------------------------------------------------------------|
//
// s to (n+4)|   These octets(s) is/are present only if explicity specified   |
type AdditionalOuterHeaderCreation struct {
	EnterpriseId                             uint16 `json:"enterpriseId,omitempty"`
	AdditionalOuterHeaderCreationDescription uint8  `json:"outerHeaderCreationDescription,omitempty"`
	Stag1                                    *STAG1 `tlv:"32793" json:"stag1,omitempty"`
	Stag2                                    *STAG2 `tlv:"32794" json:"stag2,omitempty"`
}

// MarshalBinary marshals OuterHeaderCreation, Refer to [29244 8.2.56 Outer
// Header Creation].
func (o *AdditionalOuterHeaderCreation) MarshalBinary() (data []byte, err error) {
	errMsgPrefix := "Marshal VerizonOuterHeaderCreation failed: "
	// Refer to [29244 8.2.56 Outer Header Creation], At least one bit of the
	// Outer Header Creation Description field shall be set to "1". Bits 5/1 and
	// 5/2 may both be set to "1".
	if o.AdditionalOuterHeaderCreationDescription&0x03 == 0 {
		return []byte(""), fmt.Errorf(errMsgPrefix+"At least one bit of OuterHeaderCreationDescription field shall be set%x, ", o.AdditionalOuterHeaderCreationDescription&0x03)
	}
	idx := 0
	data = make([]byte, 2)
	enterpriseId := uint16(VerizonSpecificUEPDNInfo_ENTERPRISEID)
	binary.BigEndian.PutUint16(data, enterpriseId)
	idx += 2

	octet5 := o.AdditionalOuterHeaderCreationDescription & Mask2
	Stag1Flag := utob(octet5 & BitMask1)
	Stag2Flag := utob(octet5 & BitMask2)

	// Octet t to (t+2)
	var stagData []byte
	if Stag1Flag {
		if o.Stag1 != nil {
			var err error
			stagData, err = o.Stag1.MarshalBinary()
			if err != nil {
				return nil, fmt.Errorf("VerizonOuterHeaderCreation failed to marshal STAG1: %v", err)
			}
			if len(stagData) > math.MaxUint8 {
				return nil, fmt.Errorf("STAG1 length overflowed: %d", len(stagData))
			}
			data = append(data, stagData...)
		}
	}
	if Stag2Flag {
		// Octet u to (u+2)
		if o.Stag2 != nil {
			var err error
			stagData, err = o.Stag2.MarshalBinary()
			if err != nil {
				return nil, fmt.Errorf("VerizonOuterHeaderCreation failed to marshal STAG2: %v", err)
			}
			if len(stagData) > math.MaxUint8 {
				return nil, fmt.Errorf("STAG2 length overflowed: %d", len(stagData))
			}
			data = append(data, stagData...)
		}
	}
	return data, nil
}

func (o *AdditionalOuterHeaderCreation) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	var idx uint16 = 0
	// Octet 3 to 4
	if length < idx+2 {
		return fmt.Errorf("IE VerizonOuterHeaderCreation Inadequate TLV length: %d", length)
	}

	//Octet 5 to 6
	o.EnterpriseId = binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2

	// Octet 7 to 8
	o.AdditionalOuterHeaderCreationDescription = data[idx]

	octet5 := uint8(o.AdditionalOuterHeaderCreationDescription & Mask2)
	idx = idx + 1

	Stag1Flag := utob(octet5 & BitMask1)
	Stag2Flag := utob(octet5 & BitMask2)

	// Octet t to (t+2)
	if Stag1Flag {
		if length < idx+3 {
			return fmt.Errorf("IE VerizonOuterHeaderCreation Fieid STAG1 Inadequate TLV fail,length: %d", length)
		}
		var stag1 STAG1
		if err := stag1.UnmarshalBinary(data[idx : idx+3]); err != nil {
			return fmt.Errorf("IE VerizonOuterHeaderCreation failed to unmarshal STAG1: %v", err)
		}
		o.Stag1 = &stag1
		idx = idx + 3
	}
	// Octet u to (u+2)
	if Stag2Flag {
		if length < idx+3 {
			return fmt.Errorf("IE VerizonOuterHeaderCreation Fieid STAG2 Inadequate TLV fail,length: %d", length)
		}
		var stag2 STAG2
		if err := stag2.UnmarshalBinary(data[idx : idx+3]); err != nil {
			return fmt.Errorf("IE VerizonOuterHeaderCreation failed to unmarshal STAG2: %v", err)
		}
		o.Stag2 = &stag2
		idx = idx + 3
	}
	return nil
}

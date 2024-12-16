package pfcpie

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

const (
	noSDValueAssociatedWithTheSST string = "ffffff"
)

// S-NSSAI IE. Refer to [TS29244 8.2.176 S-NSSAI]
//
// The S-NSSAI IE indicates the S-NSSAI of a PDU session. It shall be encoded
// as shown in Figure 8.2.176-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                       Type = 257(decimal)                     |
//	         |---------------------------------------------------------------|
//	3 to 4   |                           Length = n                          |
//	         |---------------------------------------------------------------|
//	  5      |                              SST                              |
//	         |---------------------------------------------------------------|
//	6 to 8   |                              SD                               |
//	         +---------------------------------------------------------------+
//
// The SST (Slice/Service Type) and SD (Slice Differentiator) fields shall be
// encoded as defined in clause 28.4.2 of 3GPP TS 23.003 [2].
//
// ============================================================================
// [TS23003 28.4.2 Format of the S-NSSAI]
//
// The structure of the S-NSSAI is depicted in Figure 28.4.2-1
//
//	       SST                                SD
//	+-----------------+-----------------------------------------------------+
//	|     8 bit       |                      24 bit                         |
//	+-----------------+-----------------------------------------------------+
//	               Figure 28.4.2-1: Structure of S-NSSAI
//
// The S-NSSAI may include both the SST and SD fields (in which case the
// S-NSSAI length is 32 bits in total), or the S-NSSAI may just include the
// SST field (in which case the S-NSSAI length is 8 bits only).
//
// The SST field may have standardized and non-standardized values. Values 0 to
// 127 belong to the standardized SST range and they are defined in 3GPP TS
// 23.501 [119]. Values 128 to 255 belong to the Operator-specific range.
//
// The SD field has a reserved value "no SD value associated with the SST"
// defined as hexadecimal FFFFFF. In certain protocols, the SD field is not
// included to indicate that no SD value is associated with the SST.
//
// ============================================================================
// [TS29571 5.4.4.2 Type: Snssai]
//
// sst | Uinterger | Mandatory
// Unsigned integer, within the range 0 to 255, representing the Slice/Service
// Type. It indicates the expected Network Slice behaviour in terms of features
// and services. Values 0 to 127 correspond to the standardized SST range.
// Values 128 to 255 correspond to the Operator-specific range. See clause
// 28.4.2 of 3GPP TS 23.003 [7]. Standardized values are defined in clause
// 5.15.2.2 of 3GPP TS 23.501 [8]
//
// sd | string | Optional
// 3-octet string, representing the Slice Differentiator, in hexadecimal
// representation. Each character in the string shall take a value of "0" to
// "9" or "A" to "F" and shall represent 4 bits. The most significant character
// representing the 4 most significant bits of the SD shall appear first in the
// string, and the character representing the 4 least significant bit of the SD
// shall appear last in the string.
type SNSSAI struct {
	Sst uint8 `json:"sst,omitempty"` // slice/Service Type

	// If "no SD value associated with the SST", Set sd to "",
	// then we'll encoded this field as 0xffffff.
	// Refer to [TS23003 28.4.2 Format of the S-NSSAI]
	Sd string `json:"sd,omitempty"` // Slice Differentiator
}

// When Snssai needs to be converted to string (e.g. when used in maps as key),
// the string shall be composed of one to three digits "sst" optionally
// followed by "-" and 6 hexadecimal digits "sd", and shall match the following
// pattern:
//
// ^([0-9]|[1-9][0-9]|1[0-9][0-9]|2([0-4][0-9]|5[0-5]))(-[A-Fa-f0-9]{6})?$
//
// Example 1: "255-19CDE0"
// Example 2: "29"
func (s *SNSSAI) String() string {
	if s.Sd != "" {
		return fmt.Sprintf("%d-%s", s.Sst, s.Sd)
	}
	return fmt.Sprintf("%d", s.Sst)
}

func (s *SNSSAI) MarshalBinary() (data []byte, err error) {
	// Reserved value "no SD value associated with the SST".
	if s.Sd == "" {
		s.Sd = noSDValueAssociatedWithTheSST
	}

	// Sd field should be 6 digits. Change to null if it is not.
	if len(s.Sd) != 6 {
		//TODO - Log error for debuggability
		//message := "Sd MarshalBinary failed: Sd field of SNSSAI is " + s.Sd +
		//	", which should be 6 hexadecimal digits string."
		s.Sd = noSDValueAssociatedWithTheSST
	}

	// convert string SD to integer.
	sd, err := strconv.ParseUint(s.Sd, 16, 24)
	if err != nil {
		//TODO - Log error for debuggability
		//message := "Set Sd of SNSSAI failed: cannot parse " +
		//	s.Sd + " as a 6 hexadecimal digits.: %s"
		// If there was error in parsing, set sd to well defined nil value
		sd, _ = strconv.ParseUint(noSDValueAssociatedWithTheSST, 16, 24)
	}

	// Octet 6 to 8: SD
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, uint32(sd))

	// Octet 5: SST
	data[0] = byte(s.Sst)

	return data, nil
}

func (s *SNSSAI) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	if length != 4 {
		return fmt.Errorf("unmarshal SNSSAI failed: Length of body SNSSAI is not 4, got %d", length)
	}
	s.Sst = uint8(data[0])
	data[0] = 0
	sd := binary.BigEndian.Uint32(data)
	s.Sd = fmt.Sprintf("%06x", sd)
	return nil
}

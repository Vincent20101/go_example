package pfcpie

import (
	"encoding/binary"
	"fmt"
)

const (
	CpFunctionFeaturesLoad  uint8 = 0x1
	CpFunctionFeaturesOvrl  uint8 = 0x2
	CpFunctionFeaturesEpfar uint8 = 0x4
	CpFunctionFeaturesSset  uint8 = 0x8
)

// CPFunctionFeatures IE, Refer to [TS29244 8.2.58 CP Function Features]
//
// The CP Function Features IE indicates the features supported by the CP
// function. Only features having an impact on the (system-wide) UP function
// behaviour are signalled in this IE. It is coded as depicted in Figure
// 8.2.58-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                       Type = 89(decimal)                      |
//	         |---------------------------------------------------------------|
//	3 to 4   |                           Length = n                          |
//	         |---------------------------------------------------------------|
//	  5      |                       Supported-Features                      |
//	         |---------------------------------------------------------------|
//	6 to 7   |                  Additional Supported-Features 1              |
//	         |---------------------------------------------------------------|
//	8 to n+4 |   These octets(s) is/are present only if explicity specified  |
//	         +---------------------------------------------------------------+
//	                      Figure 8.2.58-1: CP Function Features
//
// The CP Function Features IE takes the form of a bitmask where each bit set
// indicates that the corresponding feature is supported. Spare bits shall be
// ignored by the receiver. The same bitmask is defined for all PFCP
// interfaces.
//
// The following table specifies the features defined on PFCP interfaces and
// the interfaces on which they apply.
//
//	Table 8.2.58-1: CP Function Features
//
//	       [...]
//
// The table is not shown here, please refer to the spec for more detail.
type CPFunctionFeatures struct {

	// Octets 5: Supported-Features.

	Load  bool `json:"load,omitempty"`  // Octet5/bit1
	Ovrl  bool `json:"ovrl,omitempty"`  // Octet5/bit2
	Epfar bool `json:"epfar,omitempty"` // Octet5/bit3
	Sset  bool `json:"sset,omitempty"`  // Octet5/bit4
	Bundl bool `json:"bundl,omitempty"` // Octet5/bit5
	Mpas  bool `json:"mpas,omitempty"`  // Octet5/bit6
	Ardr  bool `json:"ardr,omitempty"`  // Octet5/bit7

	// Octets 6 to 8 Additional Supported-Features 1
	AdditionalSupportedFeatures1 uint16 `json:"additionalSupportedFeatures1,omitempty"`
}

func (c *CPFunctionFeatures) GetSupportedFeatures() uint8 {
	// Octets 5: Supported-Features
	tmpoctet := btou(c.Load) |
		btou(c.Ovrl)<<1 |
		btou(c.Epfar)<<2 |
		btou(c.Sset)<<3 |
		btou(c.Bundl)<<4 |
		btou(c.Mpas)<<5 |
		btou(c.Ardr)<<6

	return tmpoctet
}

func (c *CPFunctionFeatures) MarshalBinary() (data []byte, err error) {
	// Octets 5: Supported-Features
	tmpoctet := btou(c.Load) |
		btou(c.Ovrl)<<1 |
		btou(c.Epfar)<<2 |
		btou(c.Sset)<<3 |
		btou(c.Bundl)<<4 |
		btou(c.Mpas)<<5 |
		btou(c.Ardr)<<6
	data = append([]byte(""), byte(tmpoctet))

	tmp16 := make([]byte, 2)
	binary.BigEndian.PutUint16(tmp16, c.AdditionalSupportedFeatures1)
	data = append(data, tmp16...)
	return data, nil
}

func (c *CPFunctionFeatures) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0

	// Octets 5 to 6: Supported-Features
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Umarshal CPFunctionFeatures failed, inadequate TLV length: %d, idx: %d", length, idx)
	}

	c.Load = utob(uint8(data[idx]) & BitMask1)
	c.Ovrl = utob(uint8(data[idx]) & BitMask2)
	c.Epfar = utob(uint8(data[idx]) & BitMask3)
	c.Sset = utob(uint8(data[idx]) & BitMask4)
	c.Bundl = utob(uint8(data[idx]) & BitMask5)
	c.Mpas = utob(uint8(data[idx]) & BitMask6)
	c.Ardr = utob(uint8(data[idx]) & BitMask7)

	idx = idx + 1

	// Octets 6 to 7: Additional Supported-Features 1
	if length < idx+2 {
		return nil
		//return fmt.Errorf("Umarshal CPFunctionFeatures failed, inadequate TLV length: %d", length)
	}
	c.AdditionalSupportedFeatures1 = binary.BigEndian.Uint16(data[idx:])
	if length != idx {
		return nil
		// Octets(s) is/are present only if explicity specified
	}
	return nil
}

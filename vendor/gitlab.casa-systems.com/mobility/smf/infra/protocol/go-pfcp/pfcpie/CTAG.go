package pfcpie

import (
	"fmt"
)

// TS29.244 8.2.94 C-TAG (Customer-VLAN tag)
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 134(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |---------------------------------------+-------+-------+--------|
//	  5      |             Spare                     |  VID  |  DEI  |  PCP   |
//	         |-------------------------------+-------+-------+-------+--------|
//	  6      |        C-VID Value            | DEI   |       PCP Value        |
//	         |                               | Flag  |                        |
//	         |-------------------------------+-------+------------------------|
//	  7      |                           C-VID Value                          |
//	         |----------------------------------------------------------------|
//
// 8 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//   - Bit 1 – PCP: If this bit is set to "1", then PCP Value field shall
//     used by the receiver, otherwise the PCP Value  field shall be ignored.
//   - Bit 2 – DEI: If this bit is set to "1", then DEI flag field shall
//     used by the receiver, otherwise the DEI flag field  shall be ignored.
//   - Bit 3 – VID: If this bit is set to "1", then C-VID value field shall
//     used by the receiver, otherwise the VID Value  field shall be ignored.
//   - Bit 4 to 8 – spare and reserved for future use.
//
// The PCP value, DEI flag and C-VID Value are specified in IEEE 802.1Q [30]
// tag format.
//
// Octet 6 / Bit 3 shall contain the most significant bit of the PCP value.
//
// Octet 6 / Bit 8 shall be the most significant bit of the C-VID value and
// Octet 7 / Bit 1 shall be the least significant bit  (see clause 7.1).
//
// NOTE: The encoding of the C-Tag in PFCP differs from the encoding of the
// C-Tag defined in IEEE 802.1Q [30].
type CTAG struct {
	// octet 5
	Pcp bool `json:"pcp,omitempty"`
	Dei bool `json:"dei,omitempty"`
	Vid bool `json:"vid,omitempty"`
	// octet 6 to 7
	PcpValue  uint8  `json:"pcpValue,omitempty"`
	DeiFlag   uint8  `json:"deiFlag,omitempty"`
	CvidValue uint16 `json:"cvidValue,omitempty"`
}

func (c *CTAG) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0

	// Octet 5
	octet5 := btou(c.Pcp) |
		btou(c.Dei)<<1 |
		btou(c.Vid)<<2
	data = append([]byte(""), byte(octet5))
	idx = idx + 1

	// Octet 6
	octet6 := (c.PcpValue & 0x07) |
		((c.DeiFlag & 0x01) << 3) |
		(uint8((c.CvidValue & 0x0F00) >> 4))
	data = append(data, byte(octet6))
	idx = idx + 1

	// Octet 7
	data = append(data, byte(uint8(c.CvidValue&0xFF)))
	idx = idx + 1

	return data, nil
}

func (c *CTAG) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	var idx uint16 = 0

	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Unmarshal IE CTAG failed: empty body")
	}
	c.Pcp = utob(uint8(data[idx]) & BitMask1)
	c.Dei = utob(uint8(data[idx]) & BitMask2)
	c.Vid = utob(uint8(data[idx]) & BitMask3)
	idx = idx + 1

	// Octet 6
	if length < idx+1 {
		return fmt.Errorf("Unmarshal IE CTAG Octet6 failed, Inadequate body length, idx: %d, length: %d", idx, length)
	}
	c.PcpValue = uint8(data[idx] & 0x07)
	c.DeiFlag = uint8((data[idx] >> 3) & 0x01)
	cvid_high := uint8((data[idx] >> 4) & 0x0F)
	idx = idx + 1

	// Octet 7
	if length < idx+1 {
		return fmt.Errorf("Unmarshal IE:[CTAG] field:[C-VID Value] failed, Inadequate body length, idx: %d, length: %d", idx, length)
	}
	c.CvidValue = (uint16(cvid_high) << 8) | (uint16(data[idx]) & 0xff)
	return nil
}

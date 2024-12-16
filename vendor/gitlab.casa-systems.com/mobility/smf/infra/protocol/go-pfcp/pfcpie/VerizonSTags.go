package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// STAG1 - Verizon Custom IE based on TS29.244 8.2.95 S-TAG (Service-VLAN tag)
//
// The STAG1 IE type shall be encoded as shown in below figure. It shall
// contain Service-VLAN tag (S-TAG) as  defined in IEEE 802.1Q [30]
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 32793(decimal)                        |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	5 to 6   |                			Enterprise ID           		      |
//	         |----------------------------------------------------------------|
//	  7      |             Spare                     |  VID  |  DEI  |  PCP   |
//	         |----------------------------------------------------------------|
//	  8      |        S-VID Value            | DEI   |       PCP Value        |
//	         |                               | Flag  |                        |
//	         |----------------------------------------------------------------|
//	  9      |                           S-VID Value                          |
//	         |----------------------------------------------------------------|
//
// 10 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 7:
//   - Bit 1 – PCP: If this bit is set to "1", then PCP Value field shall
//     used by the receiver, otherwise the PCP Value  field shall be ignored.
//   - Bit 2 – DEI: If this bit is set to "1", then DEI flag field shall
//     used by the receiver, otherwise the DEI flag field  shall be ignored.
//   - Bit 3 – VID: If this bit is set to "1", then S-VID value field shall
//     used by the receiver, otherwise the VID Value  field shall be ignored.
//   - Bit 4 to 8 – spare and reserved for future use.
//
// The PCP value, DEI flag and S-VID Value are specified in IEEE 802.1Q [30]
// tag format.
//
// Octet 8 / Bit 3 shall contain the most significant bit of the PCP value.
//
// Octet 8 / Bit 8 shall be the most significant bit of the S-VID value and
// Octet 9 / Bit 1 shall be the least significant bit  (see clause 7.1).
//
// NOTE: The encoding of the S-Tag in PFCP differs from the encoding of the
// S-Tag defined in IEEE 802.1Q [30].
type STAG1 struct {
	EnterpriseId uint16 `json:"enterpriseId,omitempty"`
	// octet 5
	Pcp bool `json:"pcp,omitempty"`
	Dei bool `json:"dei,omitempty"`
	Vid bool `json:"vid,omitempty"`
	// octet 6 to 7
	PcpValue  uint8  `json:"pcpValue,omitempty"`
	DeiFlag   uint8  `json:"deiFlag,omitempty"`
	SvidValue uint16 `json:"svidValue,omitempty"`
}

func (s *STAG1) MarshalBinary() (data []byte, err error) {
	idx := 0
	data = make([]byte, 2)
	enterpriseId := uint16(VerizonSpecificUEPDNInfo_ENTERPRISEID)
	binary.BigEndian.PutUint16(data, enterpriseId)
	idx += 2

	// Octet 7
	octet7 := btou(s.Pcp) |
		btou(s.Dei)<<1 |
		btou(s.Vid)<<2
	data = append(data, byte(octet7))
	idx = idx + 1

	// Octet 8
	octet8 := (s.PcpValue & 0x07) |
		((s.DeiFlag & 0x01) << 3) |
		(uint8((s.SvidValue & 0x0F00) >> 4))
	data = append(data, byte(octet8))
	idx = idx + 1

	// Octet 9
	data = append(data, byte(uint8(s.SvidValue&0xFF)))
	idx = idx + 1

	return data, nil
}

func (s *STAG1) UnmarshalBinary(data []byte) error {
	var idx = 0
	length := len(data)
	if length < idx+1 {
		return fmt.Errorf("Inadequate length of UserLocationInfo, got zero length")
	}

	s.EnterpriseId = binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2

	// Octet 7
	if length < idx+1 {
		return fmt.Errorf("ie STAG octet5 Inadequate TLV length: %d", length)
	}
	s.Pcp = utob(uint8(data[idx]) & BitMask1)
	s.Dei = utob(uint8(data[idx]) & BitMask2)
	s.Vid = utob(uint8(data[idx]) & BitMask3)
	idx = idx + 1

	// Octet 8
	if length < idx+1 {
		return fmt.Errorf("ie STAG.PcpFlag/DeiFlag/SvidValue Inadequate TLV length: %d", length)
	}
	s.PcpValue = uint8(data[idx] & 0x07)
	s.DeiFlag = uint8((data[idx] >> 3) & 0x01)
	svid_high := uint8((data[idx] >> 4) & 0x0F)
	idx = idx + 1

	// Octet 9
	if length < idx+1 {
		return fmt.Errorf("ie STAG.SvidValue Inadequate TLV length: %d", length)
	}
	s.SvidValue = (uint16(svid_high) << 8) | (uint16(data[idx]) & 0xff)
	return nil
}

// STAG2 - Verizon Custom IE based on TS29.244 8.2.95 S-TAG (Service-VLAN tag)
//
// The STAG2 IE type shall be encoded as shown in below figure. It shall
// contain Service-VLAN tag (S-TAG) as  defined in IEEE 802.1Q [30]
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 32793(decimal)                        |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	5 to 6   |                			Enterprise ID           		      |
//	         |----------------------------------------------------------------|
//	  7      |             Spare                     |  VID  |  DEI  |  PCP   |
//	         |----------------------------------------------------------------|
//	  8      |        S-VID Value            | DEI   |       PCP Value        |
//	         |                               | Flag  |                        |
//	         |----------------------------------------------------------------|
//	  9      |                           S-VID Value                          |
//	         |----------------------------------------------------------------|
//
// 10 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 7:
//   - Bit 1 – PCP: If this bit is set to "1", then PCP Value field shall
//     used by the receiver, otherwise the PCP Value  field shall be ignored.
//   - Bit 2 – DEI: If this bit is set to "1", then DEI flag field shall
//     used by the receiver, otherwise the DEI flag field  shall be ignored.
//   - Bit 3 – VID: If this bit is set to "1", then S-VID value field shall
//     used by the receiver, otherwise the VID Value  field shall be ignored.
//   - Bit 4 to 8 – spare and reserved for future use.
//
// The PCP value, DEI flag and S-VID Value are specified in IEEE 802.1Q [30]
// tag format.
//
// Octet 8 / Bit 3 shall contain the most significant bit of the PCP value.
//
// Octet 8 / Bit 8 shall be the most significant bit of the S-VID value and
// Octet 9 / Bit 1 shall be the least significant bit  (see clause 7.1).
//
// NOTE: The encoding of the S-Tag in PFCP differs from the encoding of the
// S-Tag defined in IEEE 802.1Q [30].
type STAG2 struct {
	EnterpriseId uint16 `json:"enterpriseId,omitempty"`
	// octet 5
	Pcp bool `json:"pcp,omitempty"`
	Dei bool `json:"dei,omitempty"`
	Vid bool `json:"vid,omitempty"`
	// octet 6 to 7
	PcpValue  uint8  `json:"pcpValue,omitempty"`
	DeiFlag   uint8  `json:"deiFlag,omitempty"`
	SvidValue uint16 `json:"svidValue,omitempty"`
}

func (s *STAG2) MarshalBinary() (data []byte, err error) {
	idx := 0
	data = make([]byte, 2)
	// Octet 5 to 6
	enterpriseId := uint16(VerizonSpecificUEPDNInfo_ENTERPRISEID)
	binary.BigEndian.PutUint16(data, enterpriseId)
	idx += 2

	// Octet 7
	octet7 := btou(s.Pcp) |
		btou(s.Dei)<<1 |
		btou(s.Vid)<<2
	data = append(data, byte(octet7))
	idx = idx + 1

	// Octet 8
	octet8 := (s.PcpValue & 0x07) |
		((s.DeiFlag & 0x01) << 3) |
		(uint8((s.SvidValue & 0x0F00) >> 4))
	data = append(data, byte(octet8))
	idx = idx + 1

	// Octet 9
	data = append(data, byte(uint8(s.SvidValue&0xFF)))
	idx = idx + 1

	return data, nil
}

func (s *STAG2) UnmarshalBinary(data []byte) error {
	var idx = 0
	length := len(data)
	if length < idx+1 {
		return fmt.Errorf("Inadequate length of UserLocationInfo, got zero length")
	}
	// Octet 5 to 6
	s.EnterpriseId = binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2

	// Octet 7
	if length < idx+1 {
		return fmt.Errorf("ie STAG octet5 Inadequate TLV length: %d", length)
	}
	s.Pcp = utob(uint8(data[idx]) & BitMask1)
	s.Dei = utob(uint8(data[idx]) & BitMask2)
	s.Vid = utob(uint8(data[idx]) & BitMask3)
	idx = idx + 1

	// Octet 8
	if length < idx+1 {
		return fmt.Errorf("ie STAG.PcpFlag/DeiFlag/SvidValue Inadequate TLV length: %d", length)
	}
	s.PcpValue = uint8(data[idx] & 0x07)
	s.DeiFlag = uint8((data[idx] >> 3) & 0x01)
	svid_high := uint8((data[idx] >> 4) & 0x0F)
	idx = idx + 1

	// Octet 9
	if length < idx+1 {
		return fmt.Errorf("ie STAG.SvidValue Inadequate TLV length: %d", length)
	}
	s.SvidValue = (uint16(svid_high) << 8) | (uint16(data[idx]) & 0xff)
	return nil
}

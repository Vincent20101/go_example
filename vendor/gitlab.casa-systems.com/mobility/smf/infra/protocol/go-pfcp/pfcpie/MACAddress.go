package pfcpie

import (
	"fmt"
)

// TS29.244 8.2.93 MAC address
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 133(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      |               Spare           | UDES  | USOU  | DEST  | SOUR   |
//	         |----------------------------------------------------------------|
//
// m to m+5  |                  Source MAC address value                      |
//
//	|----------------------------------------------------------------|
//
// n to n+5  |                 Destination MAC address value                  |
//
//	|----------------------------------------------------------------|
//
// o to o+5  |                 Upper Source MAC address value                 |
//
//	|----------------------------------------------------------------|
//
// p to p+5  |               Upper Destination MAC address value              |
//
//	|----------------------------------------------------------------|
//
// s to n+4  |   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//
//   - Bit 1 – SOUR (Source): If this bit is set to "1", then the source MAC
//     address value is provided.
//
//   - Bit 2 – DEST (Destination): If this bit is set to "1", then the
//     destination MAC address value is provided.
//
//   - Bit 3 – USOU (Source): If this bit is set to "1", then the source MAC
//     address value contains the lower value and  Upper Source MAC address value
//     contains the upper value of an MAC address range.
//
//   - Bit 4 – UDES (Destination): If this bit is set to "1", then the
//     destination MAC address value contains the lower  value and Upper
//     Destination MAC address value contains the upper value of an MAC address
//     range.
//
// - Bit 5 to 8: Spare, for future use and set to "0".
//
// Octets "m to (m+5)" or "n to (n+5)" and "o to (o+5)" or "p to (p+5)", if
// present, shall contain a MAC address value (12-digit hexadecimal numbers).
type MACAddress struct {
	Sour                       bool   `json:"sour,omitempty"`
	Desc                       bool   `json:"desc,omitempty"`
	Usou                       bool   `json:"usou,omitempty"`
	Udes                       bool   `json:"udes,omitempty"`
	SourceMacAddress           []byte `json:"sourceMacAddress,omitempty"`
	DestinationMacAddress      []byte `json:"destinationMacAddress,omitempty"`
	UpperSourceMacAddress      []byte `json:"upperSourceMacAddress,omitempty"`
	UpperDestinationMacAddress []byte `json:"upperDestinationMacAddress,omitempty"`
}

func (m *MACAddress) MarshalBinary() (data []byte, err error) {
	// Octet 5
	octet5 := btou(m.Sour) |
		btou(m.Desc)<<1 |
		btou(m.Usou)<<2 |
		btou(m.Udes)<<3
	data = append([]byte(""), byte(octet5))

	if m.Sour == true {
		src := make([]byte, 6)
		copy(src, m.SourceMacAddress)
		data = append(data, src...)
	}

	if m.Desc == true {
		des := make([]byte, 6)
		copy(des, m.DestinationMacAddress)
		data = append(data, des...)
	}

	if m.Usou == true {
		usrc := make([]byte, 6)
		copy(usrc, m.UpperSourceMacAddress)
		data = append(data, usrc...)
	}

	if m.Udes == true {
		udes := make([]byte, 6)
		copy(udes, m.UpperDestinationMacAddress)
		data = append(data, udes...)
	}

	return data, nil
}

func (m *MACAddress) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	var idx uint16 = 0

	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE MACAddress Inadequate TLV length: %d", length)
	}
	m.Sour = utob(uint8(data[idx]) & BitMask1)
	m.Desc = utob(uint8(data[idx]) & BitMask2)
	m.Usou = utob(uint8(data[idx]) & BitMask3)
	m.Udes = utob(uint8(data[idx]) & BitMask4)
	idx = idx + 1

	// Source Mac address value.
	if m.Sour == true {
		if length < idx+6 {
			return fmt.Errorf("IE MACAddress Field Sour Inadequate TLV length: %d", length)
		}
		m.SourceMacAddress = make([]byte, 6)
		copy(m.SourceMacAddress, data[idx:idx+6])
		idx = idx + 6
	}

	// Destination Mac address value.
	if m.Desc == true {
		if length < idx+6 {
			return fmt.Errorf("IE MACAddress Field Desc Inadequate TLV length: %d", length)
		}
		m.DestinationMacAddress = make([]byte, 6)
		copy(m.DestinationMacAddress, data[idx:idx+6])
		idx = idx + 6
	}

	// Source Mac address value.
	if m.Usou == true {
		if length < idx+6 {
			return fmt.Errorf("IE MACAddress Field Usou Inadequate TLV length: %d", length)
		}
		m.UpperSourceMacAddress = make([]byte, 6)
		copy(m.UpperSourceMacAddress, data[idx:idx+6])
		idx = idx + 6
	}

	// Source Mac address value.
	if m.Udes == true {
		if length < idx+6 {
			return fmt.Errorf("IE MACAddress Field Udes Inadequate TLV length: %d", length)
		}
		m.UpperDestinationMacAddress = make([]byte, 6)
		copy(m.UpperDestinationMacAddress, data[idx:idx+6])
		idx = idx + 6
	}

	return nil
}

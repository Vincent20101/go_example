package pfcpie

import (
	"encoding/binary"
	"fmt"
	"net"
)

// TS29244 8.2.37 F-SEID
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                     Type = 57 (decimal)                        |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      | Spare | Spare | Spare | Spare | Spare | Spare |   V4  |   V6   |
//	         |----------------------------------------------------------------|
//	6 to 13  |                          SEID                                  |
//	         |----------------------------------------------------------------|
//
// m to m+3  |                       IPv4 Address                             |
//
//	|----------------------------------------------------------------|
//
// p to p+15 |                       IPv6 Address                             |
//
//	|----------------------------------------------------------------|
//
// k to n+4  |   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//   - Bit 1 – V6: If this bit is set to "1", then IPv6 address field shall be
//     present in the F-SEID, otherwise the IPv6  address field is not present
//     at all.
//   - Bit 2 – V4: If this bit is set to "1", then IPv4 address field shall be
//     present in the F-SEID, otherwise the IPv4  address field is not present
//     at all.
//   - Bit 3 to 8 are spare and reserved for future use.
//
// At least one of V4 and V6 shall be set to "1", and both may be set to "1".
//
// Octets "m to (m+3)" and/or "p to (p+15)" (IPv4 address / IPv6 address
// fields), if present, contain respective address
type FSEID struct {
	V4          bool   `json:"v4,omitempty"`
	V6          bool   `json:"v6,omitempty"`
	Seid        uint64 `json:"seid,omitempty"`
	Ipv4Address net.IP `json:"ipv4Address,omitempty"`
	Ipv6Address net.IP `json:"ipv6Address,omitempty"`
}

func (f *FSEID) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5
	tmpUint8 := btou(f.V4)<<1 | btou(f.V6)
	data = append([]byte(""), byte(tmpUint8))
	idx = idx + 1

	// Octet 6 to 13
	data = append(data, make([]byte, 8)...)
	binary.BigEndian.PutUint64(data[idx:], f.Seid)

	// Octet m to (m+3)
	if f.V4 {
		if len(f.Ipv4Address) == 0 || f.Ipv4Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("Marshal FSEID failed: IPv4 address shall be present if V4 is set")
		}
		data = append(data, f.Ipv4Address.To4()...)
	}

	// Octet p to (p+15)
	if f.V6 {
		if len(f.Ipv6Address) == 0 || f.Ipv6Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("Marshal FSEID failed: IPv6 address shall be present if V6 is set")
		}
		data = append(data, f.Ipv6Address.To16()...)
	}

	if !f.V4 && !f.V6 {
		return []byte(""), fmt.Errorf("Marshal FSEID failed: At least one of V4 and V6 flags shall be set")
	}

	return data, nil
}

func (f *FSEID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie FSEID octet5 Inadequate TLV length: %d", length)
	}
	f.V4 = utob(uint8(data[idx]) & BitMask2)
	f.V6 = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet 6 to 13: SEID
	if length < idx+8 {
		return fmt.Errorf("ie FSEID.Seid Inadequate TLV length: %d", length)
	}
	f.Seid = binary.BigEndian.Uint64(data[idx:])
	idx = idx + 8

	// Octet m to (m+3)
	if f.V4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("ie FSEID.Ipv4Address Inadequate TLV length: %d", length)
		}
		f.Ipv4Address = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	}

	// Octet p to (p+15)
	if f.V6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("ie FSEID.Ipv6Address Inadequate TLV length: %d", length)
		}
		f.Ipv6Address = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	if !f.V4 && !f.V6 {
		return fmt.Errorf("None of V4 and V6 flags is set")
	}

	return nil
}

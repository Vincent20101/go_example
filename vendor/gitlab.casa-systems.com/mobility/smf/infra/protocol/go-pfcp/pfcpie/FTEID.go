package pfcpie

import (
	"encoding/binary"
	"fmt"
	"net"
)

// TS29244 8.2.3 F-TEID
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                     Type = 21 (decimal)                        |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      |             Spare             | CHID  |   CH  |   V6  |   V4   |
//	         |----------------------------------------------------------------|
//	6 to 9   |                          TEID                                  |
//	         |----------------------------------------------------------------|
//
// m to m+3  |                       IPv4 Address                             |
//
//	|----------------------------------------------------------------|
//
// p to p+15 |                       IPv6 Address                             |
//
//	      |----------------------------------------------------------------|
//	q     |                         CHOOSE ID                              |
//	      |----------------------------------------------------------------|
//
// k to n+4  |   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//
//   - Bit 1 – V4: If this bit is set to "1" and the CH bit is not set, then
//     the IPv4 address field shall be present, otherwise the IPv4 address
//     field shall not be present.
//
//   - Bit 2 – V6: If this bit is set to "1" and the CH bit is not set, then
//     the IPv6 address field shall be present, otherwise the IPv6 address
//     field shall not be present.
//
//   - Bit 3 – CH (CHOOSE): If this bit is set to "1", then the TEID, IPv4
//     address and IPv6 address fields shall not be  present and the UP
//     function shall assign an F-TEID with an IP4 or an IPv6 address if the V4
//     or V6 bit is set  respectively. This bit shall only be set by the CP
//     function.
//
//   - Bit 4 – CHID (CHOOSE ID): If this bit is set to "1", then the UP
//     function shall assign the same F-TEID to the  PDRs requested to be
//     created in a PFCP Session Establishment Request or PFCP Session
//     Modification Request  with the same CHOOSE ID value. This bit may only
//     be set to "1" if the CH bit it set to "1". This bit shall only be set by
//     the CP function.
//
//   - Bit 5 to 8: Spare, for future use and set to "0".
//
// At least one of the V4 and V6 flags shall be set to "1", and both may be set
// to "1" for both scenarios:
//
//   - when the CP function is providing F-TEID, i.e. both IPv4 address field
//     and IPv6 address field may be present; or
//
//   - when the UP function is requested to allocate the F-TEID, i.e. when
//     CHOOSE bit is set to "1", and the IPv4  address and IPv6 address fields
//     are not present.
//
// Octet 6 to 9 (TEID) shall be present and shall contain a GTP-U TEID, if the
// CH bit in octet 5 is not set. When the TEID is present, if both IPv4 and
// IPv6 addresses are present in the F-TEID IE, then the TEID value shall be
// shared by both  addresses.
//
// Octets "m to (m+3)" and/or "p to (p+15)" (IPv4 address / IPv6 address
// fields), if present, it shall contain the respective
//
// Octet q shall be present and shall contain a binary integer value if the
// CHID bit in octet 5 is set to "1".
type FTEID struct {
	Chid        bool   `json:"chid,omitempty"`
	Ch          bool   `json:"ch,omitempty"`
	V6          bool   `json:"v6,omitempty"`
	V4          bool   `json:"v4,omitempty"`
	Teid        uint32 `json:"teid,omitempty"`
	Ipv4Address net.IP `json:"ipv4Address,omitempty"`
	Ipv6Address net.IP `json:"ipv6Address,omitempty"`
	ChooseId    uint8  `json:"chooseId,omitempty"`
}

func (f *FTEID) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5
	tmpUint8 := btou(f.Chid)<<3 |
		btou(f.Ch)<<2 |
		btou(f.V6)<<1 |
		btou(f.V4)
	data = append([]byte(""), byte(tmpUint8))
	idx = idx + 1

	if !f.Ch {
		// Octet 6 to 9
		data = append(data, make([]byte, 4)...)
		binary.BigEndian.PutUint32(data[idx:], f.Teid)

		// Octet m to (m+3)
		if f.V4 {
			if len(f.Ipv4Address) == 0 || f.Ipv4Address.IsUnspecified() {
				return []byte(""), fmt.Errorf("Marshal TEID failed: IPv4 address shall be present if V4 is set")
			}
			data = append(data, f.Ipv4Address.To4()...)
		}

		// Octet p to (p+15)
		if f.V6 {
			if len(f.Ipv6Address) == 0 || f.Ipv6Address.IsUnspecified() {
				return []byte(""), fmt.Errorf("Marshal TEID failed: IPv6 address shall be present if V6 is set")
			}
			data = append(data, f.Ipv6Address.To16()...)
		}

		if !f.V4 && !f.V6 {
			return []byte(""), fmt.Errorf("Marshal TEID failed:At least one of V4 and V6 flags shall be set")
		}
	} else {
		// Octet q
		if f.Chid {
			data = append(data, byte(f.ChooseId))
		}
	}

	return data, nil
}

func (f *FTEID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie FTEID octet5 Inadequate TLV length: %d", length)
	}
	f.Chid = utob(uint8(data[idx]) & BitMask4)
	f.Ch = utob(uint8(data[idx]) & BitMask3)
	f.V6 = utob(uint8(data[idx]) & BitMask2)
	f.V4 = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	if !f.Ch {
		// Octet 6 to 9
		if length < idx+4 {
			return fmt.Errorf("ie FTEID.Teid Inadequate TLV length: %d", length)
		}
		f.Teid = binary.BigEndian.Uint32(data[idx:])
		idx = idx + 4

		// Octet m to (m+3)
		if f.V4 {
			if length < idx+net.IPv4len {
				return fmt.Errorf("ie FTEID.Ipv4Address Inadequate TLV length: %d", length)
			}
			f.Ipv4Address = net.IP(data[idx : idx+net.IPv4len]).To16()
			idx = idx + net.IPv4len
		}

		// Octet p to (p+15)
		if f.V6 {
			if length < idx+net.IPv6len {
				return fmt.Errorf("ie FTEID.Ipv6Address Inadequate TLV length: %d", length)
			}
			f.Ipv6Address = net.IP(data[idx : idx+net.IPv6len]).To16()
			idx = idx + net.IPv6len
		}

		if !f.V4 && !f.V6 {
			return fmt.Errorf("None of V4 and V6 flags is set")
		}
	} else {
		// Octet q
		if f.Chid {
			if length < idx+1 {
				return fmt.Errorf("ie FTEID.ChooseId Inadequate TLV length: %d", length)
			}
			f.ChooseId = uint8(data[idx])
			idx = idx + 1
		}
	}

	return nil
}

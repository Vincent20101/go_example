package pfcpie

import (
	"fmt"
	"net"
)

// UE IP Adress IE. Refert to [TS29244 8.2.62 UE IP Address].
//
// The UE IP Address IE type shall be encoded as shown in Figure 8.2.62-1. It
// contains a source or destination IP address.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                       Type = 93(decimal)                      |
//	         |---------------------------------------------------------------|
//	3 to 4   |                           Length = n                          |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//	   5     | Spare | IP6PL | CHV6  | CHV4  | IPv6D |  S/D  |  v4   |  V6   |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//	m to m+3 |                         IPv4 address                          |
//	         |---------------------------------------------------------------|
//
// p to p+15 |                         IPv6 address                          |
//
//	         |---------------------------------------------------------------|
//	   r     |                   IPv6 Prefix Delegation Bits                 |
//	         |---------------------------------------------------------------|
//	   s     |                      IPv6 Prefix Length                       |
//	         |---------------------------------------------------------------|
//	k to n+4 |   These octets(s) is/are present only if explicity specified  |
//	         +---------------------------------------------------------------+
//	                           Figure 8.2.62-1: UE IP Address
//
// The following flags are coded within Octet 5:
//
//   - Bit 1  V6: If this bit is set to "1", then the CHV6 bit shall not be
//     set and the IPv6 address field shall be present in the UE IP Address,
//     otherwise the IPv6 address field shall not be present.
//
//   - Bit 2  V4: If this bit is set to "1", then the CHV4 bit sshall not be
//     set and the IPv4 address field shall be present in the UE IP Address,
//     otherwise the IPv4 address field shall not be present.
//
//   - Bit 3  S/D: This bit is only applicable to the UE IP Address IE in the
//     PDI IE. It shall be set to "0" and ignored by the receiver in IEs other
//     than PDI IE. In the PDI IE, if this bit is set to "0", this indicates a
//     Source IP address; if this bit is set to "1", this indicates a
//     Destination IP address.
//
//   - Bit 4  IPv6D: This bit is only applicable to the UE IP address IE in
//     the PDI IE and when the V6 bit or CHV6 bit is set to "1". If this bit is
//     set to "1", then the IPv6 Prefix Delegation Bits field shall be present,
//     otherwise the UP function shall consider IPv6 prefix is default /64.
//
//   - Bit 5  CHV4 (CHOOSE IPV4): If this bit is set to "1", then the V4 bit
//     shall not be set, the IPv4 address shall not be present and the UP
//     function shall assign an IPv4 address. This bit shall only be set by the
//     CP function.
//
//   - Bit 6  CHV6 (CHOOSE IPV6): If this bit is set to "1", then the V6 bit
//     shall not be set, the IPv6 address shall not be present and the UP
//     function shall assign an IPv6 address. This bit shall only be set by the
//     CP function.
//
//   - Bit 7  IP6PL (IPv6 Prefix Length): this bit is only applicable when the
//     V6 bit or CHV6 bit is set to "1" and the "IPv6D" bit is set to "0", for
//     an IPv6 prefix other than default /64. If this bit is set to "1", then
//     the IPv6 Prefix Length field shall be present.
//
//   - Bit 8 Spare, for future use and set to "0".
//
// Octets "m to (m+3)" or "p to (p+15)" (IPv4 address / IPv6 address fields),
// if present, shall contain the address value.
//
// Octet r, if present, shall contain the number of bits allocated for IPv6
// prefix delegation (relative to the default /64 IPv6 prefix), e.g. if /60
// IPv6 prefix is used, the value shall be set to "4". When using UE IP address
// IE in a PDI to match packets, the UP function shall only use the IPv6 prefix
// part and ignore the interface identifier part.
//
// The IPv6 Prefix Length in octet s, when present, shall be encoded as a 8
// bits binary integer, e.g. if /72 prefix is used, the value shall be set to
// to (decimal) 72. The prefix length value "128" indicates an individual /128
// IPv6 address.
type UEIPAddress struct {
	Ip6pl                    bool   `json:"ip6pl,omitempty"` // octet5/bit7
	Chv6                     bool   `json:"chv6,omitempty"`  // octet5/bit6
	Chv4                     bool   `json:"chv4,omitempty"`  // octet5/bit5
	Ipv6d                    bool   `json:"ipv6d,omitempty"` // octet5/bit4
	Sd                       bool   `json:"sd,omitempty"`    // octet5/bit3
	V4                       bool   `json:"v4,omitempty"`    // octet5/bit2
	V6                       bool   `json:"v6,omitempty"`    // octet5/bit1
	Ipv4Address              net.IP `json:"ipv4Address,omitempty" cmp:"ignore"`
	Ipv6Address              net.IP `json:"ipv6Address,omitempty" cmp:"ignore"`
	Ipv6PrefixDelegationBits uint8  `json:"ipv6PrefixDelegationBits,omitempty"`
	Ipv6PrefixLength         uint8  `json:"ipv6PrefixLength,omitempty"`
}

func (u *UEIPAddress) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(u.Ip6pl)<<6 |
		btou(u.Chv6)<<5 |
		btou(u.Chv4)<<4 |
		btou(u.Ipv6d)<<3 |
		btou(u.Sd)<<2 |
		btou(u.V4)<<1 |
		btou(u.V6)
	data = append([]byte(""), byte(tmpUint8))

	// Octet m to (m+3)
	if u.V4 {
		if u.Chv4 {
			msg := "If V4 bit is set to 1, then the CHV4 bit shall not be set "
			return []byte(""), fmt.Errorf(msg)
		}
		if len(u.Ipv4Address) == 0 || u.Ipv4Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("Encode UEIPAddress failed: IPv4 address shall be present if V4 is set")
		}
		ipv4 := u.Ipv4Address.To4()
		if ipv4 == nil {
			return []byte(""), fmt.Errorf("Encode UEIPAddress failed: invalid ipv4: %v", u.Ipv4Address)
		}
		data = append(data, ipv4...)
	}

	// Octet p to (p+15)
	if u.V6 {
		if u.Chv6 {
			msg := "If V6 bit is set to 1, then the CHV6 bit shall not be set"
			return []byte(""), fmt.Errorf(msg)
		}
		if len(u.Ipv6Address) == 0 || u.Ipv6Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("Encode UEIPAddress failed: IPv4 address shall be present if V4 is set")
		}
		ipv6 := u.Ipv6Address.To16()
		if ipv6 == nil {
			return []byte(""), fmt.Errorf("Encode UEIPAddress failed: invalid ipv4: %v", u.Ipv4Address)
		}
		data = append(data, ipv6...)
	}

	if u.Ipv6d {
		if !(u.V6 || u.Chv6) {
			msg := "Encode UEIPAddress failed: IPv6D is only applicable when the V6 " +
				"bit or CHV6 bit is set to 1"
			return []byte(""), fmt.Errorf(msg)
		}
		data = append(data, byte(u.Ipv6PrefixDelegationBits))
	}

	if u.Ip6pl {
		if !((u.V6 || u.Chv6) && !u.Ipv6d) {
			msg := "Encode UEIPAddress failed: IP6PL is only applicable when the V6 " +
				"bit or CHV6 bit is set to 1 and the IPv6D bit is set to 0 "
			return []byte(""), fmt.Errorf(msg)
		}
		data = append(data, byte(u.Ipv6PrefixDelegationBits))
	}
	return data, nil
}

func (u *UEIPAddress) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Decode UEIPAddress failed: empty bytes.")
	}
	u.Ip6pl = utob(uint8(data[idx]) & BitMask7)
	u.Chv6 = utob(uint8(data[idx]) & BitMask6)
	u.Chv4 = utob(uint8(data[idx]) & BitMask5)
	u.Ipv6d = utob(uint8(data[idx]) & BitMask4)
	u.Sd = utob(uint8(data[idx]) & BitMask3)
	u.V4 = utob(uint8(data[idx]) & BitMask2)
	u.V6 = utob(uint8(data[idx]) & BitMask1)

	idx = idx + 1

	// Octet m to (m+3)
	if u.V4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("Decode UEIPAddress failed: v4 is true but inadequte Ipv4Addr length, bytes:  %v", data)
		}
		u.Ipv4Address = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	}

	// Octet p to (p+15)
	if u.V6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("Decode UEIPAddress failed: Inadequate TLV length for Ipv6Address, bytes: %v", data)
		}
		u.Ipv6Address = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	// Octet r
	if u.Ipv6d {
		if length < idx+1 {
			return fmt.Errorf("Decode UEIPAddress failed: Inadequate length for Ipv6PrefixDelegationBits, bytes: %v", data)
		}
		u.Ipv6PrefixDelegationBits = uint8(data[idx])
		idx = idx + 1
	}

	if u.Ip6pl {
		if length < idx+1 {
			return fmt.Errorf("Decode UEIPAddress failed: Inadequate length for Ipv6PrefixLength, bytes: %v", data)
		}
		u.Ipv6PrefixLength = uint8(data[idx])
	}

	return nil
}

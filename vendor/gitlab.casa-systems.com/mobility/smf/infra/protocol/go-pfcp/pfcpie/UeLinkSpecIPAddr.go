package pfcpie

import (
	"fmt"
	"net"
)

// TS29.244 8.2.158	UE Link-Specific IP Address
// The UE Link-Specific IP Address IE shall carry link-specific IP address used for MPTCP steering function.
// It shall be encoded as shown in Figure 8.2.158-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	          +----------------------------------------------------------------+
//	1 to 2    |                   Type = 229(decimal)                          |
//	          |----------------------------------------------------------------|
//	3 to 4    |                      Length = n                                |
//	          |----------------------------------------------------------------|
//	  5       |               Spare           |   NV6  |  NV4 |  V6   |   V4   |
//	          |----------------------------------------------------------------|
//
// p to (p+3) |    UE Link-Specific IPv4 Address for 3GPP Access               |
//
//	|----------------------------------------------------------------|
//
// q to (q+15)|    UE Link-Specific IPv6 Address for 3GPP Access               |
//
//	|----------------------------------------------------------------|
//
// r to (r+3) |    UE Link-Specific IPv4 Address for Non-3GPP Access           |
//
//	|----------------------------------------------------------------|
//
// s to (s+15)|    UE Link-Specific IPv6 Adress for Non-3GPP Access            |
//
//	|----------------------------------------------------------------|
//
// m to n+4   |   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//   - Bit 1 – V4: If this bit is set to "1", then the UE Link-Specific IPv4 Address for 3GPP Access shall be
//     present in the UE Link-Specific IP Address IE.
//   - Bit 2 – V6: If this bit is set to "1", then the UE Link-Specific IPv6 Address for 3GPP Access shall be
//     present in the UE Link-Specific IP Address IE.
//   - Bit 3 – NV4: If this bit is set to "1", then the UE Link-Specific IPv4 Address for Non-3GPP Access shall
//     be present in the UE Link-Specific IP Address IE.
//   - Bit 4 – NV6: If this bit is set to "1", then the UE Link-Specific IPv6 Address for Non-3GPP Access shall
//     be present in the UE Link-Specific IP Address IE.
//   - Bit 5 to 8 Spare, for future use and set to "0".
//
// Octets "p to (p+3)" or "q to (q+15)" (IPv4 address / IPv6 address fields),
// if present, shall contain the value for UE Link-Specific IP Address for 3GPP access.
//
// Octets "r to (r+3)" or "s to (s+15)" (IPv4 address / IPv6 address fields),
// if present, shall contain the value for UE Link-Specific IP Address for Non-3GPP access.
type UeLinkSpecIPAddr struct {
	V4             bool   `json:"v4,omitempty"`
	V6             bool   `json:"v6,omitempty"`
	NV4            bool   `json:"nV4,omitempty"`
	NV6            bool   `json:"nV6,omitempty"`
	Ipv4For3GPP    net.IP `json:"ipv4For3GPP,omitempty"`
	Ipv6For3GPP    net.IP `json:"ipv6For3GPP,omitempty"`
	IpV4ForNon3GPP net.IP `json:"ipV4ForNon3GPP,omitempty"`
	Ipv6ForNon3GPP net.IP `json:"ipv6ForNon3GPP,omitempty"`
}

func (u *UeLinkSpecIPAddr) MarshalBinary() (data []byte, err error) {
	errMsgPrefix := "Marshal UeLinkSpecIPAddr failed: %s"

	octet5 := btou(u.V4) |
		btou(u.V6)<<1 |
		btou(u.NV4)<<2 |
		btou(u.NV6)<<3
	data = append([]byte(""), byte(octet5))

	if u.V4 {
		if len(u.Ipv4For3GPP) == 0 || u.Ipv4For3GPP.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix, "IPv4 address shall be present if V4 is set")
		}

		data = append(data, u.Ipv4For3GPP.To4()...)
	}

	if u.V6 {
		if len(u.Ipv6For3GPP) == 0 || u.Ipv6For3GPP.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix, "IPv6 address shall be present if V6 is set")
		}

		data = append(data, u.Ipv6For3GPP.To16()...)
	}

	if u.NV4 {
		if len(u.IpV4ForNon3GPP) == 0 || u.IpV4ForNon3GPP.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix, "IPv4 address shall be present if V4 is set")
		}

		data = append(data, u.IpV4ForNon3GPP.To4()...)
	}

	if u.NV6 {
		if len(u.Ipv6ForNon3GPP) == 0 || u.Ipv6ForNon3GPP.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix, "IPv6 address shall be present if V6 is set")
		}

		data = append(data, u.Ipv6ForNon3GPP.To16()...)
	}

	return data, nil
}

func (u *UeLinkSpecIPAddr) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	var idx uint16 = 0

	if length < idx+1 {
		return fmt.Errorf("ie UeLinkSpecIPAddr Inadequate TLV length: %d", length)
	}
	u.V4 = utob(uint8(data[idx]) & BitMask1)
	u.V6 = utob(uint8(data[idx]) & BitMask2)
	u.NV4 = utob(uint8(data[idx]) & BitMask3)
	u.NV6 = utob(uint8(data[idx]) & BitMask4)
	idx = idx + 1

	// IPv4 address for 3GPP
	if u.V4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("ie UeLinkSpecIPAddr.Ipv4For3GPP Inadequate TLV length: %d", length)
		}

		u.Ipv4For3GPP = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	}

	// IPv6 address for 3GPP
	if u.V6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("ie UeLinkSpecIPAddr.Ipv6For3GPP Inadequate TLV length: %d", length)
		}

		u.Ipv6For3GPP = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	// IPv4 address for Non-3GPP
	if u.NV4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("ie UeLinkSpecIPAddr.IpV4ForNon3GPP Inadequate TLV length: %d", length)
		}

		u.IpV4ForNon3GPP = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	}

	// IPv6 address for Non-3GPP
	if u.NV6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("ie UeLinkSpecIPAddr.Ipv6ForNon3GPP Inadequate TLV length: %d", length)
		}

		u.Ipv6ForNon3GPP = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	return nil
}

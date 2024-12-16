package pfcpie

import (
	"encoding/binary"
	"fmt"
	"net"
)

// TS29.244 8.2.157	MPTCP Address Information
// The MPTCP Address Information IE shall carry the address information of MPTCP proxy in the UPF.
// It shall be encoded as shown in Figure 8.2.157-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 228(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      |               Spare                           |  V6   |   V4   |
//	         |----------------------------------------------------------------|
//	  6      |                       MPTCP Proxy Type                         |
//	         |----------------------------------------------------------------|
//
// 7 to 8    |                       MPTCP Proxy Port                         |
//
//	|----------------------------------------------------------------|
//
// p to p+3  |                       MPTCP Proxy IPv4 Address                 |
//
//	|----------------------------------------------------------------|
//
// q to q+15 |                       MPTCP Proxy IPv6 Address                 |
//
//	|----------------------------------------------------------------|
//
// m to n+4  |   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//   - Bit 1 – V4: If this bit is set to "1", then the MPTCP Proxy IPv4 Address field
//     shall be present in the MPTCP Address Information IE.
//   - Bit 1 – V6: If this bit is set to "1", then the MPTCP IPv6 Address field
//     shall be present in the MPTCP Address Information IE.
//   - Bit 3 to 8 Spare, for future use and set to "0".
//
// Octets 6 shall indicate the MPTCP Proxy Type, with the value specified in clause 6.1.4 of 3GPP TS 24.193 [59].
// Octets 7 to 8 shall indicate the allocated TCP port number of the MPTCP Proxy.
// Octets "p to (p+3)" or "q to (q+15)" (IPv4 address / IPv6 address fields), if present, shall contain the address value.
type MptcpAddrInfo struct {
	V4            bool   `json:"v4,omitempty"`
	V6            bool   `json:"v6,omitempty"`
	ProxyType     uint8  `json:"proxyType,omitempty"`
	ProxyPort     uint16 `json:"proxyPort,omitempty"`
	ProxyIpv4Addr net.IP `json:"proxyIpv4Addr,omitempty"`
	ProxyIpv6Addr net.IP `json:"proxyIpv6Addr,omitempty"`
}

func (m *MptcpAddrInfo) MarshalBinary() (data []byte, err error) {
	errMsgPrefix := "Marshal MptcpAddrInfo failed: %s"
	// Octet 5
	octet5 := btou(m.V4) | btou(m.V6)<<1
	data = append([]byte(""), byte(octet5))

	// Octet 6
	data = append(data, byte(m.ProxyType))

	// Octets 7 to 8
	octet7 := make([]byte, 2)
	binary.BigEndian.PutUint16(octet7, m.ProxyPort)
	data = append(data, octet7...)

	// Octets "p to (p+3)" or "q to (q+15)"
	if m.V4 {
		if len(m.ProxyIpv4Addr) == 0 || m.ProxyIpv4Addr.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix, "IPv4 address shall be present if V4 is set")
		}

		data = append(data, m.ProxyIpv4Addr.To4()...)
	}

	if m.V6 {
		if len(m.ProxyIpv6Addr) == 0 || m.ProxyIpv6Addr.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix, "IPv6 address shall be present if V6 is set")
		}

		data = append(data, m.ProxyIpv6Addr.To16()...)
	}

	return data, nil
}

func (m *MptcpAddrInfo) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	var idx uint16 = 0

	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE MptcpAddrInfo Inadequate TLV length: %d", length)
	}
	m.V4 = utob(uint8(data[idx]) & BitMask1)
	m.V6 = utob(uint8(data[idx]) & BitMask2)
	idx = idx + 1

	// Octet 6
	if length < idx+1 {
		return fmt.Errorf("IE MptcpAddrInfo Field ProxyType Inadequate TLV length: %d", length)
	}
	m.ProxyType = uint8(data[idx])
	idx = idx + 1

	// Octets 7 to 8
	if length < idx+2 {
		return fmt.Errorf("IE MptcpAddrInfo Field ProxyPort Inadequate TLV length: %d", length)
	}
	m.ProxyPort = binary.BigEndian.Uint16(data[idx : idx+2])
	idx = idx + 2

	// IPv4 address
	if m.V4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("IE MptcpAddrInfo Field ProxyIpv4Addr Inadequate TLV length: %d", length)
		}
		m.ProxyIpv4Addr = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	}

	// IPv6 address
	if m.V6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("IE MptcpAddrInfo Field IE MptcpAddrInfo Field Inadequate TLV length: %d", length)
		}
		m.ProxyIpv6Addr = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	return nil
}

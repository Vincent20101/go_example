package pfcpie

import (
	"encoding/binary"
	"fmt"
	"net"
)

// TS29.244 8.2.158	UE Link-Specific IP Address
//
// The UE Link-Specific IP Address IE shall carry link-specific IP address used
// for MPTCP steering function.
//
// It shall be encoded as shown in Figure 8.2.158-1.
//
//	                                       Bits
//	Octets        8       7       6       5       4       3       2       1
//	          +---------------------------------------------------------------+
//	1 to 2    |                   Type = 229(decimal)                         |
//	          |---------------------------------------------------------------|
//	3 to 4    |                      Length = n                               |
//	          |---------------------------------------------------------------|
//	  5       |               Spare                    |  MAC |  V6   |  V4   |
//	          |---------------------------------------------------------------|
//
// p to (p+3) |                     PMF IPv4 Address                          |
//
//	|---------------------------------------------------------------|
//
// q to (q+15)|                     PMF IPv6 Address                          |
//
//	|---------------------------------------------------------------|
//
// r to (r+1) |                 PMF Port for 3GPP Access                      |
//
//	|---------------------------------------------------------------|
//
// s to (s+1) |                PMF Port for Non-3GPP Access                   |
//
//	|---------------------------------------------------------------|
//
// t to (r+5) |               PMF MAC Address for 3GPP Access                 |
//
//	|---------------------------------------------------------------|
//
// u to (r+5) |              PMF MAC Address for Non-3GPP Access              |
//
//	|---------------------------------------------------------------|
//
// m to n+4   |   These octets(s) is/are present only if explicity specified  |
//
//	+---------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//
//   - Bit 1 – V4: If this bit is set to "1", it indicates the PMF IPv4
//     Address filed shall be presend in the PMF Address Information IE,
//     together with: PMF Port for 3GPP Access, PMF Port for Non-3GPP Access.
//
//   - Bit 2 – V6: If this bit is set to "1", it indicates the PMF IPv6
//     Address filed shall be presend in the PMF Address Information IE,
//     together with: PMF Port for 3GPP Access, PMF Port for Non-3GPP Access.
//
//   - Bit 3 – MAC: If this bit is set to "1", it indicates the PMF MAC
//     Address for 3GPP Access, PMF MAC Address for Non-3GPP Access filed shall
//     be presend in the PMF Address Information IE.
//
//   - Bit 4 to 8 Spare, for future use and set to "0".
//
// Octets "p to (p+3)" or "q to (q+15)" (IPv4 address / IPv6 address fields),
// if present, shall contain the value for PMF IP Address.
//
// Octets "r to (r+1)" shall carry the allocated UDP port number associated
// with the 3GPP access network, for IP PDU session.
//
// Octets "s to (s+1)" shall carry the allocated UDP port number associated
// with the Non-3GPP access network, for IP PDU session.
//
// Octets "t to (t+5)" shall carry the allocated PMF MAC address for 3GPP
// access, for Ethernet PDU session.
//
// Octets "u to (u+5)" shall carry the allocated PMF MAC address for Non-3GPP
// access, for Ethernet PDU session.
type PmfAddrInfo struct {
	V4                bool   `json:"v4,omitempty"`
	V6                bool   `json:"v6,omitempty"`
	Mac               bool   `json:"mac,omitempty"`
	Ipv4Addr          net.IP `json:"ipv4Addr,omitempty"`
	Ipv6Addr          net.IP `json:"ipv6Addr,omitempty"`
	PortFor3GPP       uint16 `json:"portFor3GPP,omitempty"`
	PortForNon3GPP    uint16 `json:"portForNon3GPP,omitempty"`
	MacAddrFor3GPP    []byte `json:"macAddrFor3GPP,omitempty"`
	MacAddrForNon3GPP []byte `json:"macAddrForNon3GPP,omitempty"`
}

func (p *PmfAddrInfo) MarshalBinary() (data []byte, err error) {
	errMsgPrefix := "Marshal PmfAddrInfo failed:"

	octet5 := btou(p.V4) |
		btou(p.V6)<<1 |
		btou(p.Mac)<<2
	data = append([]byte(""), byte(octet5))

	if p.V4 {
		if len(p.Ipv4Addr) == 0 || p.Ipv4Addr.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix + "IPv4 address shall be present if V4 is set")
		}

		data = append(data, p.Ipv4Addr.To4()...)
	}

	if p.V6 {
		if len(p.Ipv6Addr) == 0 || p.Ipv6Addr.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix + "IPv6 address shall be present if V6 is set")
		}

		data = append(data, p.Ipv6Addr.To16()...)
	}

	if p.V4 || p.V6 {
		portFor3GPP := make([]byte, 2)
		binary.BigEndian.PutUint16(portFor3GPP, p.PortFor3GPP)
		data = append(data, portFor3GPP...)

		portForNon3Gpp := make([]byte, 2)
		binary.BigEndian.PutUint16(portForNon3Gpp, p.PortForNon3GPP)
		data = append(data, portForNon3Gpp...)
	}

	if p.Mac {
		macAddrFor3GPP := make([]byte, 6)
		copy(macAddrFor3GPP, p.MacAddrFor3GPP)
		data = append(data, macAddrFor3GPP...)

		macAddrForNon3GPP := make([]byte, 6)
		copy(macAddrForNon3GPP, p.MacAddrForNon3GPP)
		data = append(data, macAddrForNon3GPP...)
	}

	return data, nil
}

func (p *PmfAddrInfo) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	var idx uint16 = 0

	if length < idx+1 {
		return fmt.Errorf("IE PmfAddrInfo Inadequate TLV length: %d", length)
	}
	p.V4 = utob(uint8(data[idx]) & BitMask1)
	p.V6 = utob(uint8(data[idx]) & BitMask2)
	p.Mac = utob(uint8(data[idx]) & BitMask3)
	idx = idx + 1

	// IPv4 address for 3GPP
	if p.V4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("IE PmfAddrInfo Fidld Ipv4Addr Inadequate TLV length: %d", length)
		}
		p.Ipv4Addr = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	}

	// IPv6 address for 3GPP
	if p.V6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("IE PmfAddrInfo Fidld Ipv6Addr Inadequate TLV length: %d", length)
		}
		p.Ipv6Addr = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	if p.V4 || p.V6 {
		if length < idx+2 {
			return fmt.Errorf("IE PmfAddrInfo Field PortForNon3GPP Inadequate TLV length: %d", length)
		}
		p.PortFor3GPP = binary.BigEndian.Uint16(data[idx : idx+2])
		idx = idx + 2

		if length < idx+2 {
			return fmt.Errorf("IE PmfAddrInfo Field PortForNon3GPP Inadequate TLV length: %d", length)
		}
		p.PortForNon3GPP = binary.BigEndian.Uint16(data[idx : idx+2])
		idx = idx + 2
	}

	if p.Mac {
		if length < idx+12 {
			return fmt.Errorf("IE PmfAddrInfo Field MacAddrFor3GPP Inadequate TLV length: %d", length)
		}
		p.MacAddrFor3GPP = make([]byte, 6)
		copy(p.MacAddrFor3GPP, data[idx:idx+6])
		idx = idx + 6

		p.MacAddrForNon3GPP = make([]byte, 6)
		copy(p.MacAddrForNon3GPP, data[idx:idx+6])
		idx = idx + 6
	}

	return nil
}

package pfcpie

import (
	"encoding/binary"
	"fmt"
	"net"
)

/*
	The definition of this IE is consistent with the private IE PFCP_IE_CASA_K8S_FWD_SRC_IP_PORT of PGW-C.
	It is used in the following situations:
	1.
		The sess est/mod/del req received from pgw-c will carry this private IE(filled the UPF ip/port) at the end.
		pfcp-svc should get the upf ip/port from this private IE and remove it to the pfcp msg, then forward to corresponding upf.
		When received the corresponding sess est/mod/del rsp, pfcp-svc should append this private IE(filled the UPF ip/port) to the end of pfcp msg and forward to pgw-c.
	2.
		The node asso setup/update/release received from upfmgr-svc will carry this private IE(filled the UPF ip/port) at the end.
		pfcp-svc should get the upf ip/port from this private IE and remove it to the pfcp msg, then forward forward to corresponding upf.
		When receive the corresponding asso setup/update/release rsp, pfcp-svc should forward the rsp back to the same upfmgr-svc pod.
*/

const PFCP_IE_CASA_UPF_ADDR = 0x8013

type CasaUpfAddr struct {
	CasaId      uint16 `json:"casaId,omitempty"`
	IsIpv4      uint8  `json:"isIpv4,omitempty"`
	Port        uint16 `json:"port,omitempty"`
	Ipv4Address net.IP `json:"ipv4Address,omitempty"`
	Ipv6Address net.IP `json:"ipv6Address,omitempty"`
}

func (u *CasaUpfAddr) MarshalBinary() (data []byte, err error) {
	// CasaId
	casaId := make([]byte, 2)
	binary.BigEndian.PutUint16(casaId, 20858) // casa id should be set 20858
	data = append([]byte(""), casaId...)

	// IsIPv4
	data = append(data, u.IsIpv4)

	// Port
	srcPort := make([]byte, 2)
	binary.BigEndian.PutUint16(srcPort, u.Port)
	data = append(data, srcPort...)

	if u.IsIpv4 == 1 { // Ipv4Address
		if len(u.Ipv4Address) == 0 || u.Ipv4Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("Encode CasaUpfAddr failed: IPv4 address shall be present if IsIpv4 is set")
		}
		ipv4 := u.Ipv4Address.To4()
		if ipv4 == nil {
			return []byte(""), fmt.Errorf("Encode CasaUpfAddr failed: invalid ipv4: %v", u.Ipv4Address)
		}
		data = append(data, ipv4...)
	} else { // Ipv6Address
		if len(u.Ipv6Address) == 0 || u.Ipv6Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("Encode CasaUpfAddr failed: IPv6 address shall be present if IsIpv4 is not set")
		}
		ipv6 := u.Ipv6Address.To16()
		if ipv6 == nil {
			return []byte(""), fmt.Errorf("Encode CasaUpfAddr failed: invalid IPv6: %v", u.Ipv6Address)
		}
		data = append(data, ipv6...)
	}
	return data, nil
}

func (u *CasaUpfAddr) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	var idx uint16 = 0

	// CasaId
	if length < idx+2 {
		return fmt.Errorf("IE CasaUpfAddr Inadequate TLV length: %d", length)
	}
	u.CasaId = binary.BigEndian.Uint16(data[idx : idx+2])

	idx = idx + 2

	// IsIPv4
	if length < idx+1 {
		return fmt.Errorf("IE CasaUpfAddr Inadequate TLV length: %d", length)
	}
	u.IsIpv4 = data[idx]

	idx = idx + 1

	// Port
	if length < idx+2 {
		return fmt.Errorf("IE CasaUpfAddr Inadequate TLV length: %d", length)
	}
	u.Port = binary.BigEndian.Uint16(data[idx : idx+2])

	idx = idx + 2

	if u.IsIpv4 == 1 { // Ipv4Address
		if length < idx+net.IPv4len {
			return fmt.Errorf("IE CasaUpfAddr Inadequate TLV length for Ipv4Address, bytes:  %v", data)
		}
		u.Ipv4Address = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	} else {
		if length < idx+net.IPv6len {
			return fmt.Errorf("IE CasaUpfAddr Inadequate TLV length for Ipv6Address, bytes: %v", data)
		}
		u.Ipv6Address = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	return nil
}

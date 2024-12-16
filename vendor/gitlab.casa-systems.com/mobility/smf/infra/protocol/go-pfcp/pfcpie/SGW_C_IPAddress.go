package pfcpie

import (
	"errors"
	"fmt"
	"net"
)

// SGW-C IP Address
// Casa Specific IE for Verizon use
//
//	Octets        8       7       6       5       4       3       2       1
//	          +-------------------------------------------------------------+
//	1         |                   Type = 4(decimal)                         |
//	          |-------------------------------------------------------------|
//	2         |                      Length = n                             |
//	          |-------------------------------------------------------------|
//	3         |	               Spare                      |   V4   |   V6   |
//	          |-------------------------------------------------------------|
//	4 to 7    |	                  IPV4 Address                              |
//	          |-------------------------------------------------------------|
//	8 to 23   |	                  IPv6 Address                              |
//	          |-------------------------------------------------------------|
type SGW_C_IPAddress struct {
	V4          bool   `json:"v4,omitempty"` // octet3/bit2
	V6          bool   `json:"v6,omitempty"` // octet3/bit1
	Ipv4Address net.IP `json:"ipv4Address,omitempty" cmp:"ignore"`
	Ipv6Address net.IP `json:"ipv6Address,omitempty" cmp:"ignore"`
}

func (p *SGW_C_IPAddress) MarshalBinary() (data []byte, err error) {
	// Octet 5: V6 and V4
	tmpOctet := btou(p.V6) | btou(p.V4)<<1
	data = append([]byte(""), tmpOctet)

	if !p.V4 && !p.V6 {
		return nil, errors.New("At least one of V4 and V6 should be true.")
	}

	// Octet m to m+3: Ipv4 Address
	if p.V4 {
		if len(p.Ipv4Address) == 0 || p.Ipv4Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("IPv4 address shall be present if V4 is set")
		}
		ipv4 := p.Ipv4Address.To4()
		if ipv4 == nil {
			return nil, errors.New("SGW_C_IPAddress.Ipv4Address.To4() failed.")
		}
		data = append(data, ipv4...)
	}

	// octet n to n+15
	if p.V6 {
		if len(p.Ipv6Address) == 0 || p.Ipv6Address.IsUnspecified() {
			return []byte(""), fmt.Errorf("IPv6 address shall be present if V6 is set")
		}
		ipv6 := p.Ipv6Address.To16()
		if ipv6 == nil {
			return nil, errors.New("SGW_C_IPAddress.Ipv4Address.To16() failed.")
		}
		data = append(data, ipv6...)
	}
	return data, nil
}

func (p *SGW_C_IPAddress) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var ind uint16 = 0
	if length < ind+1 {
		return fmt.Errorf("Decode SGW_C_IPAddress failed: empty bytes.")
	}
	p.V6 = utob(data[ind] & BitMask1)
	p.V4 = utob(data[ind] & BitMask2)

	ind += 1
	// Octet m to m+3: Ipv4 Address
	if p.V4 {
		if length < ind+net.IPv4len {
			return fmt.Errorf("Decode SGW_C_IPAddress failed: v4 is true but inadequte Ipv4Address length, bytes: %v", data)
		}
		p.Ipv4Address = net.IP(data[ind : ind+net.IPv4len]).To4()
		ind = ind + net.IPv4len
	}
	// Octet p to (p+15)
	if p.V6 {
		if length < ind+net.IPv6len {
			return fmt.Errorf("Decode SGW_C_IPAddress failed: Inadequate TLV length for Ipv6Address, bytes: %v", data)
		}
		p.Ipv6Address = net.IP(data[ind : ind+net.IPv6len]).To16()
		ind = ind + net.IPv6len
	}
	return nil
}

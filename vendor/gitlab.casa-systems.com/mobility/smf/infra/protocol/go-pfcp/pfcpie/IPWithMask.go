package pfcpie

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	IPv4WithMasklen = 5
	IPv6WithMasklen = 17
)

// IPv4WithMask consits of a IP and a Mask.
// e.g.  182.123.12.42/24
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	   1     |                           Mask                                 |
//	         |----------------------------------------------------------------|
//	 2 to 5  |                        IPv4 Address                            |
//	         +----------------------------------------------------------------+
//	                                IPv4 With Mask
type IPv4WithMask struct {
	IP         net.IP `json:"iP,omitempty"`
	Mask       uint8  `json:"mask,omitempty"`
	parseError error
}

// GetParseError get error when new IPv4WithMask from ip/mask string.
func (i *IPv4WithMask) GetParseError() error {
	return i.parseError
}

func parseIpAndMask(s string) (ip net.IP, mask uint8, err error) {
	// Split IP and Mask
	res := strings.Split(s, "/")
	if len(res) != 2 {
		return nil, 0, fmt.Errorf("IpWithMask must be in the format of ip/mask, but got %s", s)
	}
	ipStr, maskStr := res[0], res[1]

	// Parse IP and Mask
	ip = net.ParseIP(ipStr)
	maskNum, err := strconv.Atoi(maskStr)
	if err != nil {
		return nil, 0, fmt.Errorf("mask %s not valid", maskStr)
	}
	mask = uint8(maskNum)
	return ip, mask, err
}

// NewIPv4WithMask parse the ip/mask string and return a NewIPv4WithMask.
//
// If the string is not valid, return a IPv4WithMask with a nil IP.
// You can use IPv4WithMask.GetParseError() to get the parseError.
//
// e.g.
//
//	ipWithMask := NewIPv4WithMask("1.1.1.1/24")
//	if ipWithMask.GetParseError() != nil {
//	    fmt.Println("Parse ip/mask Error!")
//	}
//
// You can ignore the handling of parse error if you're pretty sure your
// `ip/mask` string is correct.
func NewIPv4WithMask(s string) *IPv4WithMask {
	ipWithMask := &IPv4WithMask{}
	ip, mask, err := parseIpAndMask(s)
	if err != nil {
		ipWithMask.IP = nil
		ipWithMask.parseError = err
		return ipWithMask
	}
	ipWithMask.Mask = mask
	ipWithMask.IP = ip
	return ipWithMask

}

func (i *IPv4WithMask) MarshalBinary() (data []byte, err error) {
	if i.IP == nil {
		return nil, fmt.Errorf("IPv4 is nil")
	}
	data = append(data, i.Mask)
	ipv4 := i.IP.To4()
	if ipv4 == nil {
		return nil, fmt.Errorf("IP.to4() is nil")
	}
	data = append(data, ipv4...)
	return data, err
}

func (i *IPv4WithMask) UnmarshalBinary(data []byte) (err error) {
	byteReader := NewByteReader(data)
	i.Mask, err = byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("inadequate length for mask: %v", err)
	}
	ip, err := byteReader.ReadBytes(net.IPv4len)
	i.IP = net.IP(ip).To16()
	if err != nil {
		return fmt.Errorf("inadequate length for IPv4: %v", err)
	}
	return err
}

// IPv6WithMask consits of a IP and a Mask.
// e.g.  ab12:2211:0:1::11/24
//
//	Ipv6/Mask:
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	   1     |                           Mask                                 |
//	         |----------------------------------------------------------------|
//	2 to 17  |                        IPv6 Address                            |
//	         +----------------------------------------------------------------+
type IPv6WithMask struct {
	IP         net.IP `json:"iP,omitempty"`
	Mask       uint8  `json:"mask,omitempty"`
	parseError error
}

// GetParseError get error when new IPv6WithMask from ip/mask string.
func (i *IPv6WithMask) GetParseError() error {
	return i.parseError
}

func parseIpv6AndMask(s string) (ip net.IP, mask uint8, err error) {
	// Split IP and Mask
	res := strings.Split(s, "/")
	if len(res) != 2 {
		return nil, 0, fmt.Errorf("IpWithMask must be in the format of ip/mask, but got %s", s)
	}
	ipStr, maskStr := res[0], res[1]

	// Parse IP and Mask
	ip = net.ParseIP(ipStr)
	maskNum, err := strconv.Atoi(maskStr)
	if err != nil {
		return nil, 0, fmt.Errorf("mask %s not valid", maskStr)
	}
	mask = uint8(maskNum)
	return ip, mask, err
}

// NewIPv6WithMask parse the ip/mask string and return a NewIPv6WithMask.
//
// If the string is not valid, return a IPv6WithMask with a nil IP.
// You can use IPv6WithMask.GetParseError() to get the parseError.
//
// e.g.
//
//	ipWithMask := NewIPv6WithMask("1010:1010:1010::/24")
//	if ipWithMask.GetParseError() != nil {
//	    fmt.Println("Parse ip/mask Error!")
//	}
//
// You can ignore the handling of parse error if you're pretty sure your
// `ip/mask` string is correct.
func NewIPv6WithMask(s string) *IPv6WithMask {
	ipWithMask := &IPv6WithMask{}
	ip, mask, err := parseIpv6AndMask(s)
	if err != nil {
		ipWithMask.IP = nil
		ipWithMask.parseError = err
		return ipWithMask
	}
	ipWithMask.Mask = mask
	ipWithMask.IP = ip
	return ipWithMask

}

func (i *IPv6WithMask) MarshalBinary() (data []byte, err error) {
	if i.IP == nil {
		return nil, fmt.Errorf("IPv6 is nil")
	}
	data = append(data, i.Mask)
	ipv6 := i.IP.To16()
	if ipv6 == nil {
		return nil, fmt.Errorf("IP.to16() is nil")
	}
	data = append(data, ipv6...)
	return data, err
}

func (i *IPv6WithMask) UnmarshalBinary(data []byte) (err error) {
	byteReader := NewByteReader(data)
	i.Mask, err = byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("inadequate length for mask: %v", err)
	}
	i.IP, err = byteReader.ReadBytes(net.IPv6len)
	if err != nil {
		return fmt.Errorf("inadequate length for IPv6: %v", err)
	}
	return err
}

package pfcpie

import (
	"fmt"
	"net"
)

// HostPool IE.
// Refer to "R6.3-GCS-3792-VzW-Feature-Gap-SMF shall support VzW required
// Pre-defined Rules."
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                        Type = 32788(decimal)                   |
//	         |----------------------------------------------------------------|
//	3 to 4   |                          Length = n                            |
//	         |----------------------------------------------------------------|
//	5 to 6   |                        Enterprise ID                           |
//	         |-------------------------------+-------+-------+-------+--------|
//	  7      |               Spare           |   V4  |   V6  |  V4   |  V6    |
//	         |                               |       |       | Range | Range  |
//	         |-------------------------------+-------+-------+-------+--------|
//	  8      |                       Length of Pool Name                      |
//	         |----------------------------------------------------------------|
//	9 to 9+a |                            Pool Name                           |
//	         |----------------------------------------------------------------|
//	  k      |                        Number of V4 Address                    |
//	         |----------------------------------------------------------------|
//	l to l+3 |                           IPv4 With Mask                       |
//	         |----------------------------------------------------------------|
//	  ...    |                               ...                              |
//	         |----------------------------------------------------------------|
//	m to m+3 |                         IPv4 With Mask                         |
//	         |----------------------------------------------------------------|
//	  n      |                        Number of V6 Address                    |
//	         |----------------------------------------------------------------|
//
// o to o+15 |                         IPv6 With Mask                         |
//
//	        |----------------------------------------------------------------|
//	...     |                               ...                              |
//	        |----------------------------------------------------------------|
//
// p to p+15 |                         IPv6 With Mask                         |
//
//	         |----------------------------------------------------------------|
//	  q      |                        Number of V4 Range                      |
//	         |----------------------------------------------------------------|
//	r to r+3 |                        Start Ipv4Address                       |
//	s to s+3 |                         End Ipv4Address                        |
//	         |----------------------------------------------------------------|
//	 ...     |                               ...                              |
//	         |----------------------------------------------------------------|
//	t to t+3 |                        Start Ipv4Address                       |
//	u to u+3 |                         End Ipv4Address                        |
//	         |----------------------------------------------------------------|
//	  v      |                        Number of V6 Range                      |
//	         |----------------------------------------------------------------|
//
// w to w+15 |                        Start Ipv6Address                       |
// x to x+15 |                         End Ipv6Address                        |
//
//	        |----------------------------------------------------------------|
//	...     |                               ...                              |
//	        |----------------------------------------------------------------|
//
// y to y+15 |                        Start Ipv6Address                       |
// z to z+15 |                         End Ipv6Address                        |
//
//	         |----------------------------------------------------------------|
//	k to n+4 |   These octets(s) is/are present only if explicity specified   |
//	         +----------------------------------------------------------------+
//
// Host pool IE can have both address and address range at the same time!
//
// IE Type: 32788 (0x8014)
type HostPool struct {
	V4           bool           `json:"v4,omitempty"`
	V6           bool           `json:"v6,omitempty"`
	V4Range      bool           `json:"v4Range,omitempty"`
	V6Range      bool           `json:"v6Range,omitempty"`
	EnterpriseID uint16         `json:"enterpriseID,omitempty"`
	PoolName     string         `json:"poolName,omitempty"` // Host pool name
	Ipv4Addrs    []IPv4WithMask `json:"ipv4Addrs,omitempty"`
	Ipv6Addrs    []IPv6WithMask `json:"ipv6Addrs,omitempty"`
	Ipv4Ranges   []IPRange      `json:"ipv4Ranges,omitempty"`
	Ipv6Ranges   []IPRange      `json:"ipv6Ranges,omitempty"`
}

// NewHostPool allocate a new HostPool IE.
func NewHostPool() *HostPool {
	return &HostPool{
		EnterpriseID: CasaEnterpriseID,
	}
}

// encodeIpv4Addrs encodes HostPool.Ipv4Addrs to octets.
func (h *HostPool) encodeIpv4Addrs() (data []byte, err error) {
	numOfV4 := uint8(len(h.Ipv4Addrs))
	data = append(data, byte(numOfV4))
	for _, ip := range h.Ipv4Addrs {
		if ip.IP.To4() == nil {
			return nil, fmt.Errorf("wrong ip format, should be IPV4, got: %v ", ip)
		}
		ipv4, _ := ip.MarshalBinary()
		data = append(data, ipv4...)
	}
	return data, nil
}

// encodeIpv6Addrs encodes HostPool.Ipv6Addrs to octets.
func (h *HostPool) encodeIpv6Addrs() (data []byte, err error) {
	numOfV6 := uint8(len(h.Ipv6Addrs))
	data = append(data, byte(numOfV6))
	for _, ip := range h.Ipv6Addrs {
		if ip.IP.To16() == nil {
			return nil, fmt.Errorf("wrong ip format, should be IPV6, got: %v", ip)
		}
		ipv6, _ := ip.MarshalBinary()
		data = append(data, ipv6...)
	}
	return data, nil
}

// encodeIpv4Ranges encodes HostPool.Ipv4Ranges to octets.
func (h *HostPool) encodeIpv4Ranges() (data []byte, err error) {
	numOfRanges := uint8(len(h.Ipv4Ranges))
	data = append(data, byte(numOfRanges))
	for _, r := range h.Ipv4Ranges {
		s := r.StartAddr.To4()
		if s == nil {
			errStr := "encodeIpv4Ranges, IpRange.StartAddr.To4() failed, Ip %v is not a valid Ipv4Address."
			return nil, fmt.Errorf(errStr, r.StartAddr)
		}
		data = append(data, s...)
		e := r.EndAddr.To4()
		if e == nil {
			errStr := "encodeIpv4Ranges, IpRange.EndAddr.To4() failed, Ip %v is not a valid Ipv4Address."
			return nil, fmt.Errorf(errStr, r.EndAddr)
		}
		data = append(data, e...)
	}
	return data, nil
}

// encodeIpv6Ranges encodes HostPool.Ipv6Ranges to octets.
func (h *HostPool) encodeIpv6Ranges() (data []byte, err error) {
	numOfRanges := uint8(len(h.Ipv6Ranges))
	data = append(data, byte(numOfRanges))
	for _, r := range h.Ipv6Ranges {
		s := r.StartAddr.To16()
		if s == nil {
			errStr := "encodeIpv6Ranges, IpRange.StartAddr.To16() failed, Ip %v is not a valid Ipv6Address."
			return nil, fmt.Errorf(errStr, r.StartAddr)
		}
		data = append(data, s...)
		e := r.EndAddr.To16()
		if e == nil {
			errStr := "encodeIpv6Ranges, IpRange.EndAddr.To16() failed, Ip %v is not a valid Ipv6Address."
			return nil, fmt.Errorf(errStr, r.EndAddr)
		}
		data = append(data, e...)
	}
	return data, nil
}

// MarshalBinary marshal HostPool IE to octets data.
func (h *HostPool) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 6: EnterpriseId
	data = AppendUint16([]byte(""), h.EnterpriseID)

	// Octet 7: V4, V6, V4Range, V6Range
	var tmp uint8 = btou(h.V4)<<3 |
		btou(h.V6)<<2 |
		btou(h.V4Range)<<1 |
		btou(h.V6Range)
	data = append(data, byte(tmp))

	// Octet 8: Length of Pool Name
	lengthOfPoolName := uint8(len(h.PoolName))
	data = append(data, byte(lengthOfPoolName))

	// Octet 9 to 9+a: Pool Name
	data = append(data, []byte(h.PoolName)...)

	// Octet k to m+3: Ipv4Addrs
	if h.V4 {
		v4Octets, err := h.encodeIpv4Addrs()
		if err != nil {
			return nil, err
		}
		data = append(data, v4Octets...)
	}

	// Octet o to o+15: Ipv6Addrs
	if h.V6 {
		v6Octets, err := h.encodeIpv6Addrs()
		if err != nil {
			return nil, err
		}
		data = append(data, v6Octets...)
	}

	// Octet q to u+3: Ipv4Ranges
	if h.V4Range {
		b, err := h.encodeIpv4Ranges()
		if err != nil {
			return nil, err
		}
		data = append(data, b...)
	}

	// Octet q to u+3: Ipv4Ranges
	if h.V6Range {
		b, err := h.encodeIpv6Ranges()
		if err != nil {
			return nil, err
		}
		data = append(data, b...)
	}
	return data, nil
}

// UnmarshalBinary unmarshal HostPool IE.
func (h *HostPool) UnmarshalBinary(data []byte) (err error) {
	byteReader := NewByteReader(data)

	// EnterpriseID
	h.EnterpriseID, err = byteReader.ReadUint16()
	if err != nil {
		return fmt.Errorf("Inadequate length for EnterpriseId : %v", err)
	}

	// octet 7: Flags
	tmp, err := byteReader.ReadByte()
	if err != nil {
		return fmt.Errorf("Inadequate length for flags, %v", err)
	}
	h.V4 = utob(tmp & BitMask4)
	h.V6 = utob(tmp & BitMask3)
	h.V4Range = utob(tmp & BitMask2)
	h.V6Range = utob(tmp & BitMask1)

	// Length of Pool Name
	lengthOfPoolName, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for lengthOfPoolName, %v", err)
	}

	// Pool Name
	poolName, err := byteReader.ReadBytes(int(lengthOfPoolName))
	h.PoolName = string(poolName)
	if err != nil {
		return fmt.Errorf("Inadequate length for PoolName, %v", err)
	}

	// numberOfV4Addrs & Ipv4Addrs
	if h.V4 {
		numberOfV4Addrs, err := byteReader.ReadUint8()
		if err != nil {
			return fmt.Errorf("Inadequate length for numberOfV4Addrs, %v", err)
		}
		ipv4Bytes, err := byteReader.ReadBytes(IPv4WithMasklen * int(numberOfV4Addrs))
		ipv4Reader := NewByteReader(ipv4Bytes)
		if err != nil {
			return fmt.Errorf("Inadequate length for V4Addrs, %v", err)
		}
		for i := 0; i < int(numberOfV4Addrs); i++ {
			ipv4WithMask := IPv4WithMask{}
			ipv4WithMask.Mask, _ = ipv4Reader.ReadUint8()
			ipv4, _ := ipv4Reader.ReadBytes(net.IPv4len)
			ipv4WithMask.IP = net.IP(ipv4).To16()
			h.Ipv4Addrs = append(h.Ipv4Addrs, ipv4WithMask)
		}
	}

	// numberOfV6Addrs & Ipv6Addrs
	if h.V6 {
		numberOfV6Addrs, err := byteReader.ReadUint8()
		if err != nil {
			return fmt.Errorf("Inadequate length for numberOfV6Addrs, %v", err)
		}
		ipv6Bytes, err := byteReader.ReadBytes(IPv6WithMasklen * int(numberOfV6Addrs))
		ipv6Reader := NewByteReader(ipv6Bytes)
		if err != nil {
			return fmt.Errorf("Inadequate length for V6Addrs, %v", err)
		}
		for i := 0; i < int(numberOfV6Addrs); i++ {
			ipv6WithMask := IPv6WithMask{}
			ipv6WithMask.Mask, _ = ipv6Reader.ReadUint8()
			ipv6, _ := ipv6Reader.ReadBytes(net.IPv6len)
			ipv6WithMask.IP = net.IP(ipv6).To16()
			h.Ipv6Addrs = append(h.Ipv6Addrs, ipv6WithMask)
		}
	}

	// numberOfV4Range & Ipv4Ranges
	if h.V4Range {
		numberOfV4Range, err := byteReader.ReadUint8()
		if err != nil {
			return fmt.Errorf("Inadequate length for numberOfV4Ranges, %v", err)
		}
		ipv4RangeBytes, err := byteReader.ReadBytes(IPv4Rangelen * int(numberOfV4Range))
		ipv4RangeReader := NewByteReader(ipv4RangeBytes)
		if err != nil {
			return fmt.Errorf("Inadequate length for V4Ranges, %v", err)
		}
		for i := 0; i < int(numberOfV4Range); i++ {
			ipv4Range := IPRange{}
			ipv4Start, _ := ipv4RangeReader.ReadBytes(net.IPv4len)
			ipv4Range.StartAddr = net.IP(ipv4Start).To16()
			ipv4End, _ := ipv4RangeReader.ReadBytes(net.IPv4len)
			ipv4Range.EndAddr = net.IP(ipv4End).To16()
			h.Ipv4Ranges = append(h.Ipv4Ranges, ipv4Range)
		}
	}

	// numberOfv6Range & Ipv6Ranges
	if h.V6Range {
		numberOfV6Range, err := byteReader.ReadUint8()
		if err != nil {
			return fmt.Errorf("Inadequate length for numberOfV6Ranges, %v", err)
		}
		ipv6RangeBytes, err := byteReader.ReadBytes(IPv6Rangelen * int(numberOfV6Range))
		ipv6RangeReader := NewByteReader(ipv6RangeBytes)
		if err != nil {
			return fmt.Errorf("Inadequate length for V6Ranges, %v", err)
		}
		for i := 0; i < int(numberOfV6Range); i++ {
			ipv6Range := IPRange{}
			ipv6Start, _ := ipv6RangeReader.ReadBytes(net.IPv6len)
			ipv6Range.StartAddr = net.IP(ipv6Start).To16()
			ipv6End, _ := ipv6RangeReader.ReadBytes(net.IPv6len)
			ipv6Range.EndAddr = net.IP(ipv6End).To16()
			h.Ipv6Ranges = append(h.Ipv6Ranges, ipv6Range)
		}
	}
	return nil
}

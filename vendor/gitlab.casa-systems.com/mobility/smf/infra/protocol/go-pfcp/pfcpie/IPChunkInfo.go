package pfcpie

import (
	"fmt"
)

type IPv4Subnet struct {
	IPv4WithMask
}

func NewIPv4Subnet(s string) *IPv4Subnet {
	a := &IPv4Subnet{
		IPv4WithMask: *NewIPv4WithMask(s),
	}
	return a
}

type IPv6Subnet struct {
	IPv6WithMask
}

func NewIPv6Subnet(s string) *IPv6Subnet {
	a := &IPv6Subnet{
		IPv6WithMask: *NewIPv6WithMask(s),
	}
	return a
}

// IPChunkInfo
// Casa Specific IE for Verizon use
//
//	Octets        8       7       6       5       4       3       2       1
//	          +---------------------------------------------------------------+
//	   1      |                          Length = n                           |
//	          |---------------------------------------------------------------|
//	   2      |                        VrfLen = (p-3)+1                       |
//	          |---------------------------------------------------------------|
//	3 to p    |                             Vrf                               |
//	          |---------------------------------------------------------------|
//	  p+1     |                    PoolNameLen = (q-(p+2))+1                  |
//	          |---------------------------------------------------------------|
//	p+2 to q  |                           PoolName                            |
//	          |---------------------------------------------------------------|
//	  q+1     |IsIpv6(bit8)|           NumberofSubnet(bit7~bit1)              |
//	          +---------------------------------------------------------------+
//	q+2 to n  |                      IPv4/IPv6 Subnets                        |
//	          +---------------------------------------------------------------+
//
// Ipv4 Subnets
//
//	      |----------------------------------------------------------------|
//	r     |                            SubnetLen                           |
//
// r+1 to r+4 |                          Ipv4NetAddress                        |
//
//	        |----------------------------------------------------------------|
//	...     |                               ...                              |
//	        |----------------------------------------------------------------|
//	 t      |                            SubnetLen                           |
//
// t+1 to t+4 |                         Ipv4NetAddress                         |
//
//	|----------------------------------------------------------------|
//
// Ipv6 Subnets
//
//	       |----------------------------------------------------------------|
//	w      |                             SubnetLen                          |
//
// w+1 to w+15|                           Ipv6NetAddress                       |
//
//	        |----------------------------------------------------------------|
//	...     |                               ...                              |
//	        |----------------------------------------------------------------|
//	 y      |                             SubnetLen                          |
//
// y+1 to y+15|                           Ipv6NetAddress                       |
//
//	|----------------------------------------------------------------|
//
// Ipv4Subnets: 10.1.1.0/24
// Ipv6Subnets: 2021:1:2:1::/64
// Each IpPool can only be equipped with one ip type(Ipv4/Ipv6), so there is one and only one of Ipv4Subnets and Ipv6Subnets.
type IPChunkInfo struct {
	Vrf         string       `json:"vrf,omitempty"`
	PoolName    string       `json:"poolName,omitempty"`
	IPv4Subnets []IPv4Subnet `json:"iPv4Subnets,omitempty"`
	IPv6Subnets []IPv6Subnet `json:"iPv6Subnets,omitempty"`
}

func (a *IPChunkInfo) MarshalBinary() (data []byte, err error) {
	// VrfLen
	vrfLen := len(a.Vrf)
	data = append(data, uint8(vrfLen))
	// Vrf
	if vrfLen > 0 {
		data = append(data, []byte(a.Vrf)...)
	}
	// PoolNameLen
	data = append(data, uint8(len(a.PoolName)))
	// PoolName
	data = append(data, []byte(a.PoolName)...)

	isIPv6 := false
	numberOfSubnet := len(a.IPv4Subnets)
	if len(a.IPv6Subnets) != 0 {
		isIPv6 = true
		numberOfSubnet = len(a.IPv6Subnets)
	}
	tmp := btou(isIPv6)<<7 | uint8(numberOfSubnet)
	data = append(data, tmp)

	if !isIPv6 {
		// Ipv4 Subnets
		for _, sub := range a.IPv4Subnets {
			buf, err := sub.MarshalBinary()
			if err != nil {
				return nil, err
			}
			data = append(data, buf...)
		}
	} else {
		// Ipv6 Subnets
		for _, sub := range a.IPv6Subnets {
			buf, err := sub.MarshalBinary()
			if err != nil {
				return nil, err
			}
			data = append(data, buf...)
		}
	}
	return data, nil
}

func (a *IPChunkInfo) UnmarshalBinary(data []byte) (err error) {
	byteReader := NewByteReader(data)
	// Vrf Len
	vrfLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for vrfLen: %v", err)
	}

	// Vrf
	if vrfLen > 0 {
		vrf, err := byteReader.ReadBytes(int(vrfLen))
		if err != nil {
			return fmt.Errorf("Inadequate length for vrf: %v", err)
		}
		a.Vrf = string(vrf)
	}

	// PoolName Len
	poolNameLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for poolNameLen: %v", err)
	}

	// PoolName
	poolName, err := byteReader.ReadBytes(int(poolNameLen))
	if err != nil {
		return fmt.Errorf("Inadequate length for poolName: %v", err)
	}
	a.PoolName = string(poolName)

	// IsIpv6 and NumberOfSubnet
	tmpOctet, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for IsIpv6 and NumberOfSubnet: %v", err)
	}

	isIPv6 := utob(tmpOctet & uint8(BitMask8))
	numberOfSubnet := tmpOctet & uint8(Mask7)

	// Ipv4Subnets or Ipv6Subnets
	sizeOfSubnet := IPv4WithMasklen
	if isIPv6 {
		sizeOfSubnet = IPv6WithMasklen
	}
	subnetsBytes, err := byteReader.ReadBytes(int(numberOfSubnet) * sizeOfSubnet)
	if err != nil {
		return fmt.Errorf("Inadequate length for IpSubnets: %v", err)
	}
	subnetsReader := NewByteReader(subnetsBytes)
	for i := 0; i < int(numberOfSubnet); i++ {
		subnetBuf, err := subnetsReader.ReadBytes(sizeOfSubnet)
		if err != nil {
			return fmt.Errorf("Inadequate length for Subnets[%d]: %v", i, err)
		}
		if !isIPv6 {
			ip := IPv4Subnet{}
			err = ip.UnmarshalBinary(subnetBuf)
			a.IPv4Subnets = append(a.IPv4Subnets, ip)
		} else {
			ip := IPv6Subnet{}
			err = ip.UnmarshalBinary(subnetBuf)
			a.IPv6Subnets = append(a.IPv6Subnets, ip)
		}
		if err != nil {
			return fmt.Errorf("unmarshal subnet[%d] faile: %v", i, err)
		}
	}
	return nil
}

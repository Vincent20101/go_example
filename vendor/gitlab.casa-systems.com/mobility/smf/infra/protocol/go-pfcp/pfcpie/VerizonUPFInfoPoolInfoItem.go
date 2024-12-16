package pfcpie

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Each IpPoolsItem defines IPv4 pool(s) (including private and public) and/or IPv6 pool(s) (including private and public) for a vrf.
// Casa Specific IE for Verizon use
//
//	Octets        8       7       6       5       4       3       2       1
//	          +--------------------------------------------------------------+
//	1 to 2    |                   Length = n - 2                             |
//	          |--------------------------------------------------------------|
//	3         | vrf | v4priv  | v4pub | v6priv | v6pub|                      |
//	          |--------------------------------------------------------------|
//	4         |       Vrf Length = (p-4)+1 if vrf bit is set in Octet 2      |
//	          |--------------------------------------------------------------|
//	5 to p    |                   Vrf  if vrf bit is set in Octet 2          |
//	          |--------------------------------------------------------------|
//	p+1 to q  |   IpPool for v4 private if v4priv bit is set in Octet 2      |
//            |--------------------------------------------------------------|
//  q+1 to r  |  IpPool for v4 public if v4pub bit is set in Octet 2         |
//            |--------------------------------------------------------------|
//  r+1 to s  |  IpPool for v6 private if v6priv bit is set in Octet 2       |
//            |--------------------------------------------------------------|
//  s+1 to n  |  IpPool for v6 public if v6pub bit is set in Octet 2         |
//	          +--------------------------------------------------------------+

type UpfIpPoolInfoItem struct {
	Vrf           string       `json:"vrf,omitempty"`
	PublicV4Pool  *UpfInfoPool `json:"publicV4Pool,omitempty"`
	PrivateV4Pool *UpfInfoPool `json:"privateV4Pool,omitempty"`
	PublicV6Pool  *UpfInfoPool `json:"publicV6Pool,omitempty"`
	PrivateV6Pool *UpfInfoPool `json:"privateV6Pool,omitempty"`
}

func (a *UpfIpPoolInfoItem) MarshalBinary() (data []byte, err error) {
	// Bit set
	data = append(data, uint8(0))
	// VrfLen
	vrfLen := len(a.Vrf)
	// Vrf
	hasVrf := false
	if vrfLen > 0 {
		hasVrf = true
		data = append(data, uint8(vrfLen))
		data = append(data, []byte(a.Vrf)...)
	}
	//PrivateV4Pool
	hasPriV4 := false
	if a.PrivateV4Pool != nil && len(a.PrivateV4Pool.V4Ranges) > 0 {
		hasPriV4 = true
		d, e := a.PrivateV4Pool.MarshalBinary()
		if e != nil {
			return nil, e
		}
		dl := uint16(len(d))
		data = append(data, byte(dl>>8), byte(dl))
		data = append(data, d...)
	}
	//PublicV4Pool
	hasPubV4 := false
	if a.PublicV4Pool != nil && len(a.PublicV4Pool.V4Ranges) > 0 {
		hasPubV4 = true
		d, e := a.PublicV4Pool.MarshalBinary()
		if e != nil {
			return nil, e
		}
		dl := uint16(len(d))
		data = append(data, byte(dl>>8), byte(dl))
		data = append(data, d...)
	}
	//PrivateV6Pool
	hasPriV6 := false
	if a.PrivateV6Pool != nil && len(a.PrivateV6Pool.V6Ranges) > 0 {
		hasPriV6 = true
		d, e := a.PrivateV6Pool.MarshalBinary()
		if e != nil {
			return nil, e
		}
		dl := uint16(len(d))
		data = append(data, byte(dl>>8), byte(dl))
		data = append(data, d...)
	}
	//PublicV6Pool
	hasPubV6 := false
	if a.PublicV6Pool != nil && len(a.PublicV6Pool.V6Ranges) > 0 {
		hasPubV6 = true
		d, e := a.PublicV6Pool.MarshalBinary()
		if e != nil {
			return nil, e
		}
		dl := uint16(len(d))
		data = append(data, byte(dl>>8), byte(dl))
		data = append(data, d...)
	}

	tmp := btou(hasVrf)<<7 | btou(hasPriV4)<<6 | btou(hasPubV4)<<5 | btou(hasPriV6)<<4 | btou(hasPubV6)<<3
	data[0] = tmp

	return data, nil
}

func (a *UpfIpPoolInfoItem) UnmarshalBinary(data []byte) (err error) {
	byteReader := NewByteReader(data)

	bitSet, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read bit set failed,err:%s", err.Error())
	}

	hasVrf := utob(bitSet & uint8(BitMask8))
	hasPriV4 := utob(bitSet & uint8(BitMask7))
	hasPubV4 := utob(bitSet & uint8(BitMask6))
	hasPriV6 := utob(bitSet & uint8(BitMask5))
	hasPubV6 := utob(bitSet & uint8(BitMask4))

	if hasVrf {
		vrfLen, err := byteReader.ReadUint8()
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read vrf len failed,err:%s", err.Error())
		}
		vrf, err := byteReader.ReadBytes(int(vrfLen))
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read vrf failed,err:%s", err.Error())
		}
		a.Vrf = string(vrf)
	}

	if hasPriV4 {
		dataLen, err := byteReader.ReadUint16()
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read upf ippool len failed,err:%s", err.Error())
		}
		data, err := byteReader.ReadBytes(int(dataLen))
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read upf ippool failed,err:%s", err.Error())
		}
		a.PrivateV4Pool = &UpfInfoPool{}
		err = a.PrivateV4Pool.UnmarshalBinary(data)
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read PrivateV4Pool failed,err:%s", err.Error())
		}
	}

	if hasPubV4 {
		dataLen, err := byteReader.ReadUint16()
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read upf ippool len failed,err:%s", err.Error())
		}
		data, err := byteReader.ReadBytes(int(dataLen))
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read upf ippool failed,err:%s", err.Error())
		}
		a.PublicV4Pool = &UpfInfoPool{}
		err = a.PublicV4Pool.UnmarshalBinary(data)
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read PublicV4Pool failed,err:%s", err.Error())
		}
	}

	if hasPriV6 {
		dataLen, err := byteReader.ReadUint16()
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read upf ippool len failed,err:%s", err.Error())
		}
		data, err := byteReader.ReadBytes(int(dataLen))
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read upf ippool failed,err:%s", err.Error())
		}
		a.PrivateV6Pool = &UpfInfoPool{IsIpv6: true}
		err = a.PrivateV6Pool.UnmarshalBinary(data)
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read PrivateV6Pool failed,err:%s", err.Error())
		}
	}

	if hasPubV6 {
		dataLen, err := byteReader.ReadUint16()
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read upf ippool len failed,err:%s", err.Error())
		}
		data, err := byteReader.ReadBytes(int(dataLen))
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read upf ippool failed,err:%s", err.Error())
		}
		a.PublicV6Pool = &UpfInfoPool{IsIpv6: true}
		err = a.PublicV6Pool.UnmarshalBinary(data)
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read PublicV6Pool failed,err:%s", err.Error())
		}
	}

	return nil
}

// UpfInfoPool
// An IpPool includes either a list of IPv4 address range or a list of IPv6 address range.
//	Octets        8       7       6       5       4       3       2       1
//	          +--------------------------------------------------------------+
//	1 to 2    |                   Length = n - 2                             |
//	          |--------------------------------------------------------------|
//	3         |            Number of IpAddressRange(s)                       |
//	          |--------------------------------------------------------------|
//	4 to n    |       list of IpAddressRange(s) (IPv4 or IPv6)               |
//	          +--------------------------------------------------------------+

type UpfInfoPool struct {
	V4Ranges []Ipv4AddressRange `json:"v4Ranges,omitempty"`
	V6Ranges []Ipv6AddressRange `json:"v6Ranges,omitempty"`
	IsIpv6   bool               `cmp:"ignore"` //pass by UpfIpPoolInfoItem
}

func (a *UpfInfoPool) MarshalBinary() (data []byte, err error) {
	if len(a.V4Ranges) > 0 {
		data = append(data, uint8(len(a.V4Ranges)))
		for _, v := range a.V4Ranges {
			d, err := v.MarshalBinary()
			if err != nil {
				return nil, err
			}
			data = append(data, d...)
		}
	} else if len(a.V6Ranges) > 0 {
		data = append(data, uint8(len(a.V6Ranges)))
		for _, v := range a.V6Ranges {
			d, err := v.MarshalBinary()
			if err != nil {
				return nil, err
			}
			data = append(data, d...)
		}
	}

	return data, nil
}

func (a *UpfInfoPool) UnmarshalBinary(data []byte) (err error) {
	byteReader := NewByteReader(data)

	rangeLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("UpfInfoPool.UnmarshalBinary read range len failed,err:%s", err.Error())
	}

	var rangeSize int
	if a.IsIpv6 {
		rangeSize = int(rangeLen * IPv6WithMasklen * 2)
		subnetsBytes, err := byteReader.ReadBytes(rangeSize)
		if err != nil {
			return fmt.Errorf("UpfInfoPool.UnmarshalBinary,read ranges failed: %v", err)
		}
		subnetsReader := NewByteReader(subnetsBytes)

		for i := 0; i < int(rangeLen); i++ {
			subnetBuf, err := subnetsReader.ReadBytes(IPv6WithMasklen * 2)
			if err != nil {
				return fmt.Errorf("UpfInfoPool.UnmarshalBinary,read range failed[%d]: %v", i, err)
			}
			ip := Ipv6AddressRange{}
			err = ip.UnmarshalBinary(subnetBuf)
			if err != nil {
				return fmt.Errorf("UpfInfoPool.UnmarshalBinary,read v6 range failed[%d]: %v", i, err)
			}
			a.V6Ranges = append(a.V6Ranges, ip)
		}
		return nil
	}

	rangeSize = int(rangeLen * net.IPv4len * 2)
	subnetsBytes, err := byteReader.ReadBytes(rangeSize)
	if err != nil {
		return fmt.Errorf("UpfInfoPool.UnmarshalBinary,read ranges failed: %v", err)
	}
	subnetsReader := NewByteReader(subnetsBytes)

	for i := 0; i < int(rangeLen); i++ {
		subnetBuf, err := subnetsReader.ReadBytes(net.IPv4len * 2)
		if err != nil {
			return fmt.Errorf("UpfInfoPool.UnmarshalBinary,read range failed[%d]: %v", i, err)
		}
		ip := Ipv4AddressRange{}
		err = ip.UnmarshalBinary(subnetBuf)
		if err != nil {
			return fmt.Errorf("UpfInfoPool.UnmarshalBinary,read v4 range failed[%d]: %v", i, err)
		}
		a.V4Ranges = append(a.V4Ranges, ip)
	}
	return nil
}

// Ipv4AddressRange
//	Octets        8       7       6       5       4       3       2       1
//	          +--------------------------------------------------------------+
//	1 to 4    |                   start                                      |
//	          |--------------------------------------------------------------|
//	5 to 8    |                    end                                       |
//	          +--------------------------------------------------------------+

func NewIpv4AddressRange(start, end string) (Ipv4AddressRange, error) {
	return Ipv4AddressRange{StartAddr: start, EndAddr: end}, nil
}

type Ipv4AddressRange struct {
	StartAddr string `json:"startAddr,omitempty"`
	EndAddr   string `json:"endAddr,omitempty"`
}

func (a *Ipv4AddressRange) MarshalBinary() (data []byte, err error) {
	if a.EndAddr == "" || a.StartAddr == "" {
		return nil, fmt.Errorf("Ipv4AddressRange start or end address is nil")
	}

	// Parse start IP
	sa := net.ParseIP(a.StartAddr).To4()
	if sa == nil {
		return nil, fmt.Errorf("NewIpv4AddressRange,ipv4 addr %s not valid", a.StartAddr)
	}

	// Parse end IP
	ea := net.ParseIP(a.EndAddr).To4()
	if ea == nil {
		return nil, fmt.Errorf("NewIpv4AddressRange,ipv4 addr %s not valid", a.EndAddr)
	}

	data = append(data, sa...)

	data = append(data, ea...)

	return data, err
}

func (a *Ipv4AddressRange) UnmarshalBinary(data []byte) (err error) {
	byteReader := NewByteReader(data)
	var address net.IP

	address, err = byteReader.ReadBytes(net.IPv4len)
	if err != nil {
		return fmt.Errorf("Ipv4AddressRange.UnmarshalBinary failed,inadequate length for IPv4: %v", err)
	}
	a.StartAddr = address.String()

	address, err = byteReader.ReadBytes(net.IPv4len)
	if err != nil {
		return fmt.Errorf("Ipv4AddressRange.UnmarshalBinary failed,inadequate length for IPv4: %v", err)
	}
	a.EndAddr = address.String()

	return err
}

//Ipv6AddressRange,The first byte is the ipv6 prefix length in both start and end fields.
//	Octets        8       7       6       5       4       3       2       1
//	          +--------------------------------------------------------------+
//	1 to 17   |                    start                                     |
//	          |--------------------------------------------------------------|
//	18 to 34  |                      end                                     |
//	          +--------------------------------------------------------------+

func NewIpv6AddressRange(start, end string) (Ipv6AddressRange, error) {
	r := Ipv6AddressRange{}

	// Parse start IP and Mask
	res := strings.Split(start, "/")
	if len(res) != 2 {
		return r, fmt.Errorf("NewIpv6AddressRange failed,ipv6 must be in the format of ip/mask, but got %s", start)
	}
	ipStr, maskStr := res[0], res[1]

	r.StartAddr = net.ParseIP(ipStr)
	maskNum, err := strconv.Atoi(maskStr)
	if err != nil {
		return r, fmt.Errorf("mask %s not valid", maskStr)
	}
	r.StartPrefixLen = uint8(maskNum)

	// Parse end IP and Mask
	res = strings.Split(end, "/")
	if len(res) != 2 {
		return r, fmt.Errorf("NewIpv6AddressRange failed,ipv6 must be in the format of ip/mask, but got %s", end)
	}
	ipStr, maskStr = res[0], res[1]

	// Parse IP and Mask
	r.EndAddr = net.ParseIP(ipStr)
	maskNum, err = strconv.Atoi(maskStr)
	if err != nil {
		return r, fmt.Errorf("mask %s not valid", maskStr)
	}
	r.EndPrefixLen = uint8(maskNum)

	return r, nil
}

type Ipv6AddressRange struct {
	StartPrefixLen uint8  `json:"startPrefixLen,omitempty"`
	EndPrefixLen   uint8  `json:"endPrefixLen,omitempty"`
	StartAddr      net.IP `json:"startAddr,omitempty"`
	EndAddr        net.IP `json:"endAddr,omitempty"`
}

func (a *Ipv6AddressRange) MarshalBinary() (data []byte, err error) {
	if a.EndAddr == nil || a.StartAddr == nil {
		return nil, fmt.Errorf("Ipv6AddressRange start or end address is nil")
	}
	data = append(data, a.StartPrefixLen)
	ipv6 := a.StartAddr.To16()
	if ipv6 == nil {
		return nil, fmt.Errorf("StartAddr.to16() is nil,address:%v", a.StartAddr)
	}
	data = append(data, ipv6...)

	data = append(data, a.EndPrefixLen)
	ipv6 = a.EndAddr.To16()
	if ipv6 == nil {
		return nil, fmt.Errorf("EndAddr.to16() is nil,address:%v", a.EndAddr)
	}
	data = append(data, ipv6...)

	return data, err
}

func (a *Ipv6AddressRange) UnmarshalBinary(data []byte) (err error) {
	byteReader := NewByteReader(data)
	a.StartPrefixLen, err = byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("inadequate length for mask: %v", err)
	}
	a.StartAddr, err = byteReader.ReadBytes(net.IPv6len)
	if err != nil {
		return fmt.Errorf("inadequate length for IPv6: %v", err)
	}

	a.EndPrefixLen, err = byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("inadequate length for mask: %v", err)
	}
	a.EndAddr, err = byteReader.ReadBytes(net.IPv6len)
	if err != nil {
		return fmt.Errorf("inadequate length for IPv6: %v", err)
	}

	return err
}

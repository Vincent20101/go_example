package pfcpie

import (
	"fmt"
	"net"
)

//// SourceIPAddress IE. Refer to [TS29244 8.2.138 Source IP Address]
//
//// It shall be encoded as shown in Figure Figure 8.2.138-1.
//
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                    Type = 192 (decimal)                        |
//           |----------------------------------------------------------------|
//  3 to 4   |                      Length = n                                |
//           |----------------------------------------------------------------|
//    5      |             Spare                     |  MPL  |   V4   |  V6   |
//           |----------------------------------------------------------------|
// m to m+3  |                          IPv4 address                          |
//           |----------------------------------------------------------------|
// p to p+15 |                          IPv6 address                          |
//           |----------------------------------------------------------------|
//    q      |                        mask/prefix length                      |
//           |----------------------------------------------------------------|
// k to (n+4)|   These octets(s) is/are present only if explicity specified   |
//           +----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//
//    -	Bit 2 – V4: If this bit is set to "1", then the IPv4 address field shall be
//		present, otherwise the IPv4 address field shall not be present.
//
//    -	Bit 3 – Mask/Prefix Length: If this bit is set to "1", then the mask
//		(for IPv4) / prefix (for IPv6) length field shall be present, otherwise
//		this field shall not be present.
//
//    -	Bit 4 to 8 Spare, for future use and set to "0".
//
//	  Octets "m to (m+3)", "p to (p+15)" (IPv4 address / IPv6 address fields), if present,
//	  shall contain the address value.
//
//    The mask/prefix length field, if present, shall be encoded as a 8 bits binary integer.
//
//    EXAMPLE 1: this field encodes the value 24 for the IPv4 subnet 192.0.2.10/24.
//    EXAMPLE 2: this field encodes the value 64 for the /64 IPv6 prefix.

type SourceIpAddress struct {
	MPL              bool   `json:"mPL,omitempty"`
	V4               bool   `json:"v4,omitempty"`
	V6               bool   `json:"v6,omitempty"`
	Ipv4Addr         net.IP `json:"ipv4Addr,omitempty"`
	Ipv6Addr         net.IP `json:"ipv6Addr,omitempty"`
	MaskPrefixLength uint8  `json:"maskPrefixLength,omitempty"`
}

func (s *SourceIpAddress) MarshalBinary() (data []byte, err error) {
	errMsgPrefix := "Marshal SourceIpAddress failed: %s"
	// Octet 5
	octet5 := btou(s.V6) | btou(s.V4)<<1 | btou(s.MPL)<<2
	data = append([]byte(""), byte(octet5))

	// Octets m to (m+3) or p to (p+15)
	if s.V4 {
		if len(s.Ipv4Addr) == 0 || s.Ipv4Addr.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix, "IPv4 address shall be present if V4 is set")
		}
		ipV4 := s.Ipv4Addr.To4()
		if ipV4 == nil {
			return []byte(""), fmt.Errorf(errMsgPrefix, " Encode IPv4 address fail: invaild ipv4: %v", s.Ipv4Addr)
		}
		data = append(data, ipV4...)
	}

	if s.V6 {
		if len(s.Ipv6Addr) == 0 || s.Ipv6Addr.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix, "IPv6 address shall be present if V6 is set")
		}
		ipV6 := s.Ipv6Addr.To16()
		if ipV6 == nil {
			return []byte(""), fmt.Errorf(errMsgPrefix, " Encode IPv6 address fail: invaild ipv6: %v", s.Ipv6Addr)
		}
		data = append(data, ipV6...)
	}
	//q
	if s.MPL {
		data = append(data, s.MaskPrefixLength)
	}
	return data, nil
}

func (s *SourceIpAddress) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))
	var idx uint16 = 0
	if length < idx+1 {
		return fmt.Errorf("ie SourceIpAddress Inadequate TLV length: %d", length)
	}
	s.V4 = utob(uint8(data[idx]) & BitMask2)
	s.V6 = utob(uint8(data[idx]) & BitMask1)
	s.MPL = utob(uint8(data[idx]) & BitMask3)
	idx = idx + 1
	if s.V4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("ie SourceIpAddress Inadequate TLV length: %d", length)
		}
		s.Ipv4Addr = net.IP(data[idx : idx+net.IPv4len])

		idx = idx + net.IPv4len
	}
	if s.V6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("ie SourceIpAddress Inadequate TLV length: %d", length)
		}
		s.Ipv6Addr = net.IP(data[idx : idx+net.IPv6len])

		idx = idx + net.IPv6len
	}

	if s.MPL {
		if length < idx+1 {
			return fmt.Errorf("ie SourceIpAddress Inadequate TLV length: %d", length)
		}
		s.MaskPrefixLength = uint8(data[idx])
		idx = idx + 1
	}
	return nil
}

package pfcpie

import (
	"encoding/binary"
	"fmt"
	"math"
	"net"
)

const (
	// Octet 5
	OuterHeaderCreationGtpUUdpIpv4 uint16 = 1
	OuterHeaderCreationGtpUUdpIpv6 uint16 = 1 << 1
	OuterHeaderCreationUdpIpv4     uint16 = 1 << 2
	OuterHeaderCreationUdpIpv6     uint16 = 1 << 3
	OuterHeaderCreationIPv4        uint16 = 1 << 4
	OuterHeaderCreationIPv6        uint16 = 1 << 5
	OuterHeaderCreationCtag        uint16 = 1 << 6
	OuterHeaderCreationStag        uint16 = 1 << 7
	// Octet 6
	OuterHeaderCreationN19Indication uint16 = 1 << 8
	OuterHeaderCreationN6Indication  uint16 = 1 << 9
)

// OuterHeaderCreation IE: Refer to [29244 8.2.56 Outer Header Creation]
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                       Type = 84(decimal)                       |
//	         |----------------------------------------------------------------|
//	3 to 4   |                          Length = n                            |
//	         |----------------------------------------------------------------|
//	5 to 6   |                Outer Header Creation Description               |
//	         |----------------------------------------------------------------|
//
// m to m+3  |                           TEID                                 |
//
//	|----------------------------------------------------------------|
//
// p to p+3  |                       IPv4 Address                             |
//
//	|----------------------------------------------------------------|
//
// q to q+3  |                       IPv6 Address                             |
//
//	|----------------------------------------------------------------|
//
// q to q+15 |                       IPv6 Address                             |
//
//	|----------------------------------------------------------------|
//
// r to r+1  |                        Port Number                             |
//
//	|----------------------------------------------------------------|
//
// t to t+2  |                           C-TAG                                |
//
//	|----------------------------------------------------------------|
//
// u to u+2  |                           S-TAG                                |
//
//	|----------------------------------------------------------------|
//
// s to (s+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The Outer Header Creation Description field, when present, shall be encoded
// as specified in Table 8.2.56-1. It takes the  form of a bitmask where each
// bit indicates the outer header to be created in the outgoing packet. Spare
// bits shall be ignored by the receiver.
//
// Check [29244 8.2.56 Outer Header Creation] to see more details.
type OuterHeaderCreation struct {
	OuterHeaderCreationDescription uint16 `json:"outerHeaderCreationDescription,omitempty"`
	Teid                           uint32 `json:"teid,omitempty"`
	Ipv4Address                    net.IP `json:"ipv4Address,omitempty"`
	Ipv6Address                    net.IP `json:"ipv6Address,omitempty"`
	PortNumber                     uint16 `json:"portNumber,omitempty"`
	Ctag                           *CTAG  `json:"ctag,omitempty"`
	Stag                           *STAG  `json:"stag,omitempty"`
}

// MarshalBinary marshals OuterHeaderCreation, Refer to [29244 8.2.56 Outer
// Header Creation].
func (o *OuterHeaderCreation) MarshalBinary() (data []byte, err error) {
	errMsgPrefix := "Marshal OuterHeaderCreation failed: "
	// Refer to [29244 8.2.56 Outer Header Creation], At least one bit of the
	// Outer Header Creation Description field shall be set to "1". Bits 5/1 and
	// 5/2 may both be set to "1".
	if o.OuterHeaderCreationDescription&0x03ff == 0 {
		return []byte(""), fmt.Errorf(errMsgPrefix+"At least one bit of outer header description field shall be set%x, ", o.OuterHeaderCreationDescription&0x0fff)
	}

	octet5 := uint8(o.OuterHeaderCreationDescription & Mask8)
	GtpU := utob(octet5&BitMask1) || utob(octet5&BitMask2)
	Udp := utob(octet5&BitMask3) || utob(octet5&BitMask4)
	Ipv4 := utob(octet5&BitMask1) || utob(octet5&BitMask3)
	Ipv6 := utob(octet5&BitMask2) || utob(octet5&BitMask4)
	CtagFlag := utob(octet5 & BitMask7)
	StagFlag := utob(octet5 & BitMask8)

	var idx uint16 = 0
	// Octet 5 to 6
	data = make([]byte, 2)
	binary.LittleEndian.PutUint16(data[idx:], o.OuterHeaderCreationDescription)
	idx = idx + 2

	// Octet m to (m+3)
	if GtpU {
		data = append(data, make([]byte, 4)...)
		binary.BigEndian.PutUint32(data[idx:], o.Teid)
		idx = idx + 4
	}

	// Octet p to (p+3)
	if Ipv4 {
		if len(o.Ipv4Address) == 0 /*|| o.Ipv4Address.IsUnspecified()// fix GCS-15496:UPF count idle session by check N3 peer ip set to 0*/ {
			return []byte(""), fmt.Errorf(errMsgPrefix + "IPv4 address must be present")
		}
		data = append(data, o.Ipv4Address.To4()...)
		idx = idx + net.IPv4len
	}

	// Octet q to (q+15)
	if Ipv6 {
		if len(o.Ipv6Address) == 0 /*|| o.Ipv6Address.IsUnspecified()//fix GCS-15496:UPF count idle session by check N3 peer ip set to 0*/ {
			return []byte(""), fmt.Errorf(errMsgPrefix + "IPv6 address must be present")
		}
		data = append(data, o.Ipv6Address.To16()...)
		idx = idx + net.IPv6len
	}

	// Octet r to (r+1)
	if Udp {
		data = append(data, make([]byte, 2)...)
		binary.BigEndian.PutUint16(data[idx:], o.PortNumber)
		idx += 2
	}

	// Octet t to (t+2)
	if CtagFlag {
		var ctagData []byte
		if o.Ctag != nil {
			var err error
			ctagData, err = o.Ctag.MarshalBinary()
			if err != nil {
				return nil, fmt.Errorf("OuterHeaderCreationDescription failed to marshal CTAG: %v", err)
			}
			if len(ctagData) > math.MaxUint8 {
				return nil, fmt.Errorf("Ctags length overflowed: %d", len(ctagData))
			}
		}
		data = append(data, ctagData...)
		idx += 3
	}

	// Octet u to (u+2)
	if StagFlag {
		var stagData []byte
		if o.Stag != nil {
			var err error
			stagData, err = o.Stag.MarshalBinary()
			if err != nil {
				return nil, fmt.Errorf("OuterHeaderCreationDescription failed to marshal STAG: %v", err)
			}
			if len(stagData) > math.MaxUint8 {
				return nil, fmt.Errorf("Stags length overflowed: %d", len(stagData))
			}
		}
		data = append(data, stagData...)
	}

	return data, nil
}

func (o *OuterHeaderCreation) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 6
	if length < idx+2 {
		return fmt.Errorf("IE OuterHeaderCreation Inadequate TLV length: %d", length)
	}
	o.OuterHeaderCreationDescription = binary.LittleEndian.Uint16(data[idx:])

	octet5 := uint8(o.OuterHeaderCreationDescription & Mask8)
	var GtpU, Udp, Ipv4, Ipv6 bool
	GtpU = utob(octet5&BitMask1) || utob(octet5&BitMask2)
	Udp = utob(octet5&BitMask3) || utob(octet5&BitMask4)
	Ipv4 = utob(octet5&BitMask1) || utob(octet5&BitMask3)
	Ipv6 = utob(octet5&BitMask2) || utob(octet5&BitMask4)
	CtagFlag := utob(octet5 & BitMask7)
	StagFlag := utob(octet5 & BitMask8)
	/*if !GtpU && !Udp {
		return fmt.Errorf("in IE OuterHeaderCreation GtpU or Udp of outer header description field are both none. You should set one")
	}*/

	idx = idx + 2
	// Octet m to (m+3)
	if GtpU {
		if length < idx+4 {
			return fmt.Errorf("IE OuterHeaderCreation Fieid Teid Inadequate TLV length: %d", length)
		}
		o.Teid = binary.BigEndian.Uint32(data[idx:])
		idx = idx + 4
	}

	// Octet p to (p+3)
	if Ipv4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("IE OuterHeaderCreation Fieid Ipv4Address Inadequate TLV length: %d", length)
		}
		o.Ipv4Address = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	}

	// Octet q to (q+15)
	if Ipv6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("IE OuterHeaderCreation Fieid Ipv6Address Inadequate TLV length: %d", length)
		}
		o.Ipv6Address = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	// Octet r to (r+2)
	if Udp {
		if length < idx+2 {
			return fmt.Errorf("IE OuterHeaderCreation Fieid Udp Inadequate TLV length: %d", length)
		}
		o.PortNumber = binary.BigEndian.Uint16(data[idx : idx+2])
		idx = idx + 2
	}

	// Octet t to (t+2)
	if CtagFlag {
		if length < idx+3 {
			return fmt.Errorf("IE OuterHeaderCreation Fieid Ctag Inadequate TLV length: %d", length)
		}
		var ctag CTAG
		if err := ctag.UnmarshalBinary(data[idx : idx+3]); err != nil {
			return fmt.Errorf("IE OuterHeaderCreation failed to unmarshal CTAG: %v", err)
		}
		o.Ctag = &ctag
		idx = idx + 3
	}

	// Octet u to (u+2)
	if StagFlag {
		if length < idx+3 {
			return fmt.Errorf("IE OuterHeaderCreation Fieid Stag Inadequate TLV fail,length: %d", length)
		}
		var stag STAG
		if err := stag.UnmarshalBinary(data[idx : idx+3]); err != nil {
			return fmt.Errorf("IE OuterHeaderCreation failed to unmarshal STAG: %v", err)
		}
		o.Stag = &stag
		idx = idx + 3
	}

	return nil
}

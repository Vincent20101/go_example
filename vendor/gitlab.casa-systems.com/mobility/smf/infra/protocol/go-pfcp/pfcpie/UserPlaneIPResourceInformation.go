package pfcpie

import (
	"fmt"
	"math/bits"
	"net"
)

type UserPlaneIPResourceInformation struct {
	Assosi          bool   `json:"assosi,omitempty"`
	Assoni          bool   `json:"assoni,omitempty"`
	Teidri          uint8  `json:"teidri,omitempty"` // 0x00011100
	V6              bool   `json:"v6,omitempty"`
	V4              bool   `json:"v4,omitempty"`
	TeidRange       uint8  `json:"teidRange,omitempty"`
	Ipv4Address     net.IP `json:"ipv4Address,omitempty" cmp:"ignore"`
	Ipv6Address     net.IP `json:"ipv6Address,omitempty" cmp:"ignore"`
	NetworkInstance Dnn    `json:"networkInstance,omitempty"`
	SourceInterface uint8  `json:"sourceInterface,omitempty"` // 0x00001111
}

func (u *UserPlaneIPResourceInformation) MarshalBinary() (data []byte, err error) {
	errMsgPrefix := "marshal UserPlaneIPResourceInformation failed:"
	// Octet 5
	if bits.Len8(u.Teidri) > 3 {
		return []byte(""), fmt.Errorf(errMsgPrefix + "TEIDRI shall not be greater than 3 bits binary integer")
	}
	tmpUint8 := btou(u.Assosi)<<6 |
		btou(u.Assoni)<<5 |
		u.Teidri<<2 |
		btou(u.V6)<<1 |
		btou(u.V4)
	data = append([]byte(""), byte(tmpUint8))

	// Octet 6
	if u.Teidri != 0 {
		data = append(data, byte(u.TeidRange))
	}

	// Octet m to (m+3)
	if u.V4 {
		if len(u.Ipv4Address) == 0 || u.Ipv4Address.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix + "IPv4 address shall be present if V4 is set")
		}
		data = append(data, u.Ipv4Address.To4()...)
	}

	// Octet p to (p+15)
	if u.V6 {
		if len(u.Ipv6Address) == 0 || u.Ipv6Address.IsUnspecified() {
			return []byte(""), fmt.Errorf(errMsgPrefix + "IPv6 address shall be present if V6 is set")
		}
		data = append(data, u.Ipv6Address.To16()...)
	}

	if !u.V4 && !u.V6 {
		return []byte(""), fmt.Errorf(errMsgPrefix + "At least one of V4 and V6 flags shall be set")
	}

	// Octet k to l
	if u.Assoni {
		AssoniBuf, _ := u.NetworkInstance.MarshalBinary()
		data = append(data, AssoniBuf...)
	}

	// Octet r
	if u.Assosi {
		if bits.Len8(u.SourceInterface) > 4 {
			return []byte(""), fmt.Errorf(errMsgPrefix + "Source interface shall not be greater than 4 bits binary integer")
		}
		data = append(data, byte(u.SourceInterface))
	}

	return data, nil
}

func (u *UserPlaneIPResourceInformation) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie UserPlaneIPResourceInformation Inadequate TLV length: %d", length)
	}
	u.Assosi = utob(uint8(data[idx]) & BitMask7)
	u.Assoni = utob(uint8(data[idx]) & BitMask6)
	u.Teidri = uint8(data[idx]) >> 2 & Mask3
	u.V6 = utob(uint8(data[idx]) & BitMask2)
	u.V4 = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet 6
	if u.Teidri != 0 {
		if length < idx+1 {
			return fmt.Errorf("ie UserPlaneIPResourceInformation.TeidRange Inadequate TLV length: %d", length)
		}
		u.TeidRange = uint8(data[idx])
		idx = idx + 1
	}

	// Octet m to (m+3)
	if u.V4 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("ie UserPlaneIPResourceInformation.Ipv4Address Inadequate TLV length: %d", length)
		}
		u.Ipv4Address = net.IP(data[idx : idx+net.IPv4len]).To16()
		idx = idx + net.IPv4len
	}

	// Octet p to (p+15)
	if u.V6 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("ie UserPlaneIPResourceInformation.Ipv6Address Inadequate TLV length: %d", length)
		}
		u.Ipv6Address = net.IP(data[idx : idx+net.IPv6len]).To16()
		idx = idx + net.IPv6len
	}

	if !u.V4 && !u.V6 {
		return fmt.Errorf("None of V4 and V6 flags is set")
	}

	// Octet r
	if u.Assosi {
		if length < idx+1 {
			return fmt.Errorf("ie UserPlaneIPResourceInformation.SourceInterface Inadequate TLV length: %d", length)
		}
		u.SourceInterface = data[length-1] & Mask4
		data = data[:length-1]
	}

	// Octet k to l
	if u.Assoni {
		if length < idx+1 {
			return fmt.Errorf("ie UserPlaneIPResourceInformation.NetworkInstance Inadequate TLV length: %d", length)
		}
		err := u.NetworkInstance.UnmarshalBinary(data[idx:])
		if err != nil {
			return err
		}
	}

	return nil
}

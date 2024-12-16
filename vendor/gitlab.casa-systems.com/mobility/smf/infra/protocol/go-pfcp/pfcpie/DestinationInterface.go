package pfcpie

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

const (
	DestinationInterfaceAccessTag uint16 = 42
)

const (
	DestinationInterfaceAccess uint8 = iota
	DestinationInterfaceCore
	DestinationInterfaceSgiLanN6Lan
	DestinationInterfaceCpFunction
	DestinationInterfaceLiFunction
	DestinationInterface5GVnInternal
)

// Destination Interface. Refer to [TS29244-g40 8.2.24 Destination Interface]
//
// The Destination Interface IE type shall be encoded as shown in Figure
// 8.2.24-1. It indicates the type of the interface towards which an outgoing
// packet is sent.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2      1
//	         +--------------------------------------------------------------+
//	1 to 2   |                       Type = 42(decimal)                     |
//	         |--------------------------------------------------------------|
//	3 to 4   |                           Length = n                         |
//	         |-------------------------------+------------------------------|
//	   5     |                               |      Interface Value         |
//	         |-------------------------------+------------------------------|
//	6 to n+4 |  These octets(s) is/are present only if explicity specified  |
//	         +--------------------------------------------------------------+
//	                      Figure 8.2.24-1: Destination Interface
//
// The Interface value shall be encoded as a 4 bits binary integer as specified
// in Table 8.2.24-1.
//
//	                  Table 8.2.24-1: Interface value
//	+----------------------------------------+----------------------------+
//	|          Interface Value               |      Values(Decimal)       |
//	+----------------------------------------+----------------------------+
//	|    Access (NOTE 1, NOTE 3, NOTE 4)     |             0              |
//	+----------------------------------------+----------------------------+
//	|         Core (see NOTE 1): 1           |             1              |
//	+----------------------------------------+----------------------------+
//	|           SGi-LAN/N6-LAN               |             2              |
//	+----------------------------------------+----------------------------+
//	|             CP- Function               |             3              |
//	+----------------------------------------+----------------------------+
//	|       LI Function (see NOTE 2)         |             4              |
//	+----------------------------------------+----------------------------+
//	|              5G VN Internal            |             5              |
//	+----------------------------------------+----------------------------+
//	|                 Spare                  |          6 to 15           |
//	+----------------------------------------+----------------------------+
//	| NOTE 1:                                                             |
//	|     The "Access" and "Core" values denote a downlink and uplink     |
//	|     traffic direction respectively.                                 |
//	| NOTE 2:                                                             |
//	|     LI Function may denote an SX3LIF or an LMISF. See clause 5.7.   |
//	| NOTE 3:                                                             |
//	|     For indirect data forwarding, the Source Interface in the PDR   |
//	|     and the Destination Interface in the FAR shall both be set to   |
//	|     "Access", in the forwarding SGW(s). The Interface value does    |
//	|     not infer any traffic direction, in PDRs and FARs set up for    |
//	|     indirect data forwarding, i.e.  with both the Source and        |
//	|     Destination Interfaces set to Access.                           |
//	| NOTE 4:                                                             |
//	|     For a HTTP redirection, the Source Interface in the PDR to      |
//	|     match the uplink packets to be redirected and the Destination   |
//	|     Interface in the FAR to enable the HTTP redirection shall both  |
//	|     be set to "Access".                                             |
//	|                                                                     |
//	+---------------------------------------------------------------------+
type DestinationInterface struct {
	InterfaceValue uint8 `json:"interfaceValue,omitempty"` // 0x00001111
}

func (d *DestinationInterface) MarshalBinary() (data []byte, err error) {
	// Octet 5
	if bits.Len8(d.InterfaceValue) > 4 {
		return []byte(""), fmt.Errorf("Interface data shall not be greater than 4 bits binary integer")
	}
	data = append([]byte(""), byte(d.InterfaceValue))

	return data, nil
}

func (d *DestinationInterface) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if idx+1 > length {
		return fmt.Errorf("IE DestinationInterface Inadequate TLV length: %d", length)
	}
	d.InterfaceValue = uint8(data[idx]) & Mask4
	//idx = idx + 1

	//if length != idx {
	//return fmt.Errorf("Inadequate TLV length: %d", length)
	//}

	return nil
}

func (d *DestinationInterface) TlvMarshal() (data []byte, err error) {
	var tag uint16 = 42
	idx := uint16(0)
	// Encode tag.
	data = make([]byte, 8)
	binary.BigEndian.PutUint16(data[idx:], tag)
	idx += 2
	buf, err := d.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("DestinationInterface call MarshalBinary failed: %v", err)
	}
	// Encode length
	lengthOfDI := uint16(len(buf))
	binary.BigEndian.PutUint16(data[idx:], uint16(lengthOfDI))
	idx += 2
	// Encode Value
	data = append(data[:idx], buf...)
	idx += lengthOfDI

	return data, nil
}

func (d *DestinationInterface) TlvUnmarshal(data []byte) (err error) {
	length := uint16(len(data))
	idx := uint16(0)
	// Decode tag.
	tag := binary.BigEndian.Uint16(data[idx : idx+2])
	if tag != uint16(DestinationInterfaceAccessTag) {
		return fmt.Errorf("DestinationInterface TlvUnmarshal failed: tag %d is not %d", tag, DestinationInterfaceAccessTag)
	}
	idx += 2
	if idx > length {
		return fmt.Errorf("DestinationInterface TlvUnmarshal failed: idx(%d) > length(%d)", idx, length)

	}
	lengthOfDI := binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2
	if idx > length {
		return fmt.Errorf("DestinationInterface TlvUnmarshalBinary failed: length(%d)"+
			" > length(%d)", idx, length)
	}
	err = d.UnmarshalBinary(data[idx:])
	idx += lengthOfDI
	if err != nil {
		return fmt.Errorf("DestinationInterface call UnmarshalBinary failed: %v", err)
	}
	return nil
}

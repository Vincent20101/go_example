package pfcpie

import (
	"encoding/binary"
	"fmt"
	"net"
)

// Remote GTP-U Peer IE. Refer to [TS29244-g40 8.2.70 Remote GTP-U Peer].
//
// The Remote GTP-U Peer IE shall be encoded as depicted in Figure 8.2.70-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2      1
//	         +--------------------------------------------------------------+
//	1 to 2   |                       Type = 103(decimal)                    |
//	         |--------------------------------------------------------------|
//	3 to 4   |                           Length = n                         |
//	         |-------------------------------+------------------------------|
//	  5      |             Spare             |   NI  |   DI  |  V4   |  V6  |
//	         |-------------------------------+------------------------------|
//
// m to m+3  |                         IPv4 addresse                        |
//
//	|--------------------------------------------------------------|
//
// p to p+15 |                         IPv6 addresse                        |
//
//	+--------------------------------------------------------------+
//
// q to q+1  |              Length of Destination Interface Field           |
//
//	+--------------------------------------------------------------+
//
// q+2 to r  |                    Destination Interface Field               |
//
//	+--------------------------------------------------------------+
//
// s to s+1  |                 Length of Network Instance field             |
//
//	+--------------------------------------------------------------+
//
// s+2 to t  |                    Network Instance field                    |
//
//	+--------------------------------------------------------------+
//
// k to n+4  |   These octet(s) is/are present only if explicitly specifie  |
//
//	+--------------------------------------------------------------+
//
// Figure 8.2.70-1: Remote GTP-U Peer
//
// The following flags are coded within Octet 5:
//
//   - Bit 1 – V6: If this bit is set to "1", then the IPv6 address field shall
//     be present, otherwise the IPv6 address field shall not be present.
//
//   - Bit 2 – V4: If this bit is set to "1", then the IPv4 address field shall
//     be present, otherwise the IPv4 address field shall not be present.
//
//   - Bit 3 – DI: If this bit is set to "1", then the Length of Destination
//     Interface field and the Destination Interface field shall be present,
//     otherwise both the Length of Destination Interface field and the
//     Destination Interface field shall not be present.
//
//   - Bit 4 – NI: If this bit is set to "1", then the Length of Network
//     Instance field and the Network Instance field shall be present, otherwise
//     both the Length of Network Instance field and the Network Instance field
//     shall not be present.
//
// - Bit 5 to 8 - Spare, for future use and set to "0".
//
// Either the V4 or the V6 bit shall be set to "1".
//
// Octets "m to (m+3)" and/or "p to (p+15)" (IPv4 address / IPv6 address
// fields), if present, shall contain the respective address values.
//
// The Destination Interface field, when present, shall be encoded as the same
// as Destination Interface IE (from octet 5, i.e. excluding Type and Length)
// specified in clause 8.2.24.
//
// The Network Instance field, when present, shall be encoded as the same as
// Network Instance IE (from octet 5, i.e. excluding Type and Length) specified
// in clause 8.2.4.
type RemoteGTPUPeer struct {
	NI                   bool                  `json:"nI,omitempty"`
	DI                   bool                  `json:"dI,omitempty"`
	V4                   bool                  `json:"v4,omitempty"`
	V6                   bool                  `json:"v6,omitempty"`
	Ipv4Address          net.IP                `json:"ipv4Address,omitempty"`
	Ipv6Address          net.IP                `json:"ipv6Address,omitempty"`
	DestinationInterface *DestinationInterface `json:"destinationInterface,omitempty"` // `tlv:"42"`
	NetworkInstance      *NetworkInstance      `json:"networkInstance,omitempty"`      // `tlv:"22"`
}

func (r *RemoteGTPUPeer) MarshalBinary() (data []byte, err error) {
	errorMsgPrefix := "Marshal RemoteGTPUPeer failed: %s"
	// Octet 5
	tmpUint8 := btou(r.NI)<<3 |
		btou(r.DI)<<2 |
		btou(r.V4)<<1 |
		btou(r.V6)

	data = append([]byte(""), byte(tmpUint8))

	// Ipv4Address, Octet m to (m+3).
	if r.V4 {
		if len(r.Ipv4Address) == 0 || r.Ipv4Address.IsUnspecified() {
			return nil, fmt.Errorf(errorMsgPrefix, "IPv4 address shall be present if V4 is set")
		}
		data = append(data, r.Ipv4Address.To4()...)
	}

	// Ipv6Address, Octet p to (p+15).
	if r.V6 {
		if len(r.Ipv6Address) == 0 || r.Ipv6Address.IsUnspecified() {
			return nil, fmt.Errorf(errorMsgPrefix, "IPv6 address shall be present if V6 is set")
		}
		data = append(data, r.Ipv6Address.To16()...)
	}

	if r.V4 == false && r.V6 == false {
		return nil, fmt.Errorf(errorMsgPrefix, "Either V4 and V6 shall be set")
	}

	// Length of DestinationInterface and DestinationInterface
	if r.DI && r.DestinationInterface != nil {
		dIBuffer, err := r.DestinationInterface.TlvMarshal()
		if err != nil {
			return nil, err
		}
		lengthOfDI := uint16(len(dIBuffer))
		tmp := make([]byte, 2)
		binary.BigEndian.PutUint16(tmp, lengthOfDI)
		data = append(data, tmp...)
		data = append(data, dIBuffer...)
	}

	// Length of NetworkInstance and NetworkInstance
	if r.NI && r.NetworkInstance != nil {
		nIBuffer, err := r.NetworkInstance.TlvMarshal()
		if err != nil {
			return nil, fmt.Errorf(errorMsgPrefix, "NetworkInstance call TlvMarshal failed: %v", err)
		}
		// Length of NetworkInstance
		lengthOfNI := uint16(len(nIBuffer))
		tmp := make([]byte, 2)
		binary.BigEndian.PutUint16(tmp, lengthOfNI)
		// Encode NetworkInstance
		data = append(data, tmp...)
		data = append(data, nIBuffer...)
	}

	return data, nil
}

func (r *RemoteGTPUPeer) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE RemoteGTPUPeer Inadequate TLV length, empty body")
	}
	r.NI = utob(uint8(data[idx]) & BitMask4)
	r.DI = utob(uint8(data[idx]) & BitMask3)
	r.V4 = utob(uint8(data[idx]) & BitMask2)
	r.V6 = utob(uint8(data[idx]) & BitMask1)

	idx = idx + 1

	if length < idx {
		return fmt.Errorf("IE RemoteGTPUPeer Field Ipv4Address Inadequate TLV length: %d", length)
	}
	// Octet m to (m+3)
	if r.V4 {
		r.Ipv4Address = net.IP(data[idx : idx+net.IPv4len]).To4()
		idx = idx + net.IPv4len
		if length < idx {
			return fmt.Errorf("IE RemoteGTPUPeer Field Ipv4Address Inadequate TLV length: %d", length)
		}
	}

	// Octet p to (p+15)
	if r.V6 {
		r.Ipv6Address = net.IP(data[idx : idx+net.IPv6len])
		idx = idx + net.IPv6len
		if length < idx {
			return fmt.Errorf("IE RemoteGTPUPeer Field Ipv6Address Inadequate TLV length: %d", length)
		}
	}

	if r.V4 == false && r.V6 == false {
		//spew.Dump(data)
		return fmt.Errorf("IE RemoteGTPUPeer neither of V4 or V6 is set: %x", data[idx-1])
	}

	if idx > length {
		return fmt.Errorf("RemoteGTPUPeer unmarshal failed: idx is %d, greater than length %d", idx, length)
	}

	if r.DI {
		lengthOfDI := binary.BigEndian.Uint16(data[idx : idx+2])
		idx = idx + 2
		if idx > length {
			return fmt.Errorf("RemoteGTPUPeer unmarshal failed: idx is %d, greater than length %d", idx, length)
		}
		r.DestinationInterface = &DestinationInterface{}
		err := r.DestinationInterface.TlvUnmarshal(data[idx:])
		if err != nil {
			return fmt.Errorf("RemoteGTPUPeer unmarshal failed: %v", err)
		}
		idx = idx + lengthOfDI
		if idx > length {
			return fmt.Errorf("RemoteGTPUPeer unmarshal failed: idx is %d, greater than length %d", idx, length)
		}
	}

	if r.NI {
		lengthOfNI := binary.BigEndian.Uint16(data[idx : idx+2])
		idx = idx + 2
		r.NetworkInstance = new(NetworkInstance)
		if err := r.NetworkInstance.TlvUnmarshal(data[idx : idx+lengthOfNI]); err != nil {
			return err
		}
		idx = idx + lengthOfNI
	}

	return nil
}

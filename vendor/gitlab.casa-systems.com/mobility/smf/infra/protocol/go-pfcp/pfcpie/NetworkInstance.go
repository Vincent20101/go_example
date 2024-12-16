package pfcpie

import (
	"encoding/binary"
	"fmt"
)

var (
	NETWORKINSTANCE_TAG = 22
)

// Network Instance IE. Refer to [TS29244-g40 8.2.4 Network Instance].
//
// The Network Instance IE type shall be encoded as shown in Figure 8.2.4-1. It
// indicates a Network instance.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2      1
//	         +--------------------------------------------------------------+
//	1 to 2   |                       Type = 22(decimal)                     |
//	         |--------------------------------------------------------------|
//	3 to 4   |                           Length = n                         |
//	         |--------------------------------------------------------------|
//
// 5 to n+4  |                        Network Instance                      |
//
//	+--------------------------------------------------------------+
//	                   Figure 8.2.4-1: Network Instance
//
// The Network instance field shall be encoded as an OctetString and shall
// contain an identifier which uniquely identifies a particular Network
// instance (e.g. PDN instance) in the UP function. It may be encoded as a
// Domain Name or an Access Point Name (APN) as per clause 9.1 of
// 3GPP TS 23.003 [2]. In the latter case, the PDN Instance field may contain
// the APN Network Identifier only or the full APN with both the APN Network
// Identifier and the APN Operator Identifier as specified in
// 3GPP TS 23.003 [2] clauses 9.1.1 and 9.1.2.
//
// NOTE:	The APN field is not encoded as a dotted string as commonly used in
// documentation.
type NetworkInstance struct {
	Value string `json:"value,omitempty"`
}

func (n *NetworkInstance) MarshalBinary() (data []byte, err error) {
	return []byte(n.Value), nil
}

func (n *NetworkInstance) UnmarshalBinary(data []byte) error {
	n.Value = string(data)
	return nil
}

// Marshal NetworkInstance with tag and length
func (n *NetworkInstance) TlvMarshal() (data []byte, err error) {
	var tag uint16 = 22
	idx := 0
	// Encode tag.
	data = make([]byte, 8)
	binary.BigEndian.PutUint16(data[idx:], tag)
	idx += 2
	buf, err := n.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("NetworkInstance call MarshalBinary failed: %v", err)
	}
	// Encode length
	lengthOfNI := len(buf)
	binary.BigEndian.PutUint16(data[idx:], uint16(lengthOfNI))
	idx += 2
	// Encode Value
	data = append(data[:idx], buf...)
	idx += lengthOfNI
	return data, nil
}

// Unmarshal NetworkInstance with tag and length
func (n *NetworkInstance) TlvUnmarshal(data []byte) (err error) {
	idx := 0
	// Decode tag.
	tag := binary.BigEndian.Uint16(data[idx : idx+2])
	if tag != uint16(NETWORKINSTANCE_TAG) {
		return fmt.Errorf("NetworkInstance TlvUnmarshal failed: tag %d is not %d", tag, NETWORKINSTANCE_TAG)
	}
	idx += 2
	length := binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2
	if int(length)+idx > len(data) {
		return fmt.Errorf("NetworkInstance TlvUnmarshalBinary failed: length+idx is"+
			"%d greater than %d", int(length)+idx, len(data))
	}
	err = n.UnmarshalBinary(data[idx:])
	if err != nil {
		return fmt.Errorf("NetworkInstance call MarshalBinary failed: %v", err)
	}
	return nil
}

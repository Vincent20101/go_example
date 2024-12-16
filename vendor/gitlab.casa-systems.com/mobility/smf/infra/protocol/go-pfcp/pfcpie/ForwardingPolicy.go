package pfcpie

import (
	"fmt"
)

// Forwarding Policy IE. [Refer to TS29244 8.2.23 Forwarding Policy]
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                       Type = 41(decimal)                      |
//	         |---------------------------------------------------------------|
//	3 to 4   |                           Length = n                          |
//	         |---------------------------------------------------------------|
//	  5      |              Forwarding Policy Identifier Length              |
//	         |---------------------------------------------------------------|
//	j to k   |                  Forwarding Policy Indentifier                |
//	         |---------------------------------------------------------------|
//	m to n+4 |   These octets(s) is/are present only if explicity specified  |
//	         +---------------------------------------------------------------+
//
// The Forwarding Policy Identifier Length shall indicate the length of the
// Forwarding Policy Identifier.
//
// The Forwarding Policy Identifier shall be encoded as an Octet String
// containing a reference to a pre-configured Forwarding Policy in the UP
// function, with a maximum length of 255 octets.
type ForwardingPolicy struct {
	ForwardingPolicyIdentifier []byte `json:"forwardingPolicyIdentifier,omitempty"`
}

func (f *ForwardingPolicy) MarshalBinary() (data []byte, err error) {
	// Check length of ForwardingPolicyIdentifier
	lengthOfFpi := len(f.ForwardingPolicyIdentifier)
	if lengthOfFpi > 255 {
		return nil, fmt.Errorf("The maximum length of Forwading Policy Identifier is 255, and current length is %d", lengthOfFpi)
	}
	// Octet 5
	data = append([]byte(""), byte(uint8(lengthOfFpi)))
	// Octet j to k
	data = append(data, f.ForwardingPolicyIdentifier...)
	return data, nil
}

func (f *ForwardingPolicy) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE ForwardingPolicy Inadequate TLV length: %d", length)
	}
	lengthOfFpi := uint8(data[idx])
	idx = idx + 1

	// Octet 6 to (6+a)
	if length < idx+uint16(lengthOfFpi) {
		return fmt.Errorf("IE ForwardingPolicy Field ForwardingPolicyIdentifier Inadequate TLV length: %d", length)
	}
	f.ForwardingPolicyIdentifier = data[idx : idx+uint16(lengthOfFpi)]
	idx = idx + uint16(lengthOfFpi)

	return nil
}

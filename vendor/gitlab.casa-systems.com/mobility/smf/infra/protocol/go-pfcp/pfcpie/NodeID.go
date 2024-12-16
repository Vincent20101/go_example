package pfcpie

import (
	"encoding/json"
	"fmt"
	"math/bits"
	"net"
	"strings"
)

const (
	NodeIdTypeIpv4Address uint8 = iota
	NodeIdTypeIpv6Address
	NodeIdTypeFqdn
)

// NodeID IE, Refer to [TS29244 8.2.38 Node ID]
//
// The Node ID IE shall contain an FQDN or an IPv4/IPv6 address. It shall be
// encoded as shown in Figure 8.2.38-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                       Type = 60(decimal)                      |
//	         |---------------------------------------------------------------|
//	3 to 4   |                           Length = n                          |
//	         |-------------------------------+-------------------------------|
//	  5      |             Spare             |          Node ID Type         |
//	         |-------------------------------+-------------------------------|
//	6 to o   |                         Node ID Value                         |
//	         |---------------------------------------------------------------|
//	m to n+4 |   These octets(s) is/are present only if explicity specified  |
//	         +---------------------------------------------------------------+
//
// Node ID Type indicates the type of the Node ID value. It shall be encoded as
// a 4 bits binary integer as specified in Table 8.2.38-2.
//
//	           Table 8.2.38-2: Node ID Type
//	+-------------------+-------------------------------+
//	|   Node ID Type    |        Node ID Type           |
//	|  Value (Decimal)  |                               |
//	|-------------------+-------------------------------|
//	|        0          |        Ipv4 address           |
//	|-------------------+-------------------------------|
//	|        1          |        Ipv6 address           |
//	|-------------------+-------------------------------|
//	|        2          |           FQDN                |
//	|-------------------+-------------------------------|
//	|     3 to 15       |     Spare, for future use.    |
//	+-------------------+-------------------------------+
//
// If the Node ID is an IPv4 address, the Node ID value length shall be 4
// Octet.
//
// If the Node ID is an IPv6 address, the Node ID value length shall be 16
// Octet.
//
// If the Node ID is an FQDN, the Node ID value encoding shall be identical to
// the encoding of a FQDN within a DNS message of clause 3.1 of IETF RFC 1035
// [27] but excluding the trailing zero byte.
//
//	NOTE 1: The FQDN field in the IE is not encoded as a dotted string as
//	commonly used in DNS master zone files.
//
// ============== IETF RFC 1035 3.1. Name space definitions ===============
// Full specs ses: https://tools.ietf.org/html/rfc1035#section-3.1
//
// 3.1. Name space definitions
//
// Domain names in messages are expressed in terms of a sequence of labels.
// Each label is represented as a one octet length field followed by that
// number of octets.  Since every domain name ends with the null label of the
// root, a domain name is terminated by a length byte of zero.  The high order
// two bits of every length octet must be zero, and the remaining six bits of
// the length field limit the label to 63 octets or less.
//
// To simplify implementations, the total length of a domain name (i.e.,
// label octets and label length octets) is restricted to 255 octets or
// less.
//
// Although labels can contain any 8 bit values in octets that make up a
// label, it is strongly recommended that labels follow the preferred
// syntax described elsewhere in this memo, which is compatible with
// existing host naming conventions.  Name servers and resolvers must
// compare labels in a case-insensitive manner (i.e., A=a), assuming ASCII
// with zero parity.  Non-alphabetic codes must match exactly.
type NodeID struct {
	NodeIdType  uint8  `json:"nodeIdType,omitempty"` // 0x00001111
	NodeIdValue []byte `json:"nodeIdValue,omitempty"`
}

type nodeIdRaw struct {
	NodeIdType  uint8  `json:"nodeIdType,omitempty"`
	NodeIdValue string `json:"nodeIdValue,omitempty"`
}

// getNodeID return a nodeIdRaw struct after convert NodeIdValue from net.IP to
// readable string. This function is only used for NodeID.Unmarshal.
func (n *NodeID) getNodeID() *nodeIdRaw {
	nodeIdType := n.NodeIdType
	nodeIdValue := n.NodeIdValue
	node := &nodeIdRaw{}
	node.NodeIdType = n.NodeIdType
	if nodeIdType == NodeIdTypeIpv4Address || nodeIdType == NodeIdTypeIpv6Address {
		node.NodeIdValue = net.IP(nodeIdValue).String()
		return node
	}
	return node
}

// SetNodeID is used to set NodeID, it converts the nodeIdValue string to net.IP
// if nodeIdType is 0 or 1, convert the nodeId Value string to []byte if
// nodeIdType is 2.  The nodeIdValue will be such a string:
//
//   - if nodeIdtype is 0 (ipv4): an ipv4 string, i.e.
//     "192.168.1.1"
//
//   - if nodeIdtype is 1 (ipv6): an ipv6 string, i.e.
//     "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
//
//   - if nodeIdtype is 2 (fqdn): an fqdn string, i.e.
//     "example.com"
//
// The `nodeIdType` should be exactly matched with nodeIdValue.
func (n *NodeID) SetNodeID(nodeIdType uint8, nodeIdValue string) error {
	if nodeIdType > NodeIdTypeFqdn {
		return fmt.Errorf("SetNodeID failed: IE NodeID type error, unknown type: %d", nodeIdType)
	}

	// FQDN
	if nodeIdType == NodeIdTypeFqdn {
		n.NodeIdType = nodeIdType
		n.NodeIdValue = resolveFqdn(nodeIdValue)
		return nil
	}

	// Ipv4 or Ipv6
	n.NodeIdType = nodeIdType
	n.NodeIdValue = net.ParseIP(nodeIdValue)
	if n.NodeIdValue == nil || len(n.NodeIdValue) == 0 {
		return fmt.Errorf("IE NodeID(ipv4 type) Parse IP failed while unmashalling NodeIdValue: %s", string(nodeIdValue))
	}

	if nodeIdType == NodeIdTypeIpv4Address {
		n.NodeIdValue = net.IP(n.NodeIdValue).To4()
		if n.NodeIdValue == nil || len(n.NodeIdValue) == 0 {
			return fmt.Errorf("IE NodeID(ipv4 type) Ipv4 %s convert to 4-byte failed. ", string(nodeIdValue))
		}
	}

	if nodeIdType == NodeIdTypeIpv6Address {
		n.NodeIdValue = net.IP(n.NodeIdValue).To16()
		if n.NodeIdValue == nil || len(n.NodeIdValue) == 0 {
			return fmt.Errorf("IE NodeID(ipv6 type) Ipv6 %s convert to 16-byte failed. ", string(nodeIdValue))
		}
	}

	return nil
}

func (n *NodeID) MarshalBinary() (data []byte, err error) {
	// Octet 5
	if bits.Len8(n.NodeIdType) > 4 {
		return []byte(""), fmt.Errorf("Node ID type shall not be greater than 4 bits binary integer")
	}
	data = append([]byte(""), byte(n.NodeIdType))

	// Octet 6 to o
	if n.NodeIdType == 0 && len(n.NodeIdValue) != net.IPv4len {
		return []byte(""), fmt.Errorf("Length of node ID data shall be 4 Octet if node ID is an IPv4 address")
	}
	if n.NodeIdType == 1 && len(n.NodeIdValue) != net.IPv6len {
		return []byte(""), fmt.Errorf("Length of node ID data shall be 16 Octet if node ID is an IPv6 address")
	}
	if n.NodeIdType == 2 && len(n.NodeIdValue) > 256 {
		return []byte(""), fmt.Errorf("Length of node ID data shall be less than 256 Octets if node ID is an FQDN.")
	}
	data = append(data, n.NodeIdValue...)

	return data, nil
}

func resolveFqdn(str string) []byte {
	return []byte(str)
}

func (n *NodeID) UnmarshalJSON(b []byte) error {
	// If b is config as ipv4 or ipv6 address.
	node := &nodeIdRaw{}
	json.Unmarshal(b, node)
	nodeIdVal := strings.Trim(string(node.NodeIdValue), "\"")
	err := n.SetNodeID(node.NodeIdType, nodeIdVal)
	if err != nil {
		return err
	}
	return nil
}

func (n *NodeID) MarshalJSON() ([]byte, error) {
	// If b is config as ipv4 or ipv6 address.
	node := n.getNodeID()
	return json.Marshal(node)
}

func (n *NodeID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Unmarshal NodeID failed: empty body")
	}
	n.NodeIdType = uint8(data[idx]) & Mask4
	idx = idx + 1

	// Octet 6 to o
	if n.NodeIdType == 0 {
		if length < idx+net.IPv4len {
			return fmt.Errorf("Unmarshal NodeID(nodeidtype:%d)->[NodeId Value] failed, inadequate length: %d", n.NodeIdType, length)
		}
		n.NodeIdValue = data[idx : idx+net.IPv4len]
		idx = idx + net.IPv4len
	} else if n.NodeIdType == 1 {
		if length < idx+net.IPv6len {
			return fmt.Errorf("Unmarshal NodeID(nodeidtype:%d)->[NodeId Value] failed, inadequate length: %d", n.NodeIdType, length)
		}
		n.NodeIdValue = data[idx : idx+net.IPv6len]
		idx = idx + net.IPv6len
	} else {
		n.NodeIdValue = data[idx:]
		idx = idx + uint16(len(n.NodeIdValue))
	}

	return nil
}

func (n *NodeID) ResolveNodeIdToIp() (net.IP, net.IP) {
	switch n.NodeIdType {
	case NodeIdTypeIpv4Address:
		return net.IP(n.NodeIdValue).To4(), nil
	case NodeIdTypeIpv6Address:
		return nil, net.IP(n.NodeIdValue).To16()
	case NodeIdTypeFqdn:
		ns, _ := net.LookupHost(string(n.NodeIdValue))
		var ipv4, ipv6 net.IP
		for _, ip := range ns {
			if net.ParseIP(ip).To4() != nil {
				ipv4 = net.ParseIP(ip).To4()
			}
			if net.ParseIP(ip).To16() != nil {
				ipv6 = net.ParseIP(ip).To16()
			}
			if ipv4 != nil && ipv6 != nil {
				break
			}
		}
		return ipv4, ipv6
	default:
		return nil, nil
	}
}

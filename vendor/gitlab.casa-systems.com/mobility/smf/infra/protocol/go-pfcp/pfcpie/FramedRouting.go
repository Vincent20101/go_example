package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// [TS29244 8.2.109 Framed-Routing]
// The Framed-Routing field shall be encoded as the value part of the
// Framed-Routing AVP specified in IETE RFC 2856.
//
//	Octets
//	          +-------------------------------------------------------------+
//	1 to 2    |                   Type = 154(decimal)                       |
//	          |-------------------------------------------------------------|
//	3 to 4    |                      Length = 4                             |
//	          |-------------------------------------------------------------|
//	5 to 8    |                     Framed-Routing                          |
//	          +-------------------------------------------------------------+
//
// The definition of Framed-Routing could be found in
// gitlab.casa-systems.com/mobility/smf/infra/protocol/radius/rfc2865, which is uint32 type.
type FramedRouting struct {
	FramedRoutingdata uint32 `json:"framedRoutingdata,omitempty"`
}

func (f *FramedRouting) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, f.FramedRoutingdata)
	return data, nil
}

func (f *FramedRouting) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 6
	if length < idx+4 {
		return fmt.Errorf("ie FramedRouting Inadequate TLV length: %d", length)
	}
	f.FramedRoutingdata = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

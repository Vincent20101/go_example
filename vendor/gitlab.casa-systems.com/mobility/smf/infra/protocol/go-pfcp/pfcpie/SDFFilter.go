package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// SDF Filter IE. Refer to [TS29244-g40 8.2.5 SDF Filter]
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                       Type = 23(decimal)                      |
//	         |---------------------------------------------------------------|
//	3 to 4   |                          Length = n                           |
//	         |-----------------------+-------+-------+-------+-------+-------|
//	  5      |        Spare          |  BID  |   FL  |  SPI  |  TTC  |  FD   |
//	         |-----------------------+-------+-------+-------+-------+-------|
//	  6      |                             Spare                             |
//	         |---------------------------------------------------------------|
//	m to m+1 |                  Length of Flow Description                   |
//	         |---------------------------------------------------------------|
//	m+2 to p |                      Flow Description                         |
//	         |---------------------------------------------------------------|
//	s to s+1 |                      Tos Traffic Class                        |
//	         |---------------------------------------------------------------|
//	t to t+3 |                  Security Parameter Index                     |
//	         |---------------------------------------------------------------|
//	v to v+2 |                           Flow Label                          |
//	         |---------------------------------------------------------------|
//	w to w+3 |                          SDF Filter ID                        |
//	         |---------------------------------------------------------------|
//	x to n+4 |   These octets(s) is/are present only if explicity specified  |
//	         +---------------------------------------------------------------+
//	                                Figure 8.2.5-1: SDF Filte
//
// The following flags are coded within Octet 5:
//
//   - Bit 1  FD (Flow Description): If this bit is set to "1", then the Length
//     of Flow Description and the Flow Description fields shall be present,
//     otherwise they shall not be present.
//
//   - Bit 2  TTC (ToS Traffic Class): If this bit is set to "1", then the ToS
//     Traffic Class field shall be present, otherwise the ToS Traffic Class field
//     shall not be present.
//
//   - Bit 3  SPI (Security Parameter Index): If this bit is set to "1", then
//     the Security Parameter Index field shall be present, otherwise the Security
//     Parameter Index field shall not be present.
//
//   - Bit 4  FL (Flow Label): If this bit is set to "1", then the Flow Label
//     field shall be present, otherwise the Flow Label field shall not be
//     present.
//
//   - Bit 5  BID (Bidirectional SDF Filter): If this bit is set to "1", then
//     the SDF Filter ID shall be present, otherwise the SDF Filter ID shall not
//     be present.
//
//   - Bit 6 to 8: Spare, for future use and set to "0".
//
// The Flow Description field, when present, shall be encoded as an OctetString
// as specified in clause 5.4.2 of 3GPP TS 29.212 [8].
//
// The ToS Traffic Class field, when present, shall be encoded as an
// OctetString on two octets as specified in clause 5.3.15 of 3GPP TS 29.212
// [8].
//
// The Security Parameter Index field, when present, shall be encoded as an
// OctetString on four octets and shall contain the IPsec security parameter
// index (which is a 32-bit field), as specified in clause 5.3.51 of 3GPP TS
// 29.212 [8].
//
// The Flow Label field, when present, shall be encoded as an OctetString on 3
// octets as specified in clause 5.3.52 of 3GPP TS 29.212 [8] and shall contain
// an IPv6 flow label (which is a 20-bit field). The bits 8 to 5 of the octet
// "v" shall be spare and set to zero, and the remaining 20 bits shall contain
// the IPv6 flow label.
//
// ...
//
// The SDF Filter ID, when present, shall be encoded as an Unsigned32 binary integer value. It shall uniquely identify an
//
// SDF Filter among all the SDF Filters provisioned for a given PFCP Session.
// The source/destination IP address and port information, in a bidirectional
// SDF Filter, shall be set as for downlink IP flows. The SDF filter for the
// opposite direction has the same parameters, but having the source and
// destination address/port parameters swapped. When being provisioned with a
// bidirectional SDF filter in a PDR, the UP function shall apply the SDF
// filter as specified in clause 5.2.1A.2A.
type SDFFilter struct {
	Bid                    bool   `json:"bid,omitempty"`
	Fl                     bool   `json:"fl,omitempty"`
	Spi                    bool   `json:"spi,omitempty"`
	Ttc                    bool   `json:"ttc,omitempty"`
	Fd                     bool   `json:"fd,omitempty"`
	FlowDescription        []byte `json:"flowDescription,omitempty"`
	TosTrafficClass        []byte `json:"tosTrafficClass,omitempty"`
	SecurityParameterIndex []byte `json:"securityParameterIndex,omitempty"`
	FlowLabel              []byte `json:"flowLabel,omitempty"`
	SdfFilterId            uint32 `json:"sdfFilterId,omitempty"`
}

func (s *SDFFilter) SetFlowDescription(fd string) {
	s.FlowDescription = []byte(fd)
}

func (s *SDFFilter) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5
	tmpUint8 := btou(s.Bid)<<4 |
		btou(s.Fl)<<3 |
		btou(s.Spi)<<2 |
		btou(s.Ttc)<<1 |
		btou(s.Fd)
	data = append([]byte(""), byte(tmpUint8))
	idx = idx + 1

	// Octet 6 (spare)
	data = append(data, byte(0))
	idx = idx + 1

	lengthOfFlowDescription := uint16(len(s.FlowDescription))
	// Octet m to (m+1)
	// Octet (m+2) to p
	if s.Fd {
		if lengthOfFlowDescription == 0 {
			return []byte(""), fmt.Errorf("Length of flow description shall be present if FD is set")
		}
		data = append(data, make([]byte, 2)...)
		binary.BigEndian.PutUint16(data[idx:], lengthOfFlowDescription)
		idx = idx + 2

		if len(s.FlowDescription) != int(lengthOfFlowDescription) {
			return []byte(""), fmt.Errorf("Unmatch length of flow description: Expect %d, got %d", lengthOfFlowDescription, len(s.FlowDescription))
		}
		data = append(data, s.FlowDescription...)
		idx = idx + uint16(len(s.FlowDescription))
	}

	// Octet s to (s+1)
	if s.Ttc {
		if len(s.TosTrafficClass) != 2 {
			return []byte(""), fmt.Errorf("ToS traffic class shall be exactly two bytes")
		}
		data = append(data, s.TosTrafficClass...)
		idx = idx + 2
	}

	// Octet t to (t+3)
	if s.Spi {
		if len(s.SecurityParameterIndex) != 4 {
			return []byte(""), fmt.Errorf("Security parameter index shall be exactly four bytes")
		}
		data = append(data, s.SecurityParameterIndex...)
		idx = idx + 4
	}

	// Octet v to (v+2)
	if s.Fl {
		if len(s.FlowLabel) != 3 {
			return []byte(""), fmt.Errorf("Flow label shall be exactly three bytes")
		}
		data = append(data, s.FlowLabel...)
		idx = idx + 3
	}

	// Octet w to (w+3)
	if s.Bid {
		data = append(data, make([]byte, 4)...)
		binary.BigEndian.PutUint32(data[idx:], s.SdfFilterId)
	}

	return data, nil
}

func (s *SDFFilter) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie SDFFilter Octet5 Inadequate TLV length: %d", length)
	}
	s.Bid = utob(uint8(data[idx]) & BitMask5)
	s.Fl = utob(uint8(data[idx]) & BitMask4)
	s.Spi = utob(uint8(data[idx]) & BitMask3)
	s.Ttc = utob(uint8(data[idx]) & BitMask2)
	s.Fd = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet 6 (spare)
	if length < idx+1 {
		return fmt.Errorf("ie SDFFilter octet6[Spare] Inadequate TLV length: %d", length)
	}
	idx = idx + 1

	// Octet m to (m+1)
	// Octet (m+2) to p
	if s.Fd {
		if length < idx+2 {
			return fmt.Errorf("ie SDFFilter.Fd field[Flow Description] Inadequate TLV length: %d", length)
		}
		lengthOfFlowDescription := binary.BigEndian.Uint16(data[idx:])
		idx = idx + 2

		if length < idx+lengthOfFlowDescription {
			return fmt.Errorf("ie SDFFilter.FlowDescription Inadequate TLV length: %d", length)
		}
		s.FlowDescription = data[idx : idx+lengthOfFlowDescription]
		idx = idx + lengthOfFlowDescription
	}

	// Octet s to (s+1)
	if s.Ttc {
		if length < idx+2 {
			return fmt.Errorf("ie SDFFilter.Ttc Inadequate TLV length: %d", length)
		}
		s.TosTrafficClass = data[idx : idx+2]
		idx = idx + 2
	}

	// Octet t to (t+3)
	if s.Spi {
		if length < idx+4 {
			return fmt.Errorf("ie SDFFilter.Spi Inadequate TLV length: %d", length)
		}
		s.SecurityParameterIndex = data[idx : idx+4]
		idx = idx + 4
	}

	// Octet v to (v+2)
	if s.Fl {
		if length < idx+3 {
			return fmt.Errorf("ie SDFFilter.Fl Inadequate TLV length: %d", length)
		}
		s.FlowLabel = data[idx : idx+3]
		idx = idx + 3
	}

	// Octet w to (w+3)
	if s.Bid {
		if length < idx+4 {
			return fmt.Errorf("ie SDFFilter.Bid Inadequate TLV length: %d", length)
		}
		s.SdfFilterId = binary.BigEndian.Uint32(data[idx : idx+4])
		idx = idx + 4
	}

	return nil
}

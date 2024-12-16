package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// 8.2.169	Packet Delay Thresholds
// The Packet Delay Thresholds IE contains a packet delay thresholds in milliseconds to be monitored by the UP function. It shall be encoded as shown in Figure 8.2.169-1.
//
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                    Type = 245(decimal)                          |
//           |----------------------------------------------------------------|
//  3 to 4   |                      Length = n                                |
//           |----------------------------------------------------------------|
//    5      |             Spare                      |  RP  |   UL  |   DL   |
//           |----------------------------------------------------------------|
// m to m+3  |                 Downlink packet delay threshold                |
//           |----------------------------------------------------------------|
// p to p+3  |                 Uplink packet delay threshold                  |
//           |----------------------------------------------------------------|
// q to q+3  |                 Round trip packet delay threshold              |
//           |----------------------------------------------------------------|
// s to (n+4)|   These octets(s) is/are present only if explicity specified   |
//           +----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
// -	Bit 1  DL: If this bit is set to "1", then the Downlink packet delay threshold field shall be present, otherwise the Downlink packet delay threshold field shall not be present.
// -	Bit 2  UL: If this bit is set to "1", then the Uplink packet delay threshold field shall be present, otherwise the Uplink packet delay threshold field shall not be present.
// -	Bit 3  RP: If this bit is set to "1", then the Round trip packet delay threshold field shall be present, otherwise the Round trip packet delay threshold field shall not be present.
// -	Bit 4 to 8: Spare, for future use and set to "0".
//  At least one bit shall be set to "1". Several bits may be set to "1".
//The Downlink packet delay threshold, Uplink packet delay threshold and Round trip packet delay threshold fields shall be encoded as an Unsigned32 binary integer value. They shall contain the downlink, uplink or round trip packet delay in milliseconds respectively.

type PacketDelayThreshold struct {
	Rp               bool   `json:"rp,omitempty"`
	Ul               bool   `json:"ul,omitempty"`
	Dl               bool   `json:"dl,omitempty"`
	DlPktDelayThresh uint32 `json:"dlPktDelayThresh,omitempty"`
	UlPktDelayThresh uint32 `json:"ulPktDelayThresh,omitempty"`
	RpPktDelayThresh uint32 `json:"rpPktDelayThresh,omitempty"`
}

func (p *PacketDelayThreshold) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5
	tmpUint8 := btou(p.Rp)<<2 |
		btou(p.Ul)<<1 |
		btou(p.Dl)
	data = append([]byte(""), byte(tmpUint8))
	idx = idx + 1

	// Octet m to (m+3)
	if p.Dl {
		data = append(data, make([]byte, 4)...)
		binary.BigEndian.PutUint32(data[idx:], p.DlPktDelayThresh)
		idx = idx + 4
	}

	// Octet p to (p+3)
	if p.Ul {
		data = append(data, make([]byte, 4)...)
		binary.BigEndian.PutUint32(data[idx:], p.UlPktDelayThresh)
		idx = idx + 4
	}

	// Octet q to (q+3)
	if p.Rp {
		data = append(data, make([]byte, 4)...)
		binary.BigEndian.PutUint32(data[idx:], p.RpPktDelayThresh)
	}

	return data, nil
}

func (p *PacketDelayThreshold) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE PacketDelayThreshold Inadequate TLV length: %d", length)
	}
	p.Rp = utob(uint8(data[idx]) & BitMask3)
	p.Ul = utob(uint8(data[idx]) & BitMask2)
	p.Dl = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet m to (m+3)
	if p.Dl {
		if length < idx+4 {
			return fmt.Errorf("IE PacketDelayThreshold.DlPktDelayThresh  Inadequate TLV length: %d", length)
		}
		p.DlPktDelayThresh = binary.BigEndian.Uint32(data[idx:])
		idx = idx + 4
	}

	// Octet p to (p+3)
	if p.Ul {
		if length < idx+4 {
			return fmt.Errorf("IE PacketDelayThreshold.UlPktDelayThresh Inadequate TLV length: %d", length)
		}
		p.UlPktDelayThresh = binary.BigEndian.Uint32(data[idx:])
		idx = idx + 4
	}

	// Octet q to (q+3)
	if p.Rp {
		if length < idx+4 {
			return fmt.Errorf("IE PacketDelayThreshold.RpPktDelayThresh Inadequate TLV length: %d", length)
		}
		p.RpPktDelayThresh = binary.BigEndian.Uint32(data[idx:])
		idx = idx + 4
	}

	return nil
}

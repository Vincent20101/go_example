package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// SubsequentVolumeThreshold IE. Refer to [TS29244 8.2.16 Subsequent Volume Threshold]
//
// The Subsequent Volume Threshold IE contains the subsequent traffic volume
// thresholds to be monitored by the UP function after the Monitoring Time. It
// shall be encoded as shown in Figure 8.2.16-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                    Type = 34(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	  5      |             Spare                     | DLVOL | ULVOL | TOVOL  |
//	         |----------------------------------------------------------------|
//
// m to m+7  |                          Total Volume                          |
//
//	|----------------------------------------------------------------|
//
// p to p+7  |                         Uplink Volume                          |
//
//	|----------------------------------------------------------------|
//
// q to q+7  |                        Downlink Volume                         |
//
//	|----------------------------------------------------------------|
//
// 8 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//
//	+----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//
//   - Bit 1 – TOVOL: If this bit is set to "1", then the Total Volume field
//     shall be present, otherwise the Total Volume field shall not be present.
//
//   - Bit 2 – ULVOL: If this bit is set to "1", then the Uplink Volume field
//     shall be present, otherwise the Uplink Volume field shall not be present.
//
//   - Bit 3 – DLVOL: If this bit is set to "1", then the Downlink Volume field
//     shall be present, otherwise the Downlink Volume field shall not be
//     present.
//
//   - Bit 4 to 8: Spare, for future use and set to "0".
//
// At least one bit shall be set to "1". Several bits may be set to "1".
//
// The Total Volume, Uplink Volume and Downlink Volume fields shall be encoded
// as an Unsigned64 binary integer value. They shall contain the total, uplink
// or downlink number of octets respectively.
type SubsequentVolumeThreshold struct {
	Dlvol          bool   `json:"dlvol,omitempty"`
	Ulvol          bool   `json:"ulvol,omitempty"`
	Tovol          bool   `json:"tovol,omitempty"`
	TotalVolume    uint64 `json:"totalVolume,omitempty"`
	UplinkVolume   uint64 `json:"uplinkVolume,omitempty"`
	DownlinkVolume uint64 `json:"downlinkVolume,omitempty"`
}

func (s *SubsequentVolumeThreshold) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5
	tmpUint8 := btou(s.Dlvol)<<2 |
		btou(s.Ulvol)<<1 |
		btou(s.Tovol)
	data = append([]byte(""), byte(tmpUint8))
	idx = idx + 1

	// Octet m to (m+7)
	if s.Tovol {
		data = append(data, make([]byte, 8)...)
		binary.BigEndian.PutUint64(data[idx:], s.TotalVolume)
		idx = idx + 8
	}

	// Octet p to (p+7)
	if s.Ulvol {
		data = append(data, make([]byte, 8)...)
		binary.BigEndian.PutUint64(data[idx:], s.UplinkVolume)
		idx = idx + 8
	}

	// Octet q to (q+7)
	if s.Dlvol {
		data = append(data, make([]byte, 8)...)
		binary.BigEndian.PutUint64(data[idx:], s.DownlinkVolume)
	}

	return data, nil
}

func (s *SubsequentVolumeThreshold) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie SubsequentVolumeThreshold Inadequate TLV length: %d", length)
	}
	s.Dlvol = utob(uint8(data[idx]) & BitMask3)
	s.Ulvol = utob(uint8(data[idx]) & BitMask2)
	s.Tovol = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet m to (m+7)
	if s.Tovol {
		if length < idx+8 {
			return fmt.Errorf("ie SubsequentVolumeThreshold.TotalVolume Inadequate TLV length: %d", length)
		}
		s.TotalVolume = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	// Octet p to (p+7)
	if s.Ulvol {
		if length < idx+8 {
			return fmt.Errorf("ie SubsequentVolumeThreshold.UplinkVolume Inadequate TLV length: %d", length)
		}
		s.UplinkVolume = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	// Octet q to (q+7)
	if s.Dlvol {
		if length < idx+8 {
			return fmt.Errorf("ie SubsequentVolumeThreshold.DownlinkVolume Inadequate TLV length: %d", length)
		}
		s.DownlinkVolume = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	return nil
}

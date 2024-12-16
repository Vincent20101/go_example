package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// VolumeQuota IE. Refer to [TS29244 8.2.50  Volume Quota].
//
// The Volume Quota IE type shall be encoded as shown in Figure 8.2.50-1. It
// contains the volume quota to be monitored by the UP function.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                   Type = 73(decimal)                          |
//	         |---------------------------------------------------------------|
//	3 to 4   |                      Length = n                               |
//	         |---------------------------------------------------------------|
//	  5      |             Spare                     | DLVOL | ULVOL | TOVOL |
//	         |---------------------------------------------------------------|
//
// m to m+7  |                     Total Volume                              |
//
//	|---------------------------------------------------------------|
//
// p to p+7  |                     Uplink Volume                             |
//
//	|---------------------------------------------------------------|
//
// q to q+7  |                    Downlink Volume                            |
//
//	|---------------------------------------------------------------|
//
// S to (n+4)|   These octets(s) is/are present only if explicity specified  |
//
//	+---------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
//
//   - Bit 1 – TOVOL: If this bit is set to "1", then the Total Volume field
//     shall be present, otherwise the Total Volume field shall not be present.
//
//   - Bit 2 – ULVOL: If this bit is set to "1", then the Uplink Volume field
//     shall be present, otherwise the Uplink Volume field shall not be present.
//
//   - Bit 3 – DLVOL: If this bit is set to "1", then the Downlink Volume
//     field shall be present, otherwise the Downlink Volume field shall not be
//     present.
//
//   - Bit 4 to bit 8: Spare, for future use and set to "0".
//
// At least one bit shall be set to "1". Several bits may be set to "1".
//
// The Total Volume, Uplink Volume and Downlink Volume fields shall be encoded
// as an Unsigned64 binary integer value. They shall contain the total, uplink
// or downlink number of octets respectively.
type VolumeQuota struct {
	Dlvol          bool   `json:"dlvol,omitempty"`
	Ulvol          bool   `json:"ulvol,omitempty"`
	Tovol          bool   `json:"tovol,omitempty"`
	TotalVolume    uint64 `json:"totalVolume,omitempty"`
	UplinkVolume   uint64 `json:"uplinkVolume,omitempty"`
	DownlinkVolume uint64 `json:"downlinkVolume,omitempty"`
}

func (v *VolumeQuota) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5
	tmpUint8 := btou(v.Dlvol)<<2 |
		btou(v.Ulvol)<<1 |
		btou(v.Tovol)
	data = append([]byte(""), byte(tmpUint8))
	idx = idx + 1

	// Octet m to (m+7)
	if v.Tovol {
		data = append(data, make([]byte, 8)...)
		binary.BigEndian.PutUint64(data[idx:], v.TotalVolume)
		idx = idx + 8
	}

	// Octet p to (p+7)
	if v.Ulvol {
		data = append(data, make([]byte, 8)...)
		binary.BigEndian.PutUint64(data[idx:], v.UplinkVolume)
		idx = idx + 8
	}

	// Octet q to (q+7)
	if v.Dlvol {
		data = append(data, make([]byte, 8)...)
		binary.BigEndian.PutUint64(data[idx:], v.DownlinkVolume)
	}

	return data, nil
}

func (v *VolumeQuota) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie VolumeQuota Inadequate TLV length: %d", length)
	}
	v.Dlvol = utob(uint8(data[idx]) & BitMask3)
	v.Ulvol = utob(uint8(data[idx]) & BitMask2)
	v.Tovol = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet m to (m+7)
	if v.Tovol {
		if length < idx+8 {
			return fmt.Errorf("ie VolumeQuota.TotalVolume Inadequate TLV length: %d", length)
		}
		v.TotalVolume = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	// Octet p to (p+7)
	if v.Ulvol {
		if length < idx+8 {
			return fmt.Errorf("ie VolumeQuota.UplinkVolume Inadequate TLV length: %d", length)
		}
		v.UplinkVolume = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	// Octet q to (q+7)
	if v.Dlvol {
		if length < idx+8 {
			return fmt.Errorf("ie VolumeQuota.DownlinkVolume Inadequate TLV length: %d", length)
		}
		v.DownlinkVolume = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	return nil
}

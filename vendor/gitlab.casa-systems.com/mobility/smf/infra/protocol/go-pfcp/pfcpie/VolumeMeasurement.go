package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type VolumeMeasurement struct {
	Dlvol          bool   `json:"dlvol,omitempty"`
	Ulvol          bool   `json:"ulvol,omitempty"`
	Tovol          bool   `json:"tovol,omitempty"`
	TotalVolume    uint64 `json:"totalVolume,omitempty"`
	UplinkVolume   uint64 `json:"uplinkVolume,omitempty"`
	DownlinkVolume uint64 `json:"downlinkVolume,omitempty"`
}

func (v *VolumeMeasurement) MarshalBinary() (data []byte, err error) {
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

func (v *VolumeMeasurement) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie VolumeMeasurement Inadequate TLV length: %d", length)
	}
	v.Dlvol = utob(uint8(data[idx]) & BitMask3)
	v.Ulvol = utob(uint8(data[idx]) & BitMask2)
	v.Tovol = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet m to (m+7)
	if v.Tovol {
		if length < idx+8 {
			return fmt.Errorf("ie VolumeMeasurement.TotalVolume Inadequate TLV length: %d", length)
		}
		v.TotalVolume = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	// Octet p to (p+7)
	if v.Ulvol {
		if length < idx+8 {
			return fmt.Errorf("ie VolumeMeasurement.UplinkVolume Inadequate TLV length: %d", length)
		}
		v.UplinkVolume = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	// Octet q to (q+7)
	if v.Dlvol {
		if length < idx+8 {
			return fmt.Errorf("ie VolumeMeasurement.DownlinkVolume Inadequate TLV length: %d", length)
		}
		v.DownlinkVolume = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	return nil
}

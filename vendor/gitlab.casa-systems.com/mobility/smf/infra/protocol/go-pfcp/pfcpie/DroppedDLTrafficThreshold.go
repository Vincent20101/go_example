package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type DroppedDLTrafficThreshold struct {
	Dlby                        bool   `json:"dlby,omitempty"`
	Dlpa                        bool   `json:"dlpa,omitempty"`
	DownlinkPackets             uint64 `json:"downlinkPackets,omitempty"`
	NumberOfBytesOfDownlinkData uint64 `json:"numberOfBytesOfDownlinkData,omitempty"`
}

func (d *DroppedDLTrafficThreshold) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5
	tmpUint8 := btou(d.Dlby)<<1 | btou(d.Dlpa)
	data = append([]byte(""), byte(tmpUint8))
	idx = idx + 1

	// Octet m to (m+7)
	if d.Dlpa {
		data = append(data, make([]byte, 8)...)
		binary.BigEndian.PutUint64(data[idx:], d.DownlinkPackets)
		idx = idx + 8
	}

	// Octet o to (o+7)
	if d.Dlby {
		data = append(data, make([]byte, 8)...)
		binary.BigEndian.PutUint64(data[idx:], d.NumberOfBytesOfDownlinkData)
	}

	return data, nil
}

func (d *DroppedDLTrafficThreshold) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE DroppedDLTrafficThreshold Inadequate TLV length: %d", length)
	}
	d.Dlby = utob(uint8(data[idx]) & BitMask2)
	d.Dlpa = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet m to (m+7)
	if d.Dlpa {
		if length < idx+8 {
			return fmt.Errorf("IE DroppedDLTrafficThreshold Field DownlinkPackets Inadequate TLV length: %d", length)
		}
		d.DownlinkPackets = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	// Octet o to (o+7)
	if d.Dlby {
		if length < idx+8 {
			return fmt.Errorf("IE DroppedDLTrafficThreshold Field NumberOfBytesOfDownlinkData Inadequate TLV length: %d", length)
		}
		d.NumberOfBytesOfDownlinkData = binary.BigEndian.Uint64(data[idx:])
		idx = idx + 8
	}

	return nil
}

package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type AggregatedURRID struct {
	AggregatedUrrIdValue uint32 `json:"aggregatedUrrIdValue,omitempty"`
}

func (a *AggregatedURRID) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data[idx:], a.AggregatedUrrIdValue)

	return data, nil
}

func (a *AggregatedURRID) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE AggregatedURRID Inadequate TLV length: %d", length)
	}
	a.AggregatedUrrIdValue = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

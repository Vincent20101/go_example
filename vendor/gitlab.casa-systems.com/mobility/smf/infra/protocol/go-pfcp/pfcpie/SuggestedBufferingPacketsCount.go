package pfcpie

import (
	"fmt"
)

type SuggestedBufferingPacketsCount struct {
	PacketCountValue uint8 `json:"packetCountValue,omitempty"`
}

func (s *SuggestedBufferingPacketsCount) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(s.PacketCountValue))

	return data, nil
}

func (s *SuggestedBufferingPacketsCount) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie SuggestedBufferingPacketsCount Inadequate TLV length: %d", length)
	}
	s.PacketCountValue = uint8(data[idx])
	idx = idx + 1

	return nil
}

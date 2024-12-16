package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type AdditionalURLs struct {
	URLs []byte `json:"uRLs,omitempty"`
}

func (a *AdditionalURLs) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	data = append(data, make([]byte, 2)...)
	binary.BigEndian.PutUint16(data[idx:], uint16(len(a.URLs)))
	data = append(data, a.URLs...)

	return
}

func (a *AdditionalURLs) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	if length < idx+2 {
		return fmt.Errorf("IE AdditionalURLs Inadequate TLV length: %d", length)
	}
	urlLen := binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2

	if length < idx+urlLen {
		return fmt.Errorf("IE AdditionalURLs Field URLs Inadequate TLV length: %d, idx: %d, urlLen: %d", length, idx, urlLen)
	}
	a.URLs = append(a.URLs, data[idx:idx+urlLen]...)

	return nil
}

package pfcpie

import "fmt"

type EthernetPDUSessionInformation struct {
	Ethi bool `json:"ethi,omitempty"`
}

func (e *EthernetPDUSessionInformation) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(e.Ethi)
	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (e *EthernetPDUSessionInformation) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE EthernetPDUSessionInformation Inadequate TLV length: %d", length)
	}
	e.Ethi = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

package pfcpie

import fmt "fmt"

type UsageInformation struct {
	BEF bool `json:"bEF,omitempty"`
	AFT bool `json:"aFT,omitempty"`
	UAE bool `json:"uAE,omitempty"`
	UBE bool `json:"uBE,omitempty"`
}

func (ui *UsageInformation) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(ui.BEF) |
		(btou(ui.AFT) << 1) |
		(btou(ui.UAE) << 2) |
		(btou(ui.UBE) << 3)
	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (ui *UsageInformation) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie UsageInformation Inadequate TLV length: %d", length)
	}
	ui.BEF = utob(uint8(data[idx]) & BitMask1)
	ui.AFT = utob(uint8(data[idx]) & BitMask2)
	ui.UAE = utob(uint8(data[idx]) & BitMask3)
	ui.UBE = utob(uint8(data[idx]) & BitMask4)
	idx = idx + 1

	return nil
}

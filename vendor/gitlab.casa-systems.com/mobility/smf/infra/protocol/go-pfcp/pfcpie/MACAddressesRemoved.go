package pfcpie

import (
	"fmt"
	"math"
)

type MACAddressesRemoved struct {
	MACAddressValues [][]byte `json:"mACAddressValues,omitempty"`
	CTAG             *CTAG    `json:"cTAG,omitempty"`
	STAG             *STAG    `json:"sTAG,omitempty"`
}

func (m *MACAddressesRemoved) MarshalBinary() ([]byte, error) {
	var data []byte
	// Octet 5
	numberOfMACAddresses := len(m.MACAddressValues)
	if numberOfMACAddresses > math.MaxUint8 {
		return nil, fmt.Errorf("number of mac addresses overflowed: %d", numberOfMACAddresses)
	}
	data = append(data, byte(numberOfMACAddresses))
	for _, macAddressValue := range m.MACAddressValues {
		if len(macAddressValue) != 6 {
			return nil, fmt.Errorf("invalid mac address: %v", macAddressValue)
		}
		data = append(data, macAddressValue...)
	}
	var ctagData []byte
	if m.CTAG != nil {
		var err error
		ctagData, err = m.CTAG.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal CTAG: %v", err)
		}
		if len(ctagData) > math.MaxUint8 {
			return nil, fmt.Errorf("CTAGS length overflowed: %d", len(ctagData))
		}
	}
	data = append(data, byte(len(ctagData)))
	data = append(data, ctagData...)
	var stagData []byte
	if m.STAG != nil {
		var err error
		stagData, err = m.STAG.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("failed to marshal STAG: %v", err)
		}
		if len(stagData) > math.MaxUint8 {
			return nil, fmt.Errorf("STAGS length overflowed: %d", len(stagData))
		}
	}
	data = append(data, byte(len(stagData)))
	data = append(data, stagData...)
	return data, nil
}

func (m *MACAddressesRemoved) UnmarshalBinary(data []byte) error {
	length := len(data)

	if len(data) < 1 {
		return fmt.Errorf("IE MACAddressesRemoved Inadequate TLV length: %d", length)
	}
	numberOfMACAddresses := int(data[0])
	data = data[1:]

	if numberOfMACAddresses >= 1 {
		if len(data) < numberOfMACAddresses*6 {
			return fmt.Errorf("IE MACAddressesRemoved Field MACAddressValues Inadequate TLV length: %d", length)
		}
		m.MACAddressValues = make([][]byte, numberOfMACAddresses)
		for i, j := 0, 0; i < numberOfMACAddresses; i, j = i+1, j+6 {
			macAddressValue := make([]byte, 6)
			copy(macAddressValue, data[:6])
			data = data[6:]
			m.MACAddressValues[i] = macAddressValue
		}
	}

	if len(data) < 1 {
		return fmt.Errorf("IE MACAddressesRemoved Inadequate TLV length: %d", length)
	}
	ctagLength := int(data[0])
	data = data[1:]

	if ctagLength >= 1 {
		var ctag CTAG
		if err := ctag.UnmarshalBinary(data[:ctagLength]); err != nil {
			return fmt.Errorf("IE MACAddressesRemoved failed to unmarshal CTAG: %v", err)
		}
		data = data[ctagLength:]
		m.CTAG = &ctag
	}

	if len(data) < 1 {
		return fmt.Errorf("IE MACAddressesRemoved Inadequate TLV length: %d", length)
	}
	stagLength := int(data[0])
	data = data[1:]

	if stagLength >= 1 {
		var stag STAG
		if err := stag.UnmarshalBinary(data[:stagLength]); err != nil {
			return fmt.Errorf("IE MACAddressesRemoved failed to unmarshal STAG: %v", err)
		}
		data = data[stagLength:]
		m.STAG = &stag
	}

	return nil
}

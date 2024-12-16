package pfcpie

import (
	"encoding/hex"
	"fmt"
)

const (
	// ValidationAlgorithmHMACSHA256 is Validation Algorithm HMAC-SHA-256.
	ValidationAlgorithmHMACSHA256 = 1
	// ValidationAlgorithmHMACSHA3 is Validation Algorithm HMAC-SHA-3.
	ValidationAlgorithmHMACSHA3 = 2
)

// ValidationParameter implements Validation Parameter IE.
// IE Type: 32793
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	 1 to 2  |                     Type = 32793 (decimal)                     |
//	         |----------------------------------------------------------------|
//	 3 to 4  |                          Length = n                            |
//	         |----------------------------------------------------------------|
//	 5 to 6  |                         Enterprise Id                          |
//	         |----------------------------------------------------------------|
//	   7     |                     Length of Provider ID                      |
//	         |----------------------------------------------------------------|
//	 8 to m  |                           Provider ID                          |
//	         |----------------------------------------------------------------|
//	   x     |                       Validation Algorithm                     |
//	         |----------------------------------------------------------------|
//	  x+1    |                      Length of Validation Key                  |
//	         |----------------------------------------------------------------|
//	p+1 to q |                          Validation Key                        |
//	         |----------------------------------------------------------------|
//	q+1 to r |  These octet(s) is/are present only if explicitly specified    |
//	         +----------------------------------------------------------------+
//	                              Validation Prarmeter IE.
type ValidationParameter struct {
	EnterpriseID        uint16 `json:"enterpriseID,omitempty"`
	ProviderID          string `json:"providerID,omitempty"`
	ValidationAlogrithm uint8  `json:"validationAlogrithm,omitempty"`
	ValidationKey       string `json:"validationKey,omitempty"`
}

// NewValidationParameter return a new ValidationParameter.
func NewValidationParameter() *ValidationParameter {
	return &ValidationParameter{
		EnterpriseID: CasaEnterpriseID,
	}
}

// MarshalBinary encodes ValidationParameter.
func (e *ValidationParameter) MarshalBinary() (data []byte, err error) {
	data = AppendUint16(data, e.EnterpriseID)
	// ProviderID
	providerIDLen := len(e.ProviderID)
	data = append(data, byte(providerIDLen))
	data = append(data, []byte(e.ProviderID)...)
	// ValidationAlogrithm
	data = append(data, e.ValidationAlogrithm)
	// decode ValidationKey to byte
	validationKey, err := hex.DecodeString(e.ValidationKey)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode ValidationKey, %v", e)
	}
	data = append(data, byte(len(validationKey)))
	data = append(data, validationKey...)

	return data, nil
}

// UnmarshalBinary unmarshal ValidationParameter from binary
func (e *ValidationParameter) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)

	// EnterpriseID
	eid, err := byteReader.ReadUint16()
	if err != nil {
		return fmt.Errorf("Inadequate length for EnterpriseId, %v", e)
	}
	e.EnterpriseID = eid

	// Length of Provider ID
	providerIDLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Length of Provier ID', %v", e)
	}
	// Provider ID
	providerID, err := byteReader.ReadBytes(int(providerIDLen))
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Provier ID', %v", e)
	}
	e.ProviderID = string(providerID)

	// Validation Alogrithm
	validationAlogrithm, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Validation Alogrithm', %v", e)
	}
	e.ValidationAlogrithm = validationAlogrithm

	// EncyptionKey
	validationKeyLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Length of Validation Key', %v", e)
	}
	// Provider ID
	validationKey, err := byteReader.ReadBytes(int(validationKeyLen))
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Validation Key', %v", e)
	}
	e.ValidationKey = hex.EncodeToString(validationKey)

	return nil
}

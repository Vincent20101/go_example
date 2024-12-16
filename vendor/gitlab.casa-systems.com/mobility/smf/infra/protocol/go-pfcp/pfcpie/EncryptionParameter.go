package pfcpie

import (
	"encoding/hex"
	"fmt"
)

const (
	// EncryptionAlgorithmAESCBC128 is encryption algrithm AES-CBC-128.
	EncryptionAlgorithmAESCBC128 = 1
	// EncryptionAlgorithmAESCBC256 is encryption algrithm AES-CBC-256.
	EncryptionAlgorithmAESCBC256 = 2
)
const (
	// Encryption Parameter Type: USERID
	EncryptionTypeUserId = 1
	// Encryption Parameter Type: X-HEADER
	EncryptionTypeXHeader = 2
)

// EncryptionParameter implements Encryption Parameter IE.
// IE Type: 32792
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	 1 to 2  |                     Type = 32792 (decimal)                     |
//	         |----------------------------------------------------------------|
//	 3 to 4  |                          Length = n                            |
//	         |----------------------------------------------------------------|
//	 5 to 6  |                         Enterprise Id                          |
//	         |----------------------------------------------------------------|
//	   7     |                     Length of Provider ID                      |
//	         |----------------------------------------------------------------|
//	 8 to m  |                           Provider ID                          |
//	         |----------------------------------------------------------------|
//	  m+1    |                       Encryption Algorithm                     |
//	         |----------------------------------------------------------------|
//	  m+2    |                  Length of Initialization Vector               |
//	         |----------------------------------------------------------------|
//
// m+3 to o  |                       Initialization Vector                    |
//
//	         |----------------------------------------------------------------|
//	   p     |                      Length of Encryption Key                  |
//	         |----------------------------------------------------------------|
//	p+1 to q |                          Encryption Key                        |
//	         |----------------------------------------------------------------|
//	q+1 to r |  These octet(s) is/are present only if explicitly specified    |
//	         +----------------------------------------------------------------+
//	                              Encryption Prarmeter IE.
type EncryptionParameter struct {
	EnterpriseID         uint16 `json:"enterpriseID,omitempty"`
	EncryptionType       uint8  `json:"encryptionType,omitempty"` //Enumeration type : USERID: 1 ;X-HEADER: 2
	ProviderID           string `json:"providerID,omitempty"`
	EncryptionAlogrithm  uint8  `json:"encryptionAlogrithm,omitempty"` //Enumeration type : AES-CBC-128.: 1 ;X-AES-CBC-256: 2
	InitializationVector string `json:"initializationVector,omitempty"`
	EncryptionKey        string `json:"encryptionKey,omitempty"`
}

// NewEncryptionParameter return a new EncryptionParameter.
func NewEncryptionParameter() *EncryptionParameter {
	return &EncryptionParameter{
		EnterpriseID: CasaEnterpriseID,
	}
}

// MarshalBinary encodes EncryptionParameter.
func (e *EncryptionParameter) MarshalBinary() (data []byte, err error) {
	data = AppendUint16(data, e.EnterpriseID)
	//Encryption Parameter Type
	data = append(data, e.EncryptionType)
	// ProviderID
	providerIDLen := len(e.ProviderID)
	data = append(data, byte(providerIDLen))
	data = append(data, []byte(e.ProviderID)...)
	// EncryptionAlogrithm
	data = append(data, e.EncryptionAlogrithm)
	//InitializationVector
	initVetor, err := hex.DecodeString(e.InitializationVector)
	if err != nil {
		return nil, fmt.Errorf("decode InitializationVector hex to stringfailed: %v", e)
	}
	data = append(data, byte(len(initVetor)))
	data = append(data, initVetor...)

	// decode EncryptionKey to []byte
	encryptionKey, err := hex.DecodeString(e.EncryptionKey)
	if err != nil {
		return nil, fmt.Errorf("decode encryptionKey hex to string failed: %v", e)
	}
	// EncryptionKey Len
	data = append(data, byte(len(encryptionKey)))
	data = append(data, encryptionKey...)
	return data, nil
}

// UnmarshalBinary unmarshal EncryptionParameter from binary
func (e *EncryptionParameter) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)

	// EnterpriseID
	eid, err := byteReader.ReadUint16()
	if err != nil {
		return fmt.Errorf("Inadequate length for EnterpriseId, %v", e)
	}
	e.EnterpriseID = eid

	//Encryption Parameter Type
	ety, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for Encryption Parameter Type, %v", e)
	}
	e.EncryptionType = ety

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

	// Encryption Alogrithm
	encryptionAlogrithm, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Encryption Alogrithm', %v", e)
	}
	e.EncryptionAlogrithm = encryptionAlogrithm

	// Length of Initialization Vector
	initVectorLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Length of Initialization Vector', %v", e)
	}
	// Initialization Vector
	initVector, err := byteReader.ReadBytes(int(initVectorLen))
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Initialization Vector', %v", e)
	}
	e.InitializationVector = hex.EncodeToString(initVector)

	// EncyptionKey Len
	encryptionKeyLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Length of Encryption Key', %v", e)
	}
	// EncyptionKey
	encryptionKey, err := byteReader.ReadBytes(int(encryptionKeyLen))
	if err != nil {
		return fmt.Errorf("Inadequate length for 'Encryption Key', %v", e)
	}
	e.EncryptionKey = hex.EncodeToString(encryptionKey)

	return nil
}

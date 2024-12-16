package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// SniPool IE.
// SNI: Server Name Indication
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                        Type = 32789(decimal)                   |
//	         |----------------------------------------------------------------|
//	3 to 4   |                          Length = n                            |
//	         |----------------------------------------------------------------|
//	5 to 6   |                        Enterprise ID                           |
//	         |----------------------------------------------------------------|
//	  7      |                       Length of Pool Name                      |
//	         |----------------------------------------------------------------|
//	8 to 8+a |                            Pool Name                           |
//	         |----------------------------------------------------------------|
//	d to d+1 |                       Lenght of Sni Values                     |
//	         |----------------------------------------------------------------|
//	m to m+k |                            Sni Values                          |
//	         |----------------------------------------------------------------|
//	k to n+4 |                               ...                              |
//	         +----------------------------------------------------------------+
//
// SNI Values field of the Sni Pool IE shall be encoded as shown below:
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//
// a+2 to a+3|                       Length of SNI 1                          |
//
//	         |----------------------------------------------------------------|
//	a+4 to o |                             SNI 1                              |
//	         |----------------------------------------------------------------|
//	b to b+1 |                       Length of SNI 2                          |
//	         |----------------------------------------------------------------|
//
// b+2 to c  |                             SNI 2                              |
//
//	         |----------------------------------------------------------------|
//	  ...    |                             ...                                |
//	         |----------------------------------------------------------------|
//	c to c+1 |                       Lenght of SNI m                          |
//	         |----------------------------------------------------------------|
//	c to d   |                             SNI m                              |
//	         +----------------------------------------------------------------+
//	                             Figure 2-3 Sni Values Field
type SniPool struct {
	EnterpriseID uint16   `json:"enterpriseID,omitempty"`
	PoolName     string   `json:"poolName,omitempty"`  //SNI pool name
	SniValues    []string `json:"sniValues,omitempty"` //Ex: tls-sni = x-streams-revvel-stgec.uplynk.com
}

// NewSniPool return a new SniPool with some filled fields.
func NewSniPool() *SniPool {
	return &SniPool{
		EnterpriseID: CasaEnterpriseID,
		PoolName:     "",
		SniValues:    []string{},
	}
}

// encodeSniValues encodes SniVal fields of SniPool.
func (s *SniPool) encodeSniValues() (data []byte, err error) {
	data = []byte{0x00, 0x00}
	// Sni Values.
	for _, v := range s.SniValues {
		data = AppendUint16(data, uint16(len(v)))
		data = append(data, []byte(v)...)
	}
	total := uint16(len(data) - 2)
	data[0] = byte(total >> 8)
	data[1] = byte(total)

	return data, nil
}

// MarshalBinary marshal SniPool to octets.
func (s *SniPool) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 2)
	// Enterprise ID
	binary.BigEndian.PutUint16(data, s.EnterpriseID)
	// Length of PoolName
	data = append(data, byte(len(s.PoolName)))
	// PoolName
	data = append(data, []byte(s.PoolName)...)
	// Sni value
	sniValues, e := s.encodeSniValues()
	if e != nil {
		return nil, e
	}
	data = append(data, sniValues...)
	return data, nil
}

// UnmarshalBinary decode SniPool from octets.
func (s *SniPool) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)

	// EnterpriseID
	eid, e := byteReader.ReadUint16()
	if e != nil {
		return fmt.Errorf("Inadequate length for EnterpriseId, %v", e)
	}
	s.EnterpriseID = eid

	// Pool Name
	poolNameLen, e := byteReader.ReadUint8()
	if e != nil {
		return fmt.Errorf("Inadequate length for Length of Pool Name,  %v", e)
	}
	poolName, e := byteReader.ReadBytes(int(poolNameLen))
	if e != nil {
		return fmt.Errorf("Inadequate length for Pool Name, %v", e)
	}
	s.PoolName = string(poolName)

	// SNI Values
	sniValLen, e := byteReader.ReadUint16()
	if e != nil {
		return fmt.Errorf("Inadequate length for Length of SNI Values, %v", e)
	}
	sniBytes, e := byteReader.ReadBytes(int(sniValLen))
	sniReader := NewByteReader(sniBytes)
	if e != nil {
		return fmt.Errorf("Inadequate length for SNI Values, %v", e)
	}
	for {
		sniLen, _ := sniReader.ReadUint16()
		// Break when sniReader is drained.
		if sniLen == 0 {
			break
		}
		sni, _ := sniReader.ReadBytes(int(sniLen))
		s.SniValues = append(s.SniValues, string(sni))
	}

	return nil
}

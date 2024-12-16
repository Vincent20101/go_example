package pfcpie

import (
	"encoding/binary"
	"fmt"
)

type AdditionalDomainNameAndDomainNameProtocols struct {
	DomainNames         []byte `json:"domainNames,omitempty"`
	DomainNameProtocols []byte `json:"domainNameProtocols,omitempty"`
}

func (a *AdditionalDomainNameAndDomainNameProtocols) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	data = append(data, make([]byte, 2)...)
	idx += 2
	binary.BigEndian.PutUint16(data[idx:], uint16(len(a.DomainNames)))
	data = append(data, a.DomainNames...)
	idx += uint16(len(a.DomainNames))

	data = append(data, make([]byte, 2)...)
	idx += 2
	binary.BigEndian.PutUint16(data[idx:], uint16(len(a.DomainNameProtocols)))
	data = append(data, a.DomainNameProtocols...)

	return
}

func (a *AdditionalDomainNameAndDomainNameProtocols) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	if length < idx+2 {
		return fmt.Errorf("IE AdditionalDomainNameAndDomainNameProtocols Field Dn Inadequate TLV length: %d", length)
	}

	dnLen := binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2

	if length < idx+dnLen {
		return fmt.Errorf("IE AdditionalDomainNameAndDomainNameProtocols Field DomainNames Inadequate TLV length: %d, idx: %d, dnLen: %d", length, idx, dnLen)
	}
	a.DomainNames = append(a.DomainNames, data[idx:idx+dnLen]...)
	idx += dnLen

	if length < idx+2 {
		return fmt.Errorf("IE AdditionalDomainNameAndDomainNameProtocols Field Dnp Inadequate TLV length: %d", length)
	}
	dnpLen := binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2

	if length < idx+dnpLen {
		return fmt.Errorf("IE AdditionalDomainNameAndDomainNameProtocols Field DomainNameProtocols Inadequate TLV length: %d, idx: %d, dnpLen: %d", length, idx, dnpLen)
	}
	a.DomainNameProtocols = append(a.DomainNameProtocols, data[idx:idx+dnpLen]...)

	return nil
}

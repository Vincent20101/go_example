package pfcpie

import (
	"encoding/binary"
	"errors"
	"fmt"
)

// Use Casa's EnterpriseId
const VerizonSpecificUPFIpChunksInfo_ENTERPRISEID = 20858

// VerizonSpecificUPFIpChunksInfo IE
// Casa Specific IE for Verizon use
//
//	Octets        8       7       6       5       4       3       2       1
//	          +--------------------------------------------------------------+
//	1 to 2    |                   Type = xxx(decimal)                        |
//	          |--------------------------------------------------------------|
//	3 to 4    |                      Length = n-4                            |
//	          |--------------------------------------------------------------|
//	5 to 6    |                     EnterpriseID                             |
//	          |--------------------------------------------------------------|
//	  7       |                    IpPoolChunksNum                           |
//	          |--------------------------------------------------------------|
//	8 to n    |                      IpPoolChunks                            |
//	          +--------------------------------------------------------------+
//
// deprecated
type VerizonUPFIpChunksInfo struct {
	EnterpriseId uint16        `json:"enterpriseId,omitempty"`
	IpPoolChunks []IPChunkInfo `json:"ipPoolChunks,omitempty"`
}

func NewVerizonUPFIpChunksInfo() *VerizonUPFIpChunksInfo {
	return &VerizonUPFIpChunksInfo{
		EnterpriseId: CasaEnterpriseID,
	}
}

func (a *VerizonUPFIpChunksInfo) MarshalBinary() ([]byte, error) {
	// EnterpriseID
	data := make([]byte, 2)
	enterpriseId := uint16(VerizonSpecificUPFIpChunksInfo_ENTERPRISEID)
	binary.BigEndian.PutUint16(data, enterpriseId)

	// IpPoolChunksNum
	data = append(data, uint8(len(a.IpPoolChunks)))

	// IpPoolChunks
	for _, item := range a.IpPoolChunks {
		value, err := item.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Failed to IpChunkInfo.MarshalBinary(), Err: %v", err)
		}
		data = append(data, uint8(len(value)))
		data = append(data, value...)
	}
	return data, nil
}

func (a *VerizonUPFIpChunksInfo) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)
	EnterpriseID, err := byteReader.ReadUint16()
	if err != nil {
		return fmt.Errorf("Inadequate length for enterpriseId: %v", err)
	}

	// EnterpriseID
	a.EnterpriseId = EnterpriseID

	// IpChunkNum
	ipChunkNum, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("Inadequate length for IsCgNat and IpPoolChunksNum: %v", err)
	}

	// IpPoolChunks
	for i := 0; i < int(ipChunkNum); i++ {
		ipChunkInfoLen, err := byteReader.ReadUint8()
		if err != nil {
			return errors.New("The length of data(IpPoolChunks) is incorrect.")
		}
		ipChunkBytes, err := byteReader.ReadBytes(int(ipChunkInfoLen))
		if err != nil {
			return fmt.Errorf("Inadequate length for ipchunk[%d]: %v", i, err)
		}
		ipChunk := IPChunkInfo{}
		err = ipChunk.UnmarshalBinary(ipChunkBytes)
		if err != nil {
			return fmt.Errorf("unmarshal ipchunk[%d] failed: %v", i, err)
		}
		a.IpPoolChunks = append(a.IpPoolChunks, ipChunk)
	}
	return nil
}

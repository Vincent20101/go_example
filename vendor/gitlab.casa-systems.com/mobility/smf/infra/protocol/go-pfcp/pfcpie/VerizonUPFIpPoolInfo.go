package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// VerizonSpecificUPFIpPoolInfo IE
// Casa Specific IE for Verizon use
//
//	Octets        8       7       6       5       4       3       2       1
//	          +--------------------------------------------------------------+
//	1 to 2    |                   Type = 32799(decimal)                      |
//	          |--------------------------------------------------------------|
//	3 to 4    |                      Length = n                              |
//	          |--------------------------------------------------------------|
//	5 to 6    |                     EnterpriseID                             |
//	          |--------------------------------------------------------------|
//	  7       |                   Number of upfInfo(s) (up to 32 entries)    |
//	          |--------------------------------------------------------------|
//	8 to n    |                      list of upfInfo(s)                      |
//	          +--------------------------------------------------------------+

type VerizonUPFIpPoolInfo struct {
	EnterpriseId uint16              `json:"enterpriseId,omitempty"`
	UpfInfoList  []UPFInfoIpPoolInfo `json:"upfInfoList,omitempty"`
}

func (a *VerizonUPFIpPoolInfo) MarshalBinary() ([]byte, error) {
	// EnterpriseID
	data := make([]byte, 2)
	enterpriseId := uint16(VerizonSpecificUPFIpChunksInfo_ENTERPRISEID)
	binary.BigEndian.PutUint16(data, enterpriseId)

	// IpPoolChunksNum
	data = append(data, uint8(len(a.UpfInfoList)))

	// UpfInfoList
	for _, item := range a.UpfInfoList {
		value, err := item.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("VerizonUPFIpPoolInfo.MarshalBinary failed, err: %v", err)
		}
		dl := uint16(len(value))
		data = append(data, byte(dl>>8), byte(dl))
		data = append(data, value...)
	}
	return data, nil
}

func (a *VerizonUPFIpPoolInfo) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)
	EnterpriseID, err := byteReader.ReadUint16()
	if err != nil {
		return fmt.Errorf("VerizonUPFIpPoolInfo.UnmarshalBinary read enterpriseId len failed,err:%v", err)
	}

	// EnterpriseID
	a.EnterpriseId = EnterpriseID

	// UpfInfoList len
	ul, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("VerizonUPFIpPoolInfo.UnmarshalBinary read UpfInfoList len failed,err:%v", err)
	}

	// UpfInfoList
	for i := 0; i < int(ul); i++ {
		upfInfoLen, err := byteReader.ReadUint16()
		if err != nil {
			return fmt.Errorf("VerizonUPFIpPoolInfo.UnmarshalBinary read UpfInfo len failed,err:%v", err)
		}
		upfInfoBytes, err := byteReader.ReadBytes(int(upfInfoLen))
		if err != nil {
			return fmt.Errorf("VerizonUPFIpPoolInfo.UnmarshalBinary read UpfInfo Data failed,err:%v", err)
		}
		upfInfo := UPFInfoIpPoolInfo{}
		err = upfInfo.UnmarshalBinary(upfInfoBytes)
		if err != nil {
			return fmt.Errorf("VerizonUPFIpPoolInfo.UnmarshalBinary read UpfInfo failed,err:%v", err)
		}
		a.UpfInfoList = append(a.UpfInfoList, upfInfo)
	}
	return nil
}

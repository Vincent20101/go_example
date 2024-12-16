package pfcpie

import "fmt"

// VerizonSpecificUPFInfoIpPoolInfo IE
// Each upfInfo/PPE instance has a upfInfo defined in upfInfoList.
// Casa Specific IE for Verizon use
//
//	Octets        8       7       6       5       4       3       2       1
//	          +--------------------------------------------------------------+
//	1 to 2    |                       Length = n                             |
//	          |--------------------------------------------------------------|
//	3         |                ppe/upfInfo key Length = (p-3)+1              |
//	          |--------------------------------------------------------------|
//	4 to p    |                       ppe/upfInfo key                        |
//	          |--------------------------------------------------------------|
//	  p+1     |                    Number of IpPoolsItem(s)                  |
//	          |--------------------------------------------------------------|
//	p+2 to q  |                    list of IpPoolsItem(s)                    |
//	          +--------------------------------------------------------------+

type UPFInfoIpPoolInfo struct {
	UpfInfoKey string              `json:"upfInfoKey,omitempty"`
	UpfIpPools []UpfIpPoolInfoItem `json:"upfIpPools,omitempty"`
}

func (a *UPFInfoIpPoolInfo) MarshalBinary() (data []byte, err error) {
	keyLen := len(a.UpfInfoKey)
	data = append(data, uint8(keyLen))
	if keyLen > 0 {
		data = append(data, []byte(a.UpfInfoKey)...)
	}

	poolLen := len(a.UpfIpPools)
	data = append(data, uint8(poolLen))
	for _, v := range a.UpfIpPools {
		d, e := v.MarshalBinary()
		if e != nil {
			return nil, fmt.Errorf("UPFInfoIpPoolInfo.MarshalBinary UpfIpPools failed,err:%s", e.Error())
		}
		dl := uint16(len(d))
		data = append(data, byte(dl>>8), byte(dl))
		data = append(data, d...)
	}

	return data, nil
}

func (a *UPFInfoIpPoolInfo) UnmarshalBinary(data []byte) (err error) {
	byteReader := NewByteReader(data)

	keyLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read UpfInfoKey len failed,err:%s", err.Error())
	}
	if keyLen > 0 {
		key, err := byteReader.ReadBytes(int(keyLen))
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read UpfInfoKey failed,err:%s", err.Error())
		}
		a.UpfInfoKey = string(key)
	}

	poolLen, err := byteReader.ReadUint8()
	if err != nil {
		return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read UpfIpPools len failed,err:%s", err.Error())
	}
	for i := 0; i < int(poolLen); i++ {
		var pool UpfIpPoolInfoItem
		poolDataLen, err := byteReader.ReadUint16()
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read pool data len failed,err:%s", err.Error())
		}
		pd, err := byteReader.ReadBytes(int(poolDataLen))
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read pool data failed,err:%s", err.Error())
		}
		err = pool.UnmarshalBinary(pd)
		if err != nil {
			return fmt.Errorf("UpfIpPoolInfoItem.UnmarshalBinary read pool failed,err:%s", err.Error())
		}
		a.UpfIpPools = append(a.UpfIpPools, pool)
	}

	return nil
}

package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// Use Casa's EnterpriseId
const VerizonSpecificUEPDNInfo_ENTERPRISEID = 20858

const (
	tlv_type_VerizonUEPDNInfo_ApnName            = 1
	tlv_type_VerizonUEPDNInfo_PGW_C_SMFNodeName  = 2
	tlv_type_VerizonUEPDNInfo_PGW_C_SMFIPAddress = 3
	tlv_type_VerizonUEPDNInfo_RATType            = 4
	tlv_type_VerizonUEPDNInfo_UserLocationInfo   = 5
	tlv_type_VerizonUEPDNInfo_SGW_C_IPAddress    = 6
	tlv_type_VerizonUEPDNInfo_ServingPLMN        = 7
	tlv_type_VerizonUEPDNInfo_ChargingId         = 8
)

type VerizonUEPDNInfo struct {
	EnterpriseId       uint16              `json:"enterpriseId,omitempty"`
	ApnName            *ApnName            `json:"apnName,omitempty"`
	PGW_C_SMFNodeName  *PGW_C_SMFNodeName  `json:"pGW_C_SMFNodeName,omitempty"`
	PGW_C_SMFIPAddress *PGW_C_SMFIPAddress `json:"pGW_C_SMFIPAddress,omitempty"`
	RATType            *RATType            `json:"rATType,omitempty"`
	ULI                *UserLocationInfo   `json:"uLI,omitempty"`
	SGW_C_IPAddress    *SGW_C_IPAddress    `json:"sGW_C_IPAddress,omitempty"`
	ServingPLMN        *ServingPLMN        `json:"servingPLMN,omitempty"`
	ChargingId         *ChargingId         `json:"chargingId,omitempty"`
}

func (v *VerizonUEPDNInfo) MarshalBinary() ([]byte, error) {
	idx := 0
	data := make([]byte, 2)
	enterpriseId := uint16(VerizonSpecificUEPDNInfo_ENTERPRISEID)
	binary.BigEndian.PutUint16(data, enterpriseId)
	idx += 2

	if v.ApnName != nil {
		diBuf, err := v.ApnName.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("ApnName Marshal failed: %v", err)
		}
		fLen := byte(len(diBuf))
		tmp := make([]byte, 2)
		tmp[0] = tlv_type_VerizonUEPDNInfo_ApnName
		tmp[1] = fLen
		//binary.BigEndian.PutUint16(tmp, fLen)// TODO - use 2 bytes for len?
		data = append(data, tmp...)
		data = append(data, diBuf...)
	}

	if v.PGW_C_SMFNodeName != nil {
		diBuf, err := v.PGW_C_SMFNodeName.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("PGW_C_SMFNodeName Marshal failed: %v", err)
		}
		fLen := byte(len(diBuf))
		tmp := make([]byte, 2)
		tmp[0] = tlv_type_VerizonUEPDNInfo_PGW_C_SMFNodeName
		tmp[1] = fLen
		//binary.BigEndian.PutUint16(tmp, fLen)
		data = append(data, tmp...)
		data = append(data, diBuf...)
	}

	if v.PGW_C_SMFIPAddress != nil {
		diBuf, err := v.PGW_C_SMFIPAddress.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("PGW_C_SMFIPAddress Marshal failed: %v", err)
		}
		fLen := byte(len(diBuf))
		tmp := make([]byte, 2)
		tmp[0] = tlv_type_VerizonUEPDNInfo_PGW_C_SMFIPAddress
		tmp[1] = fLen
		//binary.BigEndian.PutUint16(tmp, fLen)
		data = append(data, tmp...)
		data = append(data, diBuf...)
	}

	if v.RATType != nil {
		diBuf, err := v.RATType.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("RATType Marshal failed: %v", err)
		}
		fLen := byte(len(diBuf))
		tmp := make([]byte, 2)
		tmp[0] = tlv_type_VerizonUEPDNInfo_RATType
		tmp[1] = fLen
		//binary.BigEndian.PutUint16(tmp, fLen)
		data = append(data, tmp...)
		data = append(data, diBuf...)
	}

	if v.ULI != nil {
		diBuf, err := v.ULI.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("ULI Marshal failed: %v", err)
		}
		fLen := byte(len(diBuf))
		tmp := make([]byte, 2)
		tmp[0] = tlv_type_VerizonUEPDNInfo_UserLocationInfo
		tmp[1] = fLen
		//binary.BigEndian.PutUint16(tmp, fLen)
		data = append(data, tmp...)
		data = append(data, diBuf...)
	}

	if v.SGW_C_IPAddress != nil {
		diBuf, err := v.SGW_C_IPAddress.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("SGW_C_IPAddress Marshal failed: %v", err)
		}
		fLen := byte(len(diBuf))
		tmp := make([]byte, 2)
		tmp[0] = tlv_type_VerizonUEPDNInfo_SGW_C_IPAddress
		tmp[1] = fLen
		//binary.BigEndian.PutUint16(tmp, fLen)
		data = append(data, tmp...)
		data = append(data, diBuf...)
	}

	if v.ServingPLMN != nil {
		diBuf, err := v.ServingPLMN.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("ServingPLMN Marshal failed: %v", err)
		}
		fLen := byte(len(diBuf))
		tmp := make([]byte, 2)
		tmp[0] = tlv_type_VerizonUEPDNInfo_ServingPLMN
		tmp[1] = fLen
		//binary.BigEndian.PutUint16(tmp, fLen)
		data = append(data, tmp...)
		data = append(data, diBuf...)
	}

	if v.ChargingId != nil {
		diBuf, err := v.ChargingId.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("ChargingId Marshal failed: %v", err)
		}
		fLen := byte(len(diBuf))
		tmp := make([]byte, 2)
		tmp[0] = tlv_type_VerizonUEPDNInfo_ChargingId
		tmp[1] = fLen
		//binary.BigEndian.PutUint16(tmp, fLen)
		data = append(data, tmp...)
		data = append(data, diBuf...)
	}

	return data, nil
}

func (v *VerizonUEPDNInfo) UnmarshalBinary(data []byte) error {
	var idx = 0
	length := len(data)
	if length < idx+1 {
		return fmt.Errorf("Inadequate length of UserLocationInfo, got zero length")
	}

	v.EnterpriseId = binary.BigEndian.Uint16(data[idx : idx+2])
	idx += 2

	for idx < length {
		attrType := data[idx]
		attrLen := data[idx+1]
		idx += 2
		switch attrType {
		case tlv_type_VerizonUEPDNInfo_ApnName:
			v.ApnName = &ApnName{}
			err := v.ApnName.UnmarshalBinary(data[idx : idx+int(attrLen)])
			if err != nil {
				return fmt.Errorf("ApnName Unmarshal Failed.")
			}
			idx += int(attrLen)
		case tlv_type_VerizonUEPDNInfo_PGW_C_SMFNodeName:
			v.PGW_C_SMFNodeName = &PGW_C_SMFNodeName{}
			err := v.PGW_C_SMFNodeName.UnmarshalBinary(data[idx : idx+int(attrLen)])
			if err != nil {
				return fmt.Errorf("PGW_C_SMFNodeName Unmarshal Failed.")
			}
			idx += int(attrLen)
		case tlv_type_VerizonUEPDNInfo_PGW_C_SMFIPAddress:
			v.PGW_C_SMFIPAddress = &PGW_C_SMFIPAddress{}
			err := v.PGW_C_SMFIPAddress.UnmarshalBinary(data[idx : idx+int(attrLen)])
			if err != nil {
				return fmt.Errorf("PGW_C_SMFIPAddress Unmarshal Failed.")
			}
			idx += int(attrLen)
		case tlv_type_VerizonUEPDNInfo_RATType:
			v.RATType = &RATType{}
			err := v.RATType.UnmarshalBinary(data[idx : idx+int(attrLen)])
			if err != nil {
				return fmt.Errorf("RATType Unmarshal Failed.")
			}
			idx += int(attrLen)
		case tlv_type_VerizonUEPDNInfo_UserLocationInfo:
			v.ULI = &UserLocationInfo{}
			err := v.ULI.UnmarshalBinary(data[idx : idx+int(attrLen)])
			if err != nil {
				return fmt.Errorf("UserLocationInfo Unmarshal Failed.")
			}
			idx += int(attrLen)
		case tlv_type_VerizonUEPDNInfo_SGW_C_IPAddress:
			v.SGW_C_IPAddress = &SGW_C_IPAddress{}
			err := v.SGW_C_IPAddress.UnmarshalBinary(data[idx : idx+int(attrLen)])
			if err != nil {
				return fmt.Errorf("SGW_C_IPAddress Unmarshal Failed.")
			}
			idx += int(attrLen)
		case tlv_type_VerizonUEPDNInfo_ServingPLMN:
			v.ServingPLMN = &ServingPLMN{}
			err := v.ServingPLMN.UnmarshalBinary(data[idx : idx+int(attrLen)])
			if err != nil {
				return fmt.Errorf("ServingPLMN Unmarshal Failed.")
			}
			idx += int(attrLen)
		case tlv_type_VerizonUEPDNInfo_ChargingId:
			v.ChargingId = &ChargingId{}
			err := v.ChargingId.UnmarshalBinary(data[idx : idx+int(attrLen)])
			if err != nil {
				return fmt.Errorf("ChargingId Unmarshal Failed.")
			}
			idx += int(attrLen)
		}
	}

	return nil
}

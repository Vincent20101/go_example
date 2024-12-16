package pfcpie

import (
	"encoding/json"
	"fmt"
)

// 8.2.144	Port Management Information Container
// The Port Management Information Container IE shall be encoded as shown in Figure 8.2.144-1.
// It contains an opaque container of port management information.
//                                         Bits
//  Octets        8       7       6       5       4       3       2       1
//            +----------------------------------------------------------------+
//  1 to 2    |                   Type = 202 (decimal)                         |
//            |----------------------------------------------------------------|
//  3 to 4    |                      Length = n                                |
//            |----------------------------------------------------------------|
// 5 to (n+4) |                 Port Management Information                    |
//            +----------------------------------------------------------------+
//
// The Port Management Information field shall be encoded as an Octet String.
// It shall encode an Ethernet port management message defined in clause 8 of 3GPP TS 24.519 [63].
//

type PortManagementInformationContainer struct {
	PortManagementInformation []byte `json:"portManagementInformation,omitempty"`
}

// used to marshal/unmarshal json.
type portManagementInformationRaw struct {
	PortManagementInformation string `json:"portManagementInformation,omitempty"`
}

func (a *PortManagementInformationContainer) MarshalBinary() (data []byte, err error) {
	// Octet 5 to (n+4)
	data = a.PortManagementInformation

	return data, nil
}

func (a *PortManagementInformationContainer) UnmarshalBinary(data []byte) error {
	// Octet 5 to (n+4)
	a.PortManagementInformation = data

	return nil
}

func (a *PortManagementInformationContainer) MarshalJSON() ([]byte, error) {
	portManagementInfo := &portManagementInformationRaw{PortManagementInformation: string(a.PortManagementInformation)}
	return json.Marshal(portManagementInfo)
}

func (a *PortManagementInformationContainer) UnmarshalJSON(b []byte) error {
	portManagementInfo := &portManagementInformationRaw{}
	err := json.Unmarshal(b, portManagementInfo)
	if err != nil {
		return fmt.Errorf("unmarshal portManagementInformation from json error: %v ", err)
	}
	a.PortManagementInformation = []byte(portManagementInfo.PortManagementInformation)
	return nil
}

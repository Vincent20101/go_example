package pfcpie

import (
	"encoding/json"
	"fmt"
)

// 8.2.182	Bridge Management Information Container
// The Bridge Management Information Container IE shall be encoded as shown in Figure 8.2.182-1.
// It contains an opaque container of bridge management information.
//                                         Bits
//  Octets        8       7       6       5       4       3       2       1
//            +----------------------------------------------------------------+
//  1 to 2    |                   Type = 266 (decimal)                         |
//            |----------------------------------------------------------------|
//  3 to 4    |                      Length = n                                |
//            |----------------------------------------------------------------|
// 5 to (n+4) |                 Bridge Management Information                  |
//            +----------------------------------------------------------------+
//
// The Bridge Management Information field shall be encoded as an Octet String.
// It shall encode a TSN Bridge management message defined in clause 9 of 3GPP TS 24.519 [63].
//

type BridgeManagementInformationContainer struct {
	BridgeManagementInformation []byte `json:"bridgeManagementInformation,omitempty"`
}

// used to marshal/unmarshal json.
type bridgeManagementInformationRaw struct {
	BridgeManagementInformation string `json:"bridgeManagementInformation,omitempty"`
}

func (a *BridgeManagementInformationContainer) MarshalBinary() (data []byte, err error) {
	// Octet 5 to (n+4)
	data = a.BridgeManagementInformation

	return data, nil
}

func (a *BridgeManagementInformationContainer) UnmarshalBinary(data []byte) error {
	// Octet 5 to (n+4)
	a.BridgeManagementInformation = data

	return nil
}

func (a *BridgeManagementInformationContainer) MarshalJSON() ([]byte, error) {
	bridgeManagementInfo := &bridgeManagementInformationRaw{BridgeManagementInformation: string(a.BridgeManagementInformation)}
	return json.Marshal(bridgeManagementInfo)
}

func (a *BridgeManagementInformationContainer) UnmarshalJSON(b []byte) error {
	bridgeManagementInfo := &bridgeManagementInformationRaw{}
	err := json.Unmarshal(b, bridgeManagementInfo)
	if err != nil {
		return fmt.Errorf("unmarshal BridgeManagementInformation from json error: %v ", err)
	}
	a.BridgeManagementInformation = []byte(bridgeManagementInfo.BridgeManagementInformation)
	return nil
}

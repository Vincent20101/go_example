package pfcpie

import "fmt"

// Serving PLMN
// Casa Specific IE for Verizon use
//
//	Octets        8       7       6       5       4       3       2       1
//	          +-------------------------------------------------------------+
//	1         |                   Type = 1(decimal)                         |
//	          |-------------------------------------------------------------|
//	2         |                      Length = 3                             |
//	          |-------------------------------------------------------------|
//	3         |	        MCC Digit 2          |        MCC Digit 1           |
//	          |-------------------------------------------------------------|
//	3         |	        MNC Digit 3          |        MCC Digit 3           |
//	          |-------------------------------------------------------------|
//	3         |	        MNC Digit 2          |        MNC Digit 1           |
//	          |-------------------------------------------------------------|
type ServingPLMN struct {
	Mcc string `json:"mcc,omitempty"`
	Mnc string `json:"mnc,omitempty"`
}

func (f *ServingPLMN) MarshalBinary() (data []byte, err error) {
	data, err = marshalMccMnc(f.Mcc, f.Mnc)
	if err != nil {
		return nil, fmt.Errorf("ServingPLMN marshal: %v", err)
	}

	return data, nil
}

func (f *ServingPLMN) UnmarshalBinary(data []byte) (err error) {
	length := uint16(len(data))
	if length < 3 {
		return fmt.Errorf("ServingPLMN Inadequate TLV length: %d", length)
	}
	f.Mcc, f.Mnc, err = unmarshalMccMnc(data[0:3])
	if err != nil {
		return err
	}
	return nil
}

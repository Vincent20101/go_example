package pfcpie

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// [TS 29244 8.2.117 APN/DNN]
// Access Point Name (APN) / Data Network Name (DNN) is transferred from CP
// function to UP function.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                   Type = 134(decimal)                          |
//	         |----------------------------------------------------------------|
//	3 to 4   |                      Length = n                                |
//	         |----------------------------------------------------------------|
//	5 to n+4 |                        APN/DNN                                 |
//	         +----------------------------------------------------------------+
//
// The encoding the APN/DNN field follows 3GPP TS 23.003 [2] clause 9.1. The
// content of the APN/DNN field shall be  the full APN/DNN with both the
// APN/DNN Network Identifier and APN/DNN Operator Identifier being present as
// specified in 3GPP TS 23.003 [2] clauses 9.1.1 and 9.1.2, 3GPP TS 23.060 [19]
// Annex A and 3GPP TS 23.401 [14]  clauses 4.3.8.1.
//
// NOTE: The APN/DNN field is not encoded as a dotted string as commonly used
// in documentation.
type ApnDnn struct {
	ApnDnn []byte `json:"apnDnn,omitempty"`
}

// used to marshal/unmarshal json.
type apnDnnRaw struct {
	ApnDnn string `json:"apnDnn,omitempty"`
}

func (a *ApnDnn) MarshalBinary() (data []byte, err error) {
	// Octet 5 to (n+4)
	data = a.ApnDnn

	return data, nil
}

func (a *ApnDnn) UnmarshalBinary(data []byte) error {
	// Octet 5 to (n+4)
	a.ApnDnn = data

	return nil
}

func (a *ApnDnn) MarshalJSON() ([]byte, error) {
	apnDnnStr := ""
	offset := 0
	for offset < len(a.ApnDnn) {
		length := int(a.ApnDnn[offset])
		offset++
		if length < 1 || offset+length > len(a.ApnDnn) {
			return nil, errors.New("ApnDnn convert to string format error")
		}
		token := string(a.ApnDnn[offset : offset+length])
		apnDnnStr = apnDnnStr + token + "."
		offset += length
	}
	var apndnn *apnDnnRaw
	if len(apnDnnStr) > 1 {
		apndnn = &apnDnnRaw{ApnDnn: apnDnnStr[:len(apnDnnStr)-1]}
	}
	return json.Marshal(apndnn)
}

func (a *ApnDnn) UnmarshalJSON(b []byte) error {
	apndnn := &apnDnnRaw{}
	err := json.Unmarshal(b, apndnn)
	if err != nil {
		return fmt.Errorf("Unmarshal ApnDnn from json error: %v ", err)
	}
	if len(apndnn.ApnDnn) != 0 {
		token := strings.Split(apndnn.ApnDnn, ".")
		for _, s := range token {
			a.ApnDnn = append(a.ApnDnn, byte(len(s)))
			a.ApnDnn = append(a.ApnDnn, []byte(s)...)
		}
	}

	return nil
}

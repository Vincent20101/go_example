package pfcpie

import (
	"errors"
	"strings"
)

// APN Name
// Casa Specific IE for Verizon use
// sgw-c decodes this IE in the manner similar to ApnDnn. Hence ApnName is also
// being encoded in the same manner.
// Refer to ApnDnn.go for details of the implementation.
//
//	Octets        8       7       6       5       4       3       2       1
//	          +--------------------------------------------------------------+
//	1         |                   Type = 1(decimal)                          |
//	          |--------------------------------------------------------------|
//	2         |                      Length = n                              |
//	          |--------------------------------------------------------------|
//	3 to n    |                    APN Name                                  |
//	          +--------------------------------------------------------------+
//
// APN Name is ASCII strings interspread with 1 byte lenth fields.
type ApnName struct {
	ApnName string `json:"apnName,omitempty"`
}

func (a *ApnName) MarshalBinary() (data []byte, err error) {
	// Octet 5 to (n+4)
	if len(a.ApnName) != 0 {
		token := strings.Split(a.ApnName, ".")
		for _, s := range token {
			data = append(data, byte(len(s)))
			data = append(data, []byte(s)...)
		}
	}

	return data, nil
}

func (a *ApnName) UnmarshalBinary(data []byte) error {
	// Octet 5 to (n+4)
	a.ApnName = ""
	offset := 0
	dLen := len(data)
	for offset < dLen {
		length := int(data[offset])
		offset++
		if length < 1 || offset+length > dLen {
			return errors.New("ApnName convert to string format error")
		}
		token := string(data[offset : offset+length])
		if (offset + length) < dLen {
			a.ApnName = a.ApnName + token + "."
		} else {
			a.ApnName = a.ApnName + token
		}
		offset += length
	}

	return nil
}

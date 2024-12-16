package pfcpie

import "fmt"

// 8.2.146	TSN Time Domain Number
// The TSN Time Domain Number IE type shall be encoded as shown in Figure 8.2.146-1.
// It shall contain a TSN timing related Domain Number.
//
//	                                         Bits
//	Octets        8       7       6       5       4       3       2       1
//	          +----------------------------------------------------------------+
//	1 to 2    |                      Type = 206 (decimal)                      |
//	          +----------------------------------------------------------------+
//	3 to 4    |                           Length = n                           |
//	          +----------------------------------------------------------------+
//	   5      |                  TSN Time Domain Number value                  |
//	          +----------------------------------------------------------------+
//	6 to (n+4)|	These octet(s) is/are present only if explicitly specified     |
//	          +----------------------------------------------------------------+
//
// The TSN Time Domain Number value field shall be encoded as a binary integer value.
type TSNTimeDomainNumber struct {
	TSNTimeDomainNumberValue uint8 `json:"tSNTimeDomainNumberValue,omitempty"`
}

func (t *TSNTimeDomainNumber) MarshalBinary() (data []byte, err error) {
	// Octet 5
	data = append([]byte(""), byte(t.TSNTimeDomainNumberValue))

	return data, nil
}

func (t *TSNTimeDomainNumber) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("ie TSNTimeDomainNumber Inadequate TLV length: %d", length)
	}
	t.TSNTimeDomainNumberValue = uint8(data[idx])

	return nil
}

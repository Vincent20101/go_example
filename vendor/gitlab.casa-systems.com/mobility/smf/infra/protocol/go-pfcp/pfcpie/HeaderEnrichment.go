package pfcpie

import (
	"fmt"
	"math/bits"
)

const (
	HeaderTypeHTTP = 0
)

// Header Enrichment IE. [Refer to TS29244 8.2.67 Header Enrichment]
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                       Type = 98(decimal)                      |
//	         |---------------------------------------------------------------|
//	3 to 4   |                           Length = n                          |
//	         |-----------------------+---------------------------------------|
//	  5      |         Spare         |               Header Type             |
//	         |-----------------------+---------------------------------------|
//	  6      |                 Length of Header Field Name                   |
//	         |---------------------------------------------------------------|
//	7 to m   |                        Header Field Name                      |
//	         |---------------------------------------------------------------|
//	  p      |                 Length of Header Field Value                  |
//	         |---------------------------------------------------------------|
//	p to p+1 |                        Header Field Value                     |
//	         |---------------------------------------------------------------|
//	m to n+4 |   These octets(s) is/are present only if explicity specified  |
//	         +---------------------------------------------------------------+
//	                              Figure 8.2.67-1: Header Enrichment
//
// Header Type indicates the type of the Header. It shall be encoded as defined
// in Table 8.2.67-1.
//
//	           Table 8.2.67-1: Header Type
//	+------------------------------+-----------------+
//	|        Header Type           |       Value     |
//	|------------------------------+-----------------|
//	|           HTTP               |        0        |
//	|------------------------------+-----------------|
//	|     Spare, For future use.   |      1 to 13    |
//	+------------------------------+-----------------+
//
// Length of Header Field Name indicates the length of the Header Field Name.
//
// Header Field Name shall be encoded as an OctetString.
//
// Length of Header Field Value indicates the length of the Header Field Value.
//
// Header Field Value shall be encoded as an OctetString.
//
// For a HTTP Header Type, the contents of the Header Field Name and Header
// Field Value shall comply with the HTTP header field format (see clause 3.2
// of IETF RFC 7230 [23]).
type HeaderEnrichment struct {
	HeaderType       uint8  `json:"headerType,omitempty"` // 0x00011111
	HeaderFieldName  string `json:"headerFieldName,omitempty"`
	HeaderFieldValue string `json:"headerFieldValue,omitempty"`
}

func NewHeaderEnrichment() *HeaderEnrichment {
	return &HeaderEnrichment{
		HeaderType:       0,
		HeaderFieldName:  "",
		HeaderFieldValue: "",
	}
}

func (h *HeaderEnrichment) MarshalBinary() (data []byte, err error) {

	// Octet 5
	if bits.Len8(h.HeaderType) > 5 {
		return []byte(""), fmt.Errorf("Header type shall not be greater than 5 bits binary integer")
	}
	data = append([]byte(""), byte(h.HeaderType))

	// Octet 6
	lengthOfHeaderFieldName := len(h.HeaderFieldName)
	data = append(data, byte(lengthOfHeaderFieldName))

	// Octet 7 to m
	data = append(data, (h.HeaderFieldName)...)

	// Octet p
	lengthOfHeaderFieldValue := len(h.HeaderFieldValue)
	data = append(data, byte(lengthOfHeaderFieldValue))

	// Octet (p+1) to q
	data = append(data, h.HeaderFieldValue...)

	return data, nil
}

func (h *HeaderEnrichment) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Unmarshal IE HeaderEnrichment empty body")
	}
	h.HeaderType = uint8(data[idx]) & Mask5
	idx = idx + 1

	// Octet 6
	if length < idx+1 {
		return fmt.Errorf("Unmarshal IE HeaderEnrichment failed: field [length Of HeaderFieldName] Inadequate TLV length: %d", length)
	}
	lengthOfHeaderFieldName := data[idx]
	idx = idx + 1

	// Octet 7 to m
	if length < idx+uint16(lengthOfHeaderFieldName) {
		return fmt.Errorf("Unmarshal IE HeaderEnrichment failed: field [Header FieldName] Inadequate TLV length: %d", length)
	}
	h.HeaderFieldName = string(data[idx : idx+uint16(lengthOfHeaderFieldName)])
	idx = idx + uint16(lengthOfHeaderFieldName)

	// Octet p
	if length < idx+1 {
		return fmt.Errorf("Unmarshal IE HeaderEnrichment failed: Field [length Of Header Field Value] Inadequate TLV length: %d", length)
	}
	lengthOfHeaderFieldValue := data[idx]
	idx = idx + 1

	// Octet (p+1) to q
	if length < idx+uint16(lengthOfHeaderFieldValue) {
		return fmt.Errorf("Unmarshal IE HeaderEnrichment failed: Field [Header Field Value] Inadequate TLV length: %d", length)
	}
	h.HeaderFieldValue = string(data[idx : idx+uint16(lengthOfHeaderFieldValue)])
	idx = idx + uint16(lengthOfHeaderFieldValue)

	return nil
}

package pfcpie

// PGW-C/SMF Node Name
// Casa Specific IE for Verizon use
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	   1     |                     Type = 2(decimal)                         |
//	         |---------------------------------------------------------------|
//	   2     |                      Length = n                               |
//	         |---------------------------------------------------------------|
//	3 to n   |                       PGW-C Node Name                         |
//	         +---------------------------------------------------------------+
//
// PGW-C Node Name is ASCII string.
type PGW_C_SMFNodeName struct {
	PgwCNodeName string `json:"pgwCNodeName,omitempty" cmp:"ignore"`
}

func (p *PGW_C_SMFNodeName) MarshalBinary() (data []byte, err error) {
	return []byte(p.PgwCNodeName), nil
}

func (f *PGW_C_SMFNodeName) UnmarshalBinary(data []byte) error {
	f.PgwCNodeName = string(data)
	return nil
}

package pfcpie

// Charging Id
// Casa Specific IE for Verizon use
//
//	Octets        8       7       6       5       4       3       2       1
//	          +-------------------------------------------------------------+
//	1         |                   Type = 5(decimal)                         |
//	          |-------------------------------------------------------------|
//	2         |                      Length = n                             |
//	          |-------------------------------------------------------------|
//	3 to n    |	                       Charging Id                          |
//	          |-------------------------------------------------------------|
type ChargingId struct {
	ChargingIdVal []byte `json:"chargingIdVal,omitempty" cmp:"ignore"`
}

func (f *ChargingId) MarshalBinary() (data []byte, err error) {
	data = append(data, f.ChargingIdVal...)
	return data, nil
}

func (f *ChargingId) UnmarshalBinary(data []byte) error {
	length := len(data)

	f.ChargingIdVal = data[0:length]

	return nil
}

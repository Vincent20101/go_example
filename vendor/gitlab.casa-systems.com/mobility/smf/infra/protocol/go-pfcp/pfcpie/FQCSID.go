package pfcpie

type FQCSID struct {
	FQCSIDdata []byte `json:"fQCSIDdata,omitempty"`
}

func (f *FQCSID) MarshalBinary() (data []byte, err error) {
	return f.FQCSIDdata, nil
}

func (f *FQCSID) UnmarshalBinary(data []byte) error {
	f.FQCSIDdata = make([]byte, len(data))
	copy(f.FQCSIDdata, data)
	return nil
}

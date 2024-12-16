package pfcpie

type DuplicatingParameters struct {
	DuplicatingParametersdata []byte `json:"duplicatingParametersdata,omitempty"`
}

func (d *DuplicatingParameters) MarshalBinary() (data []byte, err error) {
	return d.DuplicatingParametersdata, nil
}

func (d *DuplicatingParameters) UnmarshalBinary(data []byte) error {
	d.DuplicatingParametersdata = make([]byte, len(data))
	copy(d.DuplicatingParametersdata, data)
	return nil
}

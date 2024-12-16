package pfcpie

type Reserved struct {
	Reserveddata []byte `json:"reserveddata,omitempty"`
}

func (r *Reserved) MarshalBinary() (data []byte, err error) {
	return r.Reserveddata, nil
}

func (r *Reserved) UnmarshalBinary(data []byte) error {
	r.Reserveddata = make([]byte, len(data))
	copy(r.Reserveddata, data)
	return nil
}

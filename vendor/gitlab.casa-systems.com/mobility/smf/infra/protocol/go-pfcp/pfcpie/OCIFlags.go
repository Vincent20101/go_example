package pfcpie

type OCIFlags struct {
	OCIFlagsdata []byte `json:"oCIFlagsdata,omitempty"`
}

func (o *OCIFlags) MarshalBinary() (data []byte, err error) {
	return o.OCIFlagsdata, nil
}

func (o *OCIFlags) UnmarshalBinary(data []byte) error {
	o.OCIFlagsdata = make([]byte, len(data))
	copy(o.OCIFlagsdata, data)
	return nil
}

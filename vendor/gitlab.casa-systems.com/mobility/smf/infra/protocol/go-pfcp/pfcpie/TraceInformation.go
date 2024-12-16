package pfcpie

type TraceInformation struct {
	TraceInformationdata []byte `json:"traceInformationdata,omitempty"`
}

func (t *TraceInformation) MarshalBinary() (data []byte, err error) {
	return t.TraceInformationdata, nil
}

func (t *TraceInformation) UnmarshalBinary(data []byte) error {
	t.TraceInformationdata = make([]byte, len(data))
	copy(t.TraceInformationdata, data)
	return nil
}

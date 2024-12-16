package pfcpie

type TimeQuotaMechanism struct {
	TimeQuotaMechanismdata []byte `json:"timeQuotaMechanismdata,omitempty"`
}

func (t *TimeQuotaMechanism) MarshalBinary() (data []byte, err error) {
	return t.TimeQuotaMechanismdata, nil
}

func (t *TimeQuotaMechanism) UnmarshalBinary(data []byte) error {
	t.TimeQuotaMechanismdata = make([]byte, len(data))
	copy(t.TimeQuotaMechanismdata, data)
	return nil
}

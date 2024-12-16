package pfcpie

type Timer struct {
	Timerdata []byte `json:"timerdata,omitempty"`
}

func (t *Timer) MarshalBinary() (data []byte, err error) {
	return t.Timerdata, nil
}

func (t *Timer) UnmarshalBinary(data []byte) error {
	t.Timerdata = make([]byte, len(data))
	copy(t.Timerdata, data)
	return nil
}

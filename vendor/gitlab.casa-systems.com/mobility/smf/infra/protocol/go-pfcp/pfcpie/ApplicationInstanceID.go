package pfcpie

type ApplicationInstanceID struct {
	ApplicationInstanceIDdata []byte `json:"applicationInstanceIDdata,omitempty"`
}

func (a *ApplicationInstanceID) MarshalBinary() (data []byte, err error) {
	return a.ApplicationInstanceIDdata, nil
}

func (a *ApplicationInstanceID) UnmarshalBinary(data []byte) error {
	a.ApplicationInstanceIDdata = make([]byte, len(data))
	copy(a.ApplicationInstanceIDdata, data)
	return nil
}

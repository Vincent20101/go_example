package pfcpie

type UpdateForwardingParameters struct {
	UpdateForwardingParametersdata []byte `json:"updateForwardingParametersdata,omitempty"`
}

func (u *UpdateForwardingParameters) MarshalBinary() (data []byte, err error) {
	return u.UpdateForwardingParametersdata, nil
}

func (u *UpdateForwardingParameters) UnmarshalBinary(data []byte) error {
	u.UpdateForwardingParametersdata = make([]byte, len(data))
	copy(u.UpdateForwardingParametersdata, data)
	return nil
}

package pfcpie

type UpdateDuplicatingParameters struct {
	UpdateDuplicatingParametersdata []byte `json:"updateDuplicatingParametersdata,omitempty"`
}

func (u *UpdateDuplicatingParameters) MarshalBinary() (data []byte, err error) {
	return u.UpdateDuplicatingParametersdata, nil
}

func (u *UpdateDuplicatingParameters) UnmarshalBinary(data []byte) error {
	u.UpdateDuplicatingParametersdata = make([]byte, len(data))
	copy(u.UpdateDuplicatingParametersdata, data)
	return nil
}

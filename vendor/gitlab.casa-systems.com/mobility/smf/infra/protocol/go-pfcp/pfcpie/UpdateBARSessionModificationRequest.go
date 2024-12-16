package pfcpie

type UpdateBARSessionModificationRequest struct {
	UpdateBARSessionModificationRequestdata []byte `json:"updateBARSessionModificationRequestdata,omitempty"`
}

func (u *UpdateBARSessionModificationRequest) MarshalBinary() (data []byte, err error) {
	return u.UpdateBARSessionModificationRequestdata, nil
}

func (u *UpdateBARSessionModificationRequest) UnmarshalBinary(data []byte) error {
	u.UpdateBARSessionModificationRequestdata = make([]byte, len(data))
	copy(u.UpdateBARSessionModificationRequestdata, data)
	return nil
}

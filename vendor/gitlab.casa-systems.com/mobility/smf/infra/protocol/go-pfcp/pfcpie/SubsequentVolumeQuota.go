package pfcpie

type SubsequentVolumeQuota struct {
	SubsequentVolumeQuotadata []byte `json:"subsequentVolumeQuotadata,omitempty"`
}

func (s *SubsequentVolumeQuota) MarshalBinary() (data []byte, err error) {
	return s.SubsequentVolumeQuotadata, nil
}

func (s *SubsequentVolumeQuota) UnmarshalBinary(data []byte) error {
	s.SubsequentVolumeQuotadata = make([]byte, len(data))
	copy(s.SubsequentVolumeQuotadata, data)
	return nil
}

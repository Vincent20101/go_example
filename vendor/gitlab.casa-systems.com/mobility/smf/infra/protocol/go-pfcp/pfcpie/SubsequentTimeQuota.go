package pfcpie

type SubsequentTimeQuota struct {
	SubsequentTimeQuotadata []byte `json:"subsequentTimeQuotadata,omitempty"`
}

func (s *SubsequentTimeQuota) MarshalBinary() (data []byte, err error) {
	return s.SubsequentTimeQuotadata, nil
}

func (s *SubsequentTimeQuota) UnmarshalBinary(data []byte) error {
	s.SubsequentTimeQuotadata = make([]byte, len(data))
	copy(s.SubsequentTimeQuotadata, data)
	return nil
}

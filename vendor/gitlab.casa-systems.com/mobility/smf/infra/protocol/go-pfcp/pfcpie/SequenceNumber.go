package pfcpie

type SequenceNumber struct {
	SequenceNumberdata []byte `json:"sequenceNumberdata,omitempty"`
}

func (s *SequenceNumber) MarshalBinary() (data []byte, err error) {
	return s.SequenceNumberdata, nil
}

func (s *SequenceNumber) UnmarshalBinary(data []byte) error {
	s.SequenceNumberdata = make([]byte, len(data))
	copy(s.SequenceNumberdata, data)
	return nil
}

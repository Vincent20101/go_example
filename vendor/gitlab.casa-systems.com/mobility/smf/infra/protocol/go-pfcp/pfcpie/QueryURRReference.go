package pfcpie

type QueryURRReference struct {
	QueryURRReferencedata []byte `json:"queryURRReferencedata,omitempty"`
}

func (q *QueryURRReference) MarshalBinary() (data []byte, err error) {
	return q.QueryURRReferencedata, nil
}

func (q *QueryURRReference) UnmarshalBinary(data []byte) error {
	q.QueryURRReferencedata = make([]byte, len(data))
	copy(q.QueryURRReferencedata, data)
	return nil
}

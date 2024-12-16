package pfcpie

type Metric struct {
	Metricdata []byte `json:"metricdata,omitempty"`
}

func (m *Metric) MarshalBinary() (data []byte, err error) {
	return m.Metricdata, nil
}

func (m *Metric) UnmarshalBinary(data []byte) error {
	m.Metricdata = make([]byte, len(data))
	copy(m.Metricdata, data)
	return nil
}

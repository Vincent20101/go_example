package pfcpie

// TODO

type GTPUPathQoSReport struct {
}

func (g *GTPUPathQoSReport) MarshalBinary() (data []byte, err error) {
	return []byte{}, nil
}

func (g *GTPUPathQoSReport) UnmarshalBinary(data []byte) error {
	return nil
}

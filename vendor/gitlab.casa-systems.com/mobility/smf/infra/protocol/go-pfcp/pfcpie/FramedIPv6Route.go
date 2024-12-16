package pfcpie

type FramedIPv6Route struct {
	FramedIPv6Routedata []byte `json:"framedIPv6Routedata,omitempty"`
}

func (f *FramedIPv6Route) MarshalBinary() (data []byte, err error) {
	return f.FramedIPv6Routedata, nil
}

func (f *FramedIPv6Route) UnmarshalBinary(data []byte) error {
	f.FramedIPv6Routedata = make([]byte, len(data))
	copy(f.FramedIPv6Routedata, data)
	return nil
}

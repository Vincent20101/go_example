package pfcpie

type FramedRoute struct {
	FramedRoutedata []byte `json:"framedRoutedata,omitempty"`
}

func (f *FramedRoute) MarshalBinary() (data []byte, err error) {
	return f.FramedRoutedata, nil
}

func (f *FramedRoute) UnmarshalBinary(data []byte) error {
	f.FramedRoutedata = make([]byte, len(data))
	copy(f.FramedRoutedata, data)
	return nil
}

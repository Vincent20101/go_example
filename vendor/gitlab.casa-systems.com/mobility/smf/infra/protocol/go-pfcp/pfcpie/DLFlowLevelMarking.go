package pfcpie

type DLFlowLevelMarking struct {
	DLFlowLevelMarkingdata []byte `json:"dLFlowLevelMarkingdata,omitempty"`
}

func (d *DLFlowLevelMarking) MarshalBinary() (data []byte, err error) {
	return d.DLFlowLevelMarkingdata, nil
}

func (d *DLFlowLevelMarking) UnmarshalBinary(data []byte) error {
	d.DLFlowLevelMarkingdata = make([]byte, len(data))
	copy(d.DLFlowLevelMarkingdata, data)
	return nil
}

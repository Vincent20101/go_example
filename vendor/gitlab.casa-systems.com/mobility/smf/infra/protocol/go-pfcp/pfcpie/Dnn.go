package pfcpie

type Dnn []uint8

func (d *Dnn) MarshalBinary() (data []byte, err error) {

	//data = append(data, uint8(len(*d)))
	data = *d

	return data, nil
}

func (d *Dnn) UnmarshalBinary(data []byte) error {

	*d = data
	return nil
}

func (d *Dnn) ToString() (dnnString string) {
	return string(*d)
}

func DnnFromString(dnnString string) (d Dnn) {
	return []byte(dnnString)
}

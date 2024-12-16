package pfcpie

type AdditionalUsageReportsInformation struct {
	AdditionalUsageReportsInformationdata []byte `json:"additionalUsageReportsInformationdata,omitempty"`
}

func (a *AdditionalUsageReportsInformation) MarshalBinary() (data []byte, err error) {
	return a.AdditionalUsageReportsInformationdata, nil
}

func (a *AdditionalUsageReportsInformation) UnmarshalBinary(data []byte) error {
	a.AdditionalUsageReportsInformationdata = make([]byte, len(data))
	copy(a.AdditionalUsageReportsInformationdata, data)
	return nil
}

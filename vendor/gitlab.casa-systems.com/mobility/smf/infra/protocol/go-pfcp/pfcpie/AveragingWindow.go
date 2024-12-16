package pfcpie

import "fmt"

// AveragingWindow implements Averaging Window IE.
// IE Type: 157
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +----------------------------------------------------------------+
//	1 to 2   |                        Type = 157(decimal)                     |
//	         |----------------------------------------------------------------|
//	3 to 4   |                        Length = n                              |
//	         |----------------------------------------------------------------|
//	5 to 8   |                        Averaging Window                        |
//	         |----------------------------------------------------------------|
//	9 to(n+4)| This octet(s) is/are present only if explicitly specified      |
//	         +----------------------------------------------------------------+
type AveragingWindow struct {
	AveragingWindow uint32 `json:"averagingWindow,omitempty"`
}

// NewAveragingWindow return a new AveragingWindow.
func NewAveragingWindow() *AveragingWindow {
	return &AveragingWindow{
		AveragingWindow: 0,
	}
}

// MarshalBinary marshals AveragingWindow into octets.
func (v *AveragingWindow) MarshalBinary() (data []byte, err error) {
	data = AppendUint32(data, v.AveragingWindow)
	return data, nil
}

// UnmarshalBinary decodes octets to AveragingWindow.
func (v *AveragingWindow) UnmarshalBinary(data []byte) error {
	byteReader := NewByteReader(data)

	// AveragingWindow
	AveragingWindow, e := byteReader.ReadUint32()
	if e != nil {
		return fmt.Errorf("inadequate length for AveragingWindow: %v", e)
	}
	v.AveragingWindow = AveragingWindow
	return nil
}

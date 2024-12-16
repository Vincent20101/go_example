package pfcpie

import (
	"errors"
	"fmt"
)

// Refer to [TS29244-g40 8.2.145 Requested Clock Drift Information]
// The Requested Clock Drift Information IE shall be encoded as shown in Figure 8.2.145-1. It indicates the clock drift information to report to the CP function.
//
//	                                        Bits
//	Octets        8       7       6       5       4       3       2      1
//	          +---------------------------------------------------------------+
//	1 to 2    |                 Type = 204 (decimal)                          |
//	          |---------------------------------------------------------------|
//	3 to 4    |                            Length = n                         |
//	          |-------+-------+-------+-------+-------+-------+-------+-------|
//	5         | Spare | Spare | Spare | Spare | Spare | Spare | RRCR  | RRTO  |
//	          |-------+-------+-------+-------+-------+-------+-------+-------|
//	6 to (n+4)| These octet(s) is/are present only if explicitly specified    |
//	          +---------------------------------------------------------------+
//	                 Figure 8.2.145-1: Requested Clock Drift Information
//
// Octet 5 shall be encoded as follows:
//
//   - Bit 1: (RRTO) Request to Report Time Offset: when set to "1", this
//     indicates a request to report when the Time Offset Reporting Threshold
//     is exceeded.
//
//   - Bit 2: (RRCR) Request to Report Cumulative RateRatio: when set to "1",
//     this indicates a request to report when the cumulative RateRatio
//     Reporting Thresholds is exceeded.
//
//   - Bits 3 to 8: Spare, for future use and set to "0".
//
// At least one bit shall be set to "1". Several bits may be set to "1".
type RequestedClockDriftInformation struct {
	RRCR bool `json:"rRCR,omitempty"`
	RRTO bool `json:"rRTO,omitempty"`
}

func (c *RequestedClockDriftInformation) MarshalBinary() (data []byte, err error) {
	if !c.RRCR && !c.RRTO {
		return nil, errors.New("At least one bit shall be set to \"1\". Several bits may be set to \"1\"")
	}
	tmpUint8 := btou(c.RRCR)<<1 |
		btou(c.RRTO)

	data = append([]byte(""), byte(tmpUint8))
	return data, nil
}

func (c *RequestedClockDriftInformation) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Unmarshal RequestedClockDriftInformation err, inadequate TLV length: %d ", length)
	}

	c.RRCR = utob(uint8(data[idx]) & BitMask2)
	c.RRTO = utob(uint8(data[idx]) & BitMask1)

	return nil
}

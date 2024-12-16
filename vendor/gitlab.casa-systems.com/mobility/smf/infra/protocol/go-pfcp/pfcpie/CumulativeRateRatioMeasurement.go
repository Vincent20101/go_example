package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// 8.2.150	Cumulative rateRatio Measurement
// The Cumulative rateRatio Measurement IE type shall be encoded as shown in Figure 8.2.150-1.
//                                         Bits
//  Octets        8       7       6       5       4       3       2       1
//            +----------------------------------------------------------------+
//  1 to 2    |                    Type = 210(decimal)                         |
//            |----------------------------------------------------------------|
//  3 to 4    |                      Length = n                                |
//            |----------------------------------------------------------------|
//  5 to 8    |               Cumulative rateRatio Measurement                 |
//            |----------------------------------------------------------------|
// 9 to (n+4) |   These octet(s) is/are present only if explicitly specified   |
//            +----------------------------------------------------------------+
//
// The Cumulative rateRatio Measurement field shall be encoded as the cumulativeRateRatio (Integer32)
// specified in clauses 14.4.2 and 15.6 of IEEE Std 802.1AS-Rev/D7.3 [58],
// i.e. the quantity "(rateRatio- 1.0)(2^41)". It shall be equal to the cumulative ratio of
// the frequency of the 5GS clock to the frequency of the clock of the TSN time domain.
//

type CumulativeRateRatioMeasurement struct {
	CumulativeRateRatioMeasurement uint32 `json:"cumulativeRateRatioMeasurement,omitempty"`
}

func (s *CumulativeRateRatioMeasurement) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, s.CumulativeRateRatioMeasurement)

	return data, nil
}

func (s *CumulativeRateRatioMeasurement) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("ie CumulativeRateRatioMeasurement Inadequate TLV length: %d", length)
	}
	s.CumulativeRateRatioMeasurement = binary.BigEndian.Uint32(data[idx:])

	return nil
}

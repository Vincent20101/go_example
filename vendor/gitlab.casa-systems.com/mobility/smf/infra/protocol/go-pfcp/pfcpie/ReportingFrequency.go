package pfcpie

import (
	"fmt"
)

// 8.2.168	Reporting Frequency
// The Reporting Frequency IE shall be encoded as shown in Figure 8.2.168-1. It indicates the frequency for the UP function to send a QoS Monitoring report to the CP function.
//
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                    Type = 244(decimal)                          |
//           |----------------------------------------------------------------|
//  3 to 4   |                      Length = n                                |
//           |----------------------------------------------------------------|
//    5      | Spare | Spare | Spare | Spare | Spare | SESRL | PERIO | EVETT  |
//           |----------------------------------------------------------------|
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//           +----------------------------------------------------------------+
//
// Octet 5 shall be encoded as follows:
// -	Bit 1  EVETT (Event Triggered QoS monitoring reporting): when set to "1", this indicates the delay for QoS monitoring is reported when the delay exceeds a threshold.
// -	Bit 2  PERIO (Periodic QoS monitoring reporting): when set to "1", this indicates the delay for QoS monitoring is reported periodicly.
// -	Bit 3  SESRL (Session Released QoS monitoring reporting): when set to "1", this indicates the delay for QoS monitoring is reported when the PDU session is released.
// -	Bits 4 to 8: Spare, for future use and set to "0".
// At least one bit shall be set to "1". Several bits may be set to "1".
//

type ReportingFrequency struct {
	SessRel   bool `json:"sessRel,omitempty"`
	Periodic  bool `json:"periodic,omitempty"`
	EventTrig bool `json:"eventTrig,omitempty"`
}

func (r *ReportingFrequency) MarshalBinary() (data []byte, err error) {
	// Octet 5
	if !r.SessRel && !r.Periodic && !r.EventTrig {
		return []byte(""), fmt.Errorf("At least one of SESRL, PERIO or EVETT shall be set for ReportingFrequency IE")
	}
	tmpUint8 := btou(r.SessRel)<<2 |
		btou(r.Periodic)<<1 |
		btou(r.EventTrig)
	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (r *ReportingFrequency) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE ReportingFrequency Inadequate TLV length: %d", length)
	}
	r.SessRel = utob(uint8(data[idx]) & BitMask3)
	r.Periodic = utob(uint8(data[idx]) & BitMask2)
	r.EventTrig = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

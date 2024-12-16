package pfcpie

import (
	"fmt"
)

// Requested Qos Monitoring
// The Requested Qos Monitoring IE shall be encoded as shown in Figure 8.2.167-1.
//
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                    Type = 243(decimal)                          |
//           |----------------------------------------------------------------|
//  3 to 4   |                      Length = n                                |
//           |----------------------------------------------------------------|
//    5      | Spare | Spare | Spare | Spare | GTPUPM |  RP  |   UL   |  DL   |
//           |----------------------------------------------------------------|
// 6 to (n+4)|   These octets(s) is/are present only if explicity specified   |
//           +----------------------------------------------------------------+
//
//  Octet 5 shall be encoded as follows:
// -	Bit 1  DL (Downlink): when set to "1", this indicates a request for measuring the downlink packet delay from the UPF (PSA) to the UE.
// -	Bit 2  UL (Uplink): when set to "1", this indicates a request for measuring the uplink packet delay from the UE to the UPF (PSA).
// -	Bit 3  RP (Round Trip): when set to "1", this indicates a request for measuring the round trip packet delay between the UPF (PSA) and UE.
// -	Bit 4  GTPUPM (GTP-U Path Monitoring):
// -	when set to "1, this indicates that the Downlink, Uplink or Round Trip delay shall be measured by using GTP-U path monitoring, i.e. by the I-UPF reporting the one-way end to end accumulated transport delay in UL GTP-U packets towards the PSA; see clause 5.33.3.3 of 3GPP TS 23.501 [28];
// -	when set to "0", this indicates that the Downlink, Uplink or Round Trip delay shall be measured by using RAN and UPF time information in GTP-U packets; see clause 5.33.3.2 of 3GPP TS 23.501 [28].
// -	Bit 5 to 8: Spare, for future use and set to "0".
// At least one bit shall be set to "1". Several bits may be set to "1".
//

type RequestedQosMonitoring struct {
	GtpuPathMon bool `json:"gtpuPathMon,omitempty"`
	RoundTrip   bool `json:"roundTrip,omitempty"`
	UL          bool `json:"uL,omitempty"`
	DL          bool `json:"dL,omitempty"`
}

func (r *RequestedQosMonitoring) MarshalBinary() (data []byte, err error) {
	// Octet 5
	if !r.GtpuPathMon && !r.RoundTrip && !r.UL && !r.DL {
		return []byte(""), fmt.Errorf("At least one of GTPUPM, RP, UL or DL bit shall be set for RequestedQosMonitoring")
	}
	tmpUint8 := btou(r.GtpuPathMon)<<3 |
		btou(r.RoundTrip)<<2 |
		btou(r.UL)<<1 |
		btou(r.DL)
	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (r *RequestedQosMonitoring) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE RequestedQosMonitoring Inadequate TLV length: %d", length)
	}
	r.GtpuPathMon = utob(uint8(data[idx]) & BitMask4)
	r.RoundTrip = utob(uint8(data[idx]) & BitMask3)
	r.UL = utob(uint8(data[idx]) & BitMask2)
	r.DL = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

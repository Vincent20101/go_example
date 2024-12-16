package pfcpie

import (
	"encoding/binary"
	"fmt"
)

//  8.2.171	QoS Monitoring Measurement
// The QoS Monitoring Measurement IE contains the packet delay monitored by the UP function. It shall be encoded as shown in Figure 8.2.171-1.
//
//                                         Bits
//  Octets       8       7       6       5       4       3       2       1
//           +----------------------------------------------------------------+
//  1 to 2   |                    Type = 248(decimal)                          |
//           |----------------------------------------------------------------|
//  3 to 4   |                      Length = n                                |
//           |----------------------------------------------------------------|
//    5      |             Spare              |  PLMF |  RP  |   UL  |   DL   |
//           |----------------------------------------------------------------|
// m to m+3  |                 Downlink packet delay                          |
//           |----------------------------------------------------------------|
// p to p+3  |                 Uplink packet delay                            |
//           |----------------------------------------------------------------|
// q to q+3  |                 Round trip packet delay                        |
//           |----------------------------------------------------------------|
// s to (n+4)|   These octets(s) is/are present only if explicity specified   |
//           +----------------------------------------------------------------+
//
// The following flags are coded within Octet 5:
// -	Bit 1  DL (Downlink): If this bit is set to "1", then the Downlink packet delay field shall be present, otherwise the Downlink packet delay field shall not be present.
// -	Bit 2  UL (Uplink): If this bit is set to "1", then the Uplink packet delay field shall be present, otherwise the Uplink packet delay field shall not be present.
// -	Bit 3  RP (Round Trip): If this bit is set to "1", then the Round trip packet delay field shall be present, otherwise the Round trip packet delay field shall not be present.
// -	Bit 4  PLMF (Packet Delay Measurement Failure): If this bit is set to "1", this indicates no timestamp is received in uplink packet for a delay exceeding the Packet Delay Thresholds or the Measurement Period.
// -	Bit 5 to 8: Spare, for future use and set to "0".
// At least one bit shall be set to "1". Several bits may be set to "1".
// The Downlink packet delay, Uplink packet delay and Round trip packet delay fields shall be encoded as an Unsigned32 binary integer value. They shall contain the downlink, uplink or round trip packet delay in milliseconds respectively.

type QosMonitoringMeasurement struct {
	Plmf       bool   `json:"plmf,omitempty"`
	Rp         bool   `json:"rp,omitempty"`
	Ul         bool   `json:"ul,omitempty"`
	Dl         bool   `json:"dl,omitempty"`
	DlPktDelay uint32 `json:"dlPktDelay,omitempty"`
	UlPktDelay uint32 `json:"ulPktDelay,omitempty"`
	RpPktDelay uint32 `json:"rpPktDelay,omitempty"`
}

func (q *QosMonitoringMeasurement) MarshalBinary() (data []byte, err error) {
	var idx uint16 = 0
	// Octet 5
	tmpUint8 := btou(q.Plmf)<<3 |
		btou(q.Rp)<<2 |
		btou(q.Ul)<<1 |
		btou(q.Dl)
	data = append([]byte(""), byte(tmpUint8))
	idx = idx + 1

	// Octet m to (m+3)
	if q.Dl {
		data = append(data, make([]byte, 4)...)
		binary.BigEndian.PutUint32(data[idx:], q.DlPktDelay)
		idx = idx + 4
	}

	// Octet p to (p+3)
	if q.Ul {
		data = append(data, make([]byte, 4)...)
		binary.BigEndian.PutUint32(data[idx:], q.UlPktDelay)
		idx = idx + 4
	}

	// Octet q to (q+3)
	if q.Rp {
		data = append(data, make([]byte, 4)...)
		binary.BigEndian.PutUint32(data[idx:], q.RpPktDelay)
	}

	return data, nil
}

func (q *QosMonitoringMeasurement) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE QosMonitoringMeasurement Inadequate TLV length: %d", length)
	}

	q.Plmf = utob(uint8(data[idx]) & BitMask4)
	q.Rp = utob(uint8(data[idx]) & BitMask3)
	q.Ul = utob(uint8(data[idx]) & BitMask2)
	q.Dl = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet m to (m+3)
	if q.Dl {
		if length < idx+4 {
			return fmt.Errorf("IE QosMonitoringMeasurement.DlPktDelayh  Inadequate TLV length: %d", length)
		}
		q.DlPktDelay = binary.BigEndian.Uint32(data[idx:])
		idx = idx + 4
	}

	// Octet p to (p+3)
	if q.Ul {
		if length < idx+4 {
			return fmt.Errorf("IE QosMonitoringMeasurement.UlPktDelay Inadequate TLV length: %d", length)
		}
		q.UlPktDelay = binary.BigEndian.Uint32(data[idx:])
		idx = idx + 4
	}

	// Octet q to (q+3)
	if q.Rp {
		if length < idx+4 {
			return fmt.Errorf("IE QosMonitoringMeasurement.RpPktDelay Inadequate TLV length: %d", length)
		}
		q.RpPktDelay = binary.BigEndian.Uint32(data[idx:])
		idx = idx + 4
	}

	return nil
}

package pfcpie

import (
	"fmt"
)

// The Node Report Type IE. Refer to [TS29244-g40 8.2.69 Node Report Type].
// The Node Report Type IE shall be encoded as shown in Figure 8.2.69-1. It indicates the type of the node report the UP function sends to the CP function.

//	    Bits
//	Octets       8       7       6       5       4      3       2    1
//	          +---------------------------------------------------------------+
//	1 to 2    |                      Type = 101 (decimal)                     |
//	          |---------------------------------------------------------------|
//	3 to 4    |                            Length = n                         |
//	          |-------+-------+-------+-------+-------+-------+-------+-------|
//	5         | Spare | Spare | Spare | Spare | GPQR  | CKDR  | UPRR  |  UPFR |
//	          |-------+-------+-------+-------+-------+-------+-------+-------|
//	6 to (n+4)| These octet(s) is/are present only if explicitly specified    |
//	          +---------------------------------------------------------------+
//
// Figure 8.2.69-1: Node Report Type
// Octet 5 shall be encoded as follows:
//   - Bit 1 – UPFR (User Plane Path Failure Report): when set to "1", this indicates a User Plane Path Failure Report.
//   - Bit 2 – UPRR (User Plane Path Recovery Report): when set to "1", this indicates a User Plane Path Recovery Report.
//     Bit 3 – CKDR (Clock Drift Report): when set to "1", this indicates a Clock Drift Report.
//   - Bit 4 – GPQR (GTP-U Path QoS Report): when set to "1", this indicates a GTP-U Path QoS Report.
//   - Bit 5 to 8 – Spare, for future use and set to "0".
//
// At least one bit shall be set to "1". Several bits may be set to "1".
// NOTE: If both UPFR and UPRR bits are set to "1", the Remote GTP-U Peer IEs in the User Plane Path Failure Report IE and in the User Plane Path Recovery Report IE are different.
type NodeReportType struct {
	GPQR bool `json:"gPQR,omitempty"`
	CKDR bool `json:"cKDR,omitempty"`
	UPRR bool `json:"uPRR,omitempty"`
	UPFR bool `json:"uPFR,omitempty"`
}

func (n *NodeReportType) MarshalBinary() (data []byte, err error) {
	tmpUint8 := btou(n.GPQR)<<3 |
		btou(n.CKDR)<<2 |
		btou(n.UPRR)<<1 |
		btou(n.UPFR)

	data = append([]byte(""), byte(tmpUint8))

	return data, nil
}

func (n *NodeReportType) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE NodeReportType Unmarshal NodeReportType err, inadequate TLV length: %d", length)
	}
	n.GPQR = utob(uint8(data[idx]) & BitMask4)
	n.CKDR = utob(uint8(data[idx]) & BitMask3)
	n.UPRR = utob(uint8(data[idx]) & BitMask2)
	n.UPFR = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

package pfcpie

//type UsageReportTrigger struct {
//UsageReportTriggerdata []byte
//}

import (
	"fmt"
)

// Usage Report Trigger IE. Refer to [TS29244-g40 8.2.41 Usage Report Trigger].
// It indicates the trigger of the usage report.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                    Type = 63(decimal)                         |
//	         |---------------------------------------------------------------|
//	3 to 4   |                      Length = n                               |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//	  5      | IMMER | DROTH | STOPT | START | QUHTI | TIMTH | VOLTH | PERIO |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//	  6      | EVETH | MACAR | EVECL | MONIT | TERMR | LIUSA | TIMQU | VOLQU |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//	  7      | Spare | Spare | Spare | EMRRE | QUVTI | IPMJL | TEBUR | EVEQU |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//
// 8 to (n+4)|   These octets(s) is/are present only if explicity specified  |
//
//	+---------------------------------------------------------------+
//
// Octet 5 shall be encoded as follows:
//
//   - Bit 1 – PERIO (Periodic Reporting): when set to "1", this indicates a
//     periodic report.
//
//   - Bit 2 – VOLTH (Volume Threshold): when set to "1", this indicates that
//     the data volume usage reaches a volume threshold.
//
//   - Bit 3 – TIMTH (Time Threshold): when set to "1", this indicates that the
//     time usage reaches a time threshold.
//
//   - Bit 4 – QUHTI (Quota Holding Time): when set to "1", this indicates that
//     no packets have been received for a period exceeding the Quota Holding
//     Time.
//
//   - Bit 5 – START (Start of Traffic): when set to "1", this indicates that
//     the start of traffic is detected.
//
//   - Bit 6 – STOPT (Stop of Traffic): when set to "1", this indicates that
//     the stop of traffic is detected.
//
//   - Bit 7 - DROTH (Dropped DL Traffic Threshold): when set to "1", this
//     indicates that the DL traffic being dropped reaches a threshold.
//
//   - Bit 8 – IMMER (Immediate Report): when set to "1", this indicates an
//     immediate report reported on CP function demand.
//
// Octet 6 shall be encoded as follows:
//
//   - Bit 1 – VOLQU (Volume Quota): when set to "1", this indicates that the
//     Volume Quota has been exhausted.
//
//   - Bit 2 – TIMQU (Time Quota): when set to "1", this indicates that the
//     Time Quota has been exhausted.
//
//   - Bit 3 - LIUSA (Linked Usage Reporting): when set to "1", this indicates
//     a linked usage report, i.e. a usage report being reported for a URR due to
//     a usage report being also reported for a linked URR (see clause 5.2.2.4).
//
//   - Bit 4 – TERMR (Termination Report): when set to "1", this indicates a
//     usage report being reported (in a PFCP Session Deletion Response) for a
//     URR due to the termination of the PFCP session, or a usage report being
//     reported (in a PFCP Session Modification Response) for a URR due to the
//     removal of the URR or dissociated from the last PDR.
//
//   - Bit 5 – MONIT (Monitoring Time): when set to "1", this indicates a usage
//     report being reported for a URR due to the Monitoring Time being reached.
//
//   - Bit 6 – ENVCL (Envelope Closure): when set to "1", this indicates the
//     usage report is generated for closure of an envelope (see clause 5.2.2.3).
//
//   - Bit 7 – MACAR (MAC Addresses Reporting): when set to "1", this indicates
//     a usage report to report MAC (Ethernet) addresses used as source address
//     of frames sent UL by the UE.
//
//   - Bits 8: EVETH (Event Threshold): when set to "1", this indicates a usage
//     report is generated when an event threshold is reached.
//
// Octet 7 shall be encoded as follows:
//
//   - Bit 1 – EVEQU (Event Quota): when set to "1", this indicates that the
//     Event Quota has been exhausted.
//
//   - Bit 2 – TEBUR (Termination By UP function Report): when set to "1", this
//     indicates a usage report being reported for a URR due to the termination
//     of the PFCP session which is initiated by the UP function.
//
//   - Bit 3 – IPMJL (IP Multicast Join/Leave): when set to "1", this indicates
//     a usage report being reported for a URR due to the UPF adding or removing
//     the PDU session to/from the DL replication tree associated with an IP
//     multicast flow.
//
//   - Bit 4 – QUVTI (Quota Validity Time): when set to "1", this indicates a
//     usage report being reported for a URR due to the quota validity timer
//     expiry.
//
//   - Bit 5 – EMRRE (End Marker Reception REport): this indicates that the UP
//     function has received End Marker from the old I-UPF. See clauses 4.2.3.2
//     and 4.23.4.3 in 3GPP TS 23.502 [29])
//
// - Bit 6 to 8: Spare, for future use and set to "0".
//
// At least one bit shall be set to "1". Several bits may be set to "1".
type UsageReportTrigger struct {
	// Octet 5
	Immer bool `json:"immer,omitempty"`
	Droth bool `json:"droth,omitempty"`
	Stopt bool `json:"stopt,omitempty"`
	Start bool `json:"start,omitempty"`
	Quhti bool `json:"quhti,omitempty"`
	Timth bool `json:"timth,omitempty"`
	Volth bool `json:"volth,omitempty"`
	Perio bool `json:"perio,omitempty"`

	// Octet 6
	Eveth bool `json:"eveth,omitempty"`
	Macar bool `json:"macar,omitempty"`
	Envcl bool `json:"envcl,omitempty"`
	Monit bool `json:"monit,omitempty"`
	Termr bool `json:"termr,omitempty"`
	Liusa bool `json:"liusa,omitempty"`
	Timqu bool `json:"timqu,omitempty"`
	Volqu bool `json:"volqu,omitempty"`

	// Octet 7
	// TODO: Unused, need upf support.
	Emrre bool `json:"emrre,omitempty"`
	Quvti bool `json:"quvti,omitempty"`
	Ipmjl bool `json:"ipmjl,omitempty"`
	Tebur bool `json:"tebur,omitempty"`
	Evequ bool `json:"evequ,omitempty"`
}

func (u *UsageReportTrigger) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(u.Immer)<<7 |
		btou(u.Droth)<<6 |
		btou(u.Stopt)<<5 |
		btou(u.Start)<<4 |
		btou(u.Quhti)<<3 |
		btou(u.Timth)<<2 |
		btou(u.Volth)<<1 |
		btou(u.Perio)

	data = append([]byte(""), byte(tmpUint8))

	// Octet 6
	tmpUint8 = btou(u.Eveth)<<7 |
		btou(u.Macar)<<6 |
		btou(u.Envcl)<<5 |
		btou(u.Monit)<<4 |
		btou(u.Termr)<<3 |
		btou(u.Liusa)<<2 |
		btou(u.Timqu)<<1 |
		btou(u.Volqu)<<0
	data = append(data, byte(tmpUint8))

	// Octet 7
	tmpUint8 = btou(u.Emrre)<<4 |
		btou(u.Quvti)<<3 |
		btou(u.Ipmjl)<<2 |
		btou(u.Tebur)<<1 |
		btou(u.Evequ)
	data = append(data, byte(tmpUint8))

	return data, nil
}

func (r *UsageReportTrigger) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Unmarshal UsageReportTrigger err, empty body: %d", length)
	}

	r.Immer = utob(uint8(data[idx]) & BitMask8)
	r.Droth = utob(uint8(data[idx]) & BitMask7)
	r.Stopt = utob(uint8(data[idx]) & BitMask6)
	r.Start = utob(uint8(data[idx]) & BitMask5)
	r.Quhti = utob(uint8(data[idx]) & BitMask4)
	r.Timth = utob(uint8(data[idx]) & BitMask3)
	r.Volth = utob(uint8(data[idx]) & BitMask2)
	r.Perio = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	if length < idx+1 {
		return nil
	}

	r.Eveth = utob(uint8(data[idx]) & BitMask8)
	r.Macar = utob(uint8(data[idx]) & BitMask7)
	r.Envcl = utob(uint8(data[idx]) & BitMask6)
	r.Monit = utob(uint8(data[idx]) & BitMask5)
	r.Termr = utob(uint8(data[idx]) & BitMask4)
	r.Liusa = utob(uint8(data[idx]) & BitMask3)
	r.Timqu = utob(uint8(data[idx]) & BitMask2)
	r.Volqu = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// upf may use old spec which not contain octet 7, return nil here.
	if length < idx+1 {
		//fmt.Errorf("Unmarshal UsageReportTrigger Octet failed, inadequate TLV length: %d", length)
		return nil
	}

	r.Emrre = utob(uint8(data[idx]) & BitMask5)
	r.Quvti = utob(uint8(data[idx]) & BitMask4)
	r.Ipmjl = utob(uint8(data[idx]) & BitMask3)
	r.Tebur = utob(uint8(data[idx]) & BitMask2)
	r.Evequ = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	return nil
}

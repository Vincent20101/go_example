package pfcpie

import (
	"fmt"
)

// ReportingTriggers IE. Refer to [TS29244 8.2.19 Reporting Triggers].
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                   Type = 134(decimal)                         |
//	         |---------------------------------------------------------------|
//	3 to 4   |                      Length = n                               |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//	  5      | LIUSA | DROTH | STOPT | START | QUHTI | TIMTH | VOLTH | PERIO |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//	  6      | QUVTI | IPMJL | EVEQU | EVETH | MACAR | ENVCL | TIMQU | VOLQU |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//	  7      | Spare | Spare | Spare | Spare | Spare | Spare | Spare | REEMR |
//	         |-------+-------+-------+-------+-------+-------+-------+-------|
//
// 8 to (n+4)|   These octets(s) is/are present only if explicity specified  |
//
//	+---------------------------------------------------------------+
//	                 Figure 8.2.19-1: Reporting Triggers
//
// ...
//
// At least one bit shall be set to "1". Several bits may be set to "1".
type ReportingTriggers struct {

	// Octet 5

	Perio bool `json:"perio,omitempty"` // Octet5/bit1
	Volth bool `json:"volth,omitempty"` // Octet5/bit2
	Timth bool `json:"timth,omitempty"` // Octet5/bit3
	Quhti bool `json:"quhti,omitempty"` // Octet5/bit4
	Start bool `json:"start,omitempty"` // Octet5/bit5
	Stopt bool `json:"stopt,omitempty"` // Octet5/bit6
	Droth bool `json:"droth,omitempty"` // Octet5/bit7
	Liusa bool `json:"liusa,omitempty"` // Octet5/bit8

	// Octet 6

	Volqu bool `json:"volqu,omitempty"` // Octet6/bit1
	Timqu bool `json:"timqu,omitempty"` // Octet6/bit2
	Envcl bool `json:"envcl,omitempty"` // Octet6/bit3
	Macar bool `json:"macar,omitempty"` // Octet6/bit4
	Eveth bool `json:"eveth,omitempty"` // Octet6/bit5
	Evequ bool `json:"evequ,omitempty"` // Octet6/bit6
	Ipmjl bool `json:"ipmjl,omitempty"` // Octet6/bit7
	Quvti bool `json:"quvti,omitempty"` // Octet6/bit8

	// Octet 7

	Reemr bool `json:"reemr,omitempty"` // Octet7/bit1
}

func (r *ReportingTriggers) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := btou(r.Liusa)<<7 |
		btou(r.Droth)<<6 |
		btou(r.Stopt)<<5 |
		btou(r.Start)<<4 |
		btou(r.Quhti)<<3 |
		btou(r.Timth)<<2 |
		btou(r.Volth)<<1 |
		btou(r.Perio)
	data = append([]byte(""), byte(tmpUint8))

	// Octet 6
	tmpUint8 = btou(r.Quvti)<<7 |
		btou(r.Ipmjl)<<6 |
		btou(r.Evequ)<<5 |
		btou(r.Eveth)<<4 |
		btou(r.Macar)<<3 |
		btou(r.Envcl)<<2 |
		btou(r.Timqu)<<1 |
		btou(r.Volqu)
	data = append(data, byte(tmpUint8))

	// Octet 7
	tmpUint8 = btou(r.Reemr)
	data = append(data, byte(tmpUint8))

	return data, nil
}

func (r *ReportingTriggers) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("UnmarshalBinary ReportingTriggers octet5 failed: empty body")
	}
	r.Liusa = utob(uint8(data[idx]) & BitMask8)
	r.Droth = utob(uint8(data[idx]) & BitMask7)
	r.Stopt = utob(uint8(data[idx]) & BitMask6)
	r.Start = utob(uint8(data[idx]) & BitMask5)
	r.Quhti = utob(uint8(data[idx]) & BitMask4)
	r.Timth = utob(uint8(data[idx]) & BitMask3)
	r.Volth = utob(uint8(data[idx]) & BitMask2)
	r.Perio = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	// Octet 6
	if length < idx+1 {
		return fmt.Errorf("UnmarshalBinary ReportingTriggers octet6 failed: inadequate length:%d", length)
	}
	r.Quvti = utob(uint8(data[idx]) & BitMask8)
	r.Ipmjl = utob(uint8(data[idx]) & BitMask7)
	r.Evequ = utob(uint8(data[idx]) & BitMask6)
	r.Eveth = utob(uint8(data[idx]) & BitMask5)
	r.Macar = utob(uint8(data[idx]) & BitMask4)
	r.Envcl = utob(uint8(data[idx]) & BitMask3)
	r.Timqu = utob(uint8(data[idx]) & BitMask2)
	r.Volqu = utob(uint8(data[idx]) & BitMask1)
	idx = idx + 1

	if length < idx+1 {
		//return fmt.Errorf("UnmarshalBinary ReportingTriggers octet7 failed: inadequate length:%d", length)
		// TODO: chenglong: should return error here, to support old version, we return nil for now.
		return nil
	}
	r.Reemr = utob(uint8(data[idx]) & BitMask1)

	return nil
}

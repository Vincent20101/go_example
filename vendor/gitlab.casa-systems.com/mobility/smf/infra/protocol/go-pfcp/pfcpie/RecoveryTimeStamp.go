package pfcpie

import (
	"encoding/binary"
	"fmt"
)

// TS29.244-g40 8.2.65	Recovery Time Stamp
// The Recovery Time Stamp IE is coded as shown in Figure 8.2.65-1. It indicates the UTC time when the PFCP entity started. Octets 5 to 8 are encoded in the same format as the first four octets of the 64-bit timestamp format as defined in clause 6 of IETF RFC 5905 [26].
// NOTE:	The encoding is defined as the time in seconds relative to 00:00:00 on 1 January 1900.

// 		                                      Bits
// 	Octets	      8	     7	       6	    5	     4	    3	    2	    1
//            +----------------------------------------------------------------+
// 	1 to 2	  |                       Type = 96 (decimal)                      |
//            +----------------------------------------------------------------+
// 	3 to 4	  |                           Length = n                           |
//            +----------------------------------------------------------------+
// 	5 to 8	  |                    Recovery Time Stamp value                   |
//            +----------------------------------------------------------------+
// 	9 to (n+4)|	These octet(s) is/are present only if explicitly specified     |
//            +----------------------------------------------------------------+

type RecoveryTimeStamp struct {
	RecoveryTimeStamp uint32 `json:"recoveryTimeStamp,omitempty"`
}

func (r *RecoveryTimeStamp) MarshalBinary() (data []byte, err error) {
	// Octet 5 to 8
	data = make([]byte, 4)
	binary.BigEndian.PutUint32(data, r.RecoveryTimeStamp)

	return data, nil
}

func (r *RecoveryTimeStamp) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5 to 8
	if length < idx+4 {
		return fmt.Errorf("IE RecoveryTimeStamp Inadequate TLV length: %d", length)
	}
	r.RecoveryTimeStamp = binary.BigEndian.Uint32(data[idx:])
	idx = idx + 4

	return nil
}

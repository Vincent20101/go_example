package pfcpie

import (
	"encoding/binary"
	"fmt"
	"strconv"
)

// User Location Info
// Casa Specific IE for Verizon use
//
//	                                      Bits Octets
//	             8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	   1     |                     Type = 5(decimal)                         |
//	         |---------------------------------------------------------------|
//	   2     |                      Length = n                               |
//	         |------+--------+-------+-------+-------+-------+-------+-------|
//	   3     | Spare|  NCGI  |  LAI  |  ECGI |  TAI  |  RAI  |  SAI  |  CGI  |
//	         |------+--------+-------+-------+-------+-------+-------+-------|
//	a to a+6 |                           CGI Field                           |
//	         |---------------------------------------------------------------|
//	b to b+6 |                           SAI Field                           |
//	         |---------------------------------------------------------------|
//	c to c+6 |                           RAI Field                           |
//	         |---------------------------------------------------------------|
//	d to d+6 |                           TAI Field                           |
//	         |---------------------------------------------------------------|
//	e to e+6 |                          ECGI Field                           |
//	         |---------------------------------------------------------------|
//	f to f+4 |                           LAI Field                           |
//	         |---------------------------------------------------------------|
//	g to g+7 |                          NCGI Field                           |
//	         +---------------------------------------------------------------+
//
// The ULI IE shall contain only one identity of the same type (e.g. more than
// one CGI cannot be included), but ULI IE may contain more than one identity
// of a different type (e.g. ECGI and TAI). The flags LAI, ECGI, TAI, RAI, SAI
// , CGI in octet 5 indicate if the corresponding type shall be present in a
// respective field or not. If one of these flags is set to "0", the
// corresponding field shall not be present at all. If more than one identity
// of different type is present, then they shall be sorted in the following
// order: CGI, SAI, RAI, TAI, ECGI, LAI, NCGI.
//
// For each identity, if an Administration decides to include only two digits
// in the MNC, then "MNC digit 3" field of corresponding location shall be
// coded as "1111"
//
// MCC and MNC of each field refer to [TS 23.003 2.2 Composition of IMSI]:
//
//   - Mobile Country Code(MCC) consisting of three digits. The MCC identifies
//     uniquely the country of domicile of the mobile subscriber;
//
//   - Mobile Network Code (MNC) consisting of two or three digits for
//     GSM/UMTS applications. The MNC identifies the home PLMN of the mobile
//     subscriber. The length of the MNC (two or three digits) depends on the
//     value of the MCC. A mixture of two and three digit MNC codes within a
//     single MCC area is not recommended and is outside the scope of this
//     specification.n
type UserLocationInfo struct {
	NCGI      bool       `json:"nCGI,omitempty"` // octet5/bit7
	LAI       bool       `json:"lAI,omitempty"`  // octet5/bit6
	ECGI      bool       `json:"eCGI,omitempty"` // octet5/bit5
	TAI       bool       `json:"tAI,omitempty"`  // octet5/bit4
	RAI       bool       `json:"rAI,omitempty"`  // octet5/bit3
	SAI       bool       `json:"sAI,omitempty"`  // octet5/bit2
	CGI       bool       `json:"cGI,omitempty"`  // octet5/bit1
	NCGIField *NCGIField `json:"nCGIField,omitempty"`
	LAIField  *LAIField  `json:"lAIField,omitempty"`
	ECGIField *ECGIField `json:"eCGIField,omitempty"`
	TAIField  *TAIField  `json:"tAIField,omitempty"`
	RAIField  *RAIField  `json:"rAIField,omitempty"`
	SAIField  *SAIField  `json:"sAIField,omitempty"`
	CGIField  *CGIField  `json:"cGIField,omitempty"`
}

const (
	UliCgiFieldLen  = 7
	UliSaiFieldLen  = 7
	UliRaiFieldLen  = 7
	UliTaiFieldLen  = 7
	UliEcgiFieldLen = 7
	UliLaiFieldLen  = 5
	UliNcgiFieldLen = 8
)

// The coding of NCGI (E-UTRAN Cell Global Identifier) is depicted in Figure
// 8.21.5-1. Only zero or one NCGI field shall be present in ULI IE.
//
//	                                        Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +-------------------------------+-------------------------------+
//	   a     |           MCC Digit 2         |          MCC Digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+1    |           MNC Digit 3         |          MCC Digit 3          |
//	         |-------------------------------+-------------------------------|
//	  a+2    |           MNC Digit 2         |          MNC digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+3    |            Spare              |              NCI              |
//	         |-------------------------------+-------------------------------|
//
// a+4 to a+7|                  		    NCI               				 |
//
//	|---------------------------------------------------------------|
//
// The NR Cell Identifier (NCI) consists of 36 bits. The NCI field shall
// start with Bit 4 of octet a+3, which is the most significant bit. Bit 1 of
// Octet a+7 is the least significant bit. The coding of the NR cell
// identifier is the responsibility of each administration. Coding using ASCII encoding
type NCGIField struct {
	Mcc string `json:"mcc,omitempty"`
	Mnc string `json:"mnc,omitempty"`
	NCI uint64 `json:"nCI,omitempty"`
}

func (e *NCGIField) MarshalBinary() (data []byte, err error) {
	data, err = marshalMccMnc(e.Mcc, e.Mnc)
	if err != nil {
		return nil, err
	}

	data = append(data, byte((e.NCI&0xf00000000)>>32))
	data = append(data, byte((e.NCI&0xff000000)>>24))
	data = append(data, byte((e.NCI&0xff0000)>>16))
	data = append(data, byte((e.NCI&0xff00)>>8))
	data = append(data, byte(e.NCI&0xff))

	return data, nil
}

func (e *NCGIField) UnmarshalBinary(data []byte) (err error) {
	length := uint16(len(data))
	if length < 8 {
		return fmt.Errorf("ECGIField Inadequate TLV length: %d", length)
	}
	e.Mcc, e.Mnc, err = unmarshalMccMnc(data[0:3])
	if err != nil {
		return err
	}
	tmp := append([]byte{0, 0, 0}, data[3:]...)
	e.NCI = binary.BigEndian.Uint64(tmp)

	return nil
}

// TS29274 8.21.1 CGI field
//
// The coding of CGI (Cell Global Identifier) is depicted in 8.21.1-1 Figure.
// Only zero or one CGI field shall be present in ULI IE.
//
//	                                        Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +-------------------------------+-------------------------------+
//	   a     |           MCC Digit 2         |          MCC Digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+1    |           MNC Digit 3         |          MCC Digit 3          |
//	         |-------------------------------+-------------------------------|
//	  a+2    |           MNC Digit 2         |          MNC digit 1          |
//	         |-------------------------------+-------------------------------|
//
// a+3 to a+4|                   Location Area Code (LAC)                    |
//
//	|---------------------------------------------------------------|
//
// a+5 to a+6|                      Cell Identity (CI)                       |
//
//			 +---------------------------------------------------------------+
//	                       Figure 8.21.1-1: CGI field
//
// The Location Area Code (LAC) consists of 2 octets. Bit 8 of Octet a+3 is the
// most significant bit and bit 1 of Octet a+4 the least significant bit. The
// coding of the location area code is the responsibility of each administration.
// Coding using full hexadecimal representation (binary, not ASCII encoding) shall
// be used.
//
// The Cell Identity (CI) consists of 2 octets. Bit 8 of Octet a+5 is the most
// significant bit and bit 1 of Octet a+6 the least significant bit. The coding
// of the cell identity is the responsibility of each administration. Coding
// using full hexadecimal representation (binary, not ASCII encoding) shall be
// used.
//
// MCC and MNC refer to [TS 23.003 2.2 Composition of IMSI]:
//
//   - Mobile Country Code(MCC) consisting of three digits. The MCC identifies
//     uniquely the country of domicile of the mobile subscriber;
//
//   - Mobile Network Code (MNC) consisting of two or three digits for
//     GSM/UMTS applications. The MNC identifies the home PLMN of the mobile
//     subscriber. The length of the MNC (two or three digits) depends on the
//     value of the MCC. A mixture of two and three digit MNC codes within a
//     single MCC area is not recommended and is outside the scope of this
//     specification.
type CGIField struct {
	Mcc string `json:"mcc,omitempty"` // 3 digits
	Mnc string `json:"mnc,omitempty"` // 2 or 3 digits
	LAC uint16 `json:"lAC,omitempty"`
	CI  uint16 `json:"cI,omitempty"`
}

func marshalMccMnc(Mcc string, Mnc string) (data []byte, err error) {
	if len(Mcc) != 3 {
		return nil, fmt.Errorf("MCC should consist of three digits, got %d.", len(Mcc))
	}
	mccVal, err := strconv.Atoi(Mcc)
	if err != nil {
		return nil, fmt.Errorf("parse mcc to int failed, err: %s, mcc: %s", err.Error(), Mcc)
	}

	// Get Digits
	mccDigit1 := uint8(mccVal / 100)
	mccDigit2 := uint8((mccVal / 10) % 10)
	mccDigit3 := uint8(mccVal % 10)

	//octet a: MCC Digit2 and MCC Digit1
	data = append(data, mccDigit2<<4|mccDigit1)

	lengthOfMnc := len(Mnc)
	if lengthOfMnc > 3 || lengthOfMnc < 2 {
		return nil, fmt.Errorf("MNC should consist of two or three digits, got %d.", lengthOfMnc)
	}
	mncVal, err := strconv.Atoi(Mnc)
	if err != nil {
		return nil, fmt.Errorf("parse mnc to int failed, err: %s, mnc: %s", err.Error(), Mnc)
	}

	var mncDigit1, mncDigit2, mncDigit3 uint8

	if lengthOfMnc == 2 {
		mncDigit1 = uint8(mncVal / 10)
		mncDigit2 = uint8(mncVal % 10)
		// If an Administration decides to include only two digits in the MNC,
		// then "MNC digit 3" field of corresponding location shall be coded as
		// "1111".
		mncDigit3 = uint8(0x0f)
	} else {
		// MNC is 3-digit
		mncDigit1 = uint8(mncVal / 100)
		mncDigit2 = uint8((mncVal / 10) % 10)
		mncDigit3 = uint8(mncVal % 10)
	}

	// Octet a+1: MNC Digit 3 (high 4 bits) and MCC Digit 3.
	data = append(data, mncDigit3<<4|mccDigit3)
	// Octet a+2: MNC Digit2 (high 4 bits) and MNC digit 1.
	data = append(data, mncDigit2<<4|mncDigit1)

	return data, nil
}

func unmarshalMccMnc(data []byte) (Mcc string, Mnc string, err error) {
	if len(data) != 3 {
		return "", "", fmt.Errorf("Cannot unmarshal Mcc/Mnc len:%v", len(data))
	}
	Mcc = fmt.Sprintf("%d%d%d", data[0]&0xf, (data[0]&0xf0)>>4, data[1]&0xf)
	if data[1]&0xf0 == 0xf0 {
		Mnc = fmt.Sprintf("%d%d", data[2]&0xf, (data[2]&0xf0)>>4)
	} else {
		Mnc = fmt.Sprintf("%d%d%d", data[2]&0xf, (data[2]&0xf0)>>4, (data[1]&0xf0)>>4)
	}
	return Mcc, Mnc, nil
}

func (c *CGIField) MarshalBinary() (data []byte, err error) {
	data, err = marshalMccMnc(c.Mcc, c.Mnc)
	if err != nil {
		return nil, err
	}
	tmp := make([]byte, 4)
	binary.BigEndian.PutUint16(tmp, c.LAC)
	binary.BigEndian.PutUint16(tmp[2:], c.CI)
	data = append(data, tmp...)

	return data, nil
}

func (c *CGIField) UnmarshalBinary(data []byte) (err error) {
	length := uint16(len(data))
	if length < 7 {
		return fmt.Errorf("CGIField Inadequate TLV length: %d", length)
	}
	c.Mcc, c.Mnc, err = unmarshalMccMnc(data[0:3])
	if err != nil {
		return err
	}
	c.LAC = binary.BigEndian.Uint16(data[3:])
	c.CI = binary.BigEndian.Uint16(data[5:])

	return nil
}

// The coding of LAI (Location Area Identifier) is depicted in Figure 8.21.6-1.
//
//	                                        Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +-------------------------------+-------------------------------+
//	   a     |           MCC Digit 2         |          MCC Digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+1    |           MNC Digit 3         |          MCC Digit 3          |
//	         |-------------------------------+-------------------------------|
//	  a+2    |           MNC Digit 2         |          MNC digit 1          |
//	         |-------------------------------+-------------------------------|
//
// a+3 to a+4|                  Location Area Code (LAC)                     |
//
//	|---------------------------------------------------------------|
//
// The Location Area Code (LAC) consists of 2 octets. Bit 8 of Octet f+3 is the
// most significant bit and bit 1 of Octet f+4 the least significant bit. The
// coding of the location area code is the responsibility of each administration.
// Coding using full hexadecimal representation (binary, not ASCII encoding)
// shall be used.
type LAIField struct {
	Mcc string `json:"mcc,omitempty"`
	Mnc string `json:"mnc,omitempty"`
	LAC uint16 `json:"lAC,omitempty"`
}

func (l *LAIField) MarshalBinary() (data []byte, err error) {
	data, err = marshalMccMnc(l.Mcc, l.Mnc)
	if err != nil {
		return nil, err
	}

	tmp := make([]byte, 2)
	binary.BigEndian.PutUint16(tmp, l.LAC)
	data = append(data, tmp...)

	return data, nil
}

func (l *LAIField) UnmarshalBinary(data []byte) (err error) {
	length := uint16(len(data))
	if length < 5 {
		return fmt.Errorf("LAIField Inadequate TLV length: %d", length)
	}

	l.Mcc, l.Mnc, err = unmarshalMccMnc(data[0:3])
	if err != nil {
		return err
	}
	l.LAC = binary.BigEndian.Uint16(data[3:])

	return nil
}

// The coding of ECGI (E-UTRAN Cell Global Identifier) is depicted in Figure
// 8.21.5-1. Only zero or one ECGI field shall be present in ULI IE.
//
//	                                        Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +-------------------------------+-------------------------------+
//	   a     |           MCC Digit 2         |          MCC Digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+1    |           MNC Digit 3         |          MCC Digit 3          |
//	         |-------------------------------+-------------------------------|
//	  a+2    |           MNC Digit 2         |          MNC digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+3    |            Spare              |              ECI              |
//	         |-------------------------------+-------------------------------|
//
// a+4 to a+6|                    ECI (E-UTRAN Cell Identifier)              |
//
//	|---------------------------------------------------------------|
//
// The E-UTRAN Cell Identifier (ECI) consists of 28 bits. The ECI field shall
// start with Bit 4 of octet a+3, which is the most significant bit. Bit 1 of
// Octet a+6 is the least significant bit. The coding of the E-UTRAN cell
// identifier is the responsibility of each administration. Coding using full
// hexadecimal representation (binary, not ASCII encoding) shall be used.
type ECGIField struct {
	Mcc string `json:"mcc,omitempty"`
	Mnc string `json:"mnc,omitempty"`
	ECI uint32 `json:"eCI,omitempty"`
}

func (e *ECGIField) MarshalBinary() (data []byte, err error) {
	data, err = marshalMccMnc(e.Mcc, e.Mnc)
	if err != nil {
		return nil, err
	}

	data = append(data, byte((e.ECI&0xf000000)>>24))
	data = append(data, byte((e.ECI&0xff0000)>>16))
	data = append(data, byte((e.ECI&0xff00)>>8))
	data = append(data, byte(e.ECI&0xff))

	return data, nil
}

func (e *ECGIField) UnmarshalBinary(data []byte) (err error) {
	length := uint16(len(data))
	if length < 7 {
		return fmt.Errorf("ECGIField Inadequate TLV length: %d", length)
	}
	e.Mcc, e.Mnc, err = unmarshalMccMnc(data[0:3])
	if err != nil {
		return err
	}
	e.ECI = binary.BigEndian.Uint32(data[3:])

	return nil
}

// The coding of TAI (Tracking Area Identity) is depicted in Figure 8.21.4-1.
// Only zero or one TAI field shall be present in ULI IE.
//
//	                                        Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +-------------------------------+-------------------------------+
//	   a     |           MCC Digit 2         |          MCC Digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+1    |           MNC Digit 3         |          MCC Digit 3          |
//	         |-------------------------------+-------------------------------|
//	  a+2    |           MNC Digit 2         |          MNC digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+3    |          				TAC Value Len          				 |
//           |-------------------------------+-------------------------------|
// a+4 to a+6|                   Tracking Area Code (TAC)                    |
//
//			 |---------------------------------------------------------------|
//
// A Tracking Area Code (TAC) consists of 2 or 3 octets (clause9.3.3.10 of 3GPP TS 38.413 [11] in hexadecimal representation)
// When the value of octet a+3 is 2, it means that the length of tac is two octets. When the value of octet a+3 is 3, it means that the length of tac is three octets.
// When the tac value length is 2, the value of octet a+4 is 0xff, bit 8 of octet a+5 is the most significant bit, and bit 1 of octet a+6 is the least significant bit;
// When the tac value length is 3, bit 8 of octet a+4 is the most significant bit, and bit 1 of octet a+6 is the least significant bit;
// Each character in the string shall take a value of "0" to "9", "a" to "f" or "A" to "F" and shall represent 4 bits.
// The most significant character representing the 4 most significant bits of the TAC shall appear first in the string, and the character representing the 4 least significant bit of the TAC shall appear last in the string.
// Examples:
// A legacy TAC 0x4305 shall be encoded as "4305".
// An extended TAC 0x63F84B shall be encoded as "63F84B"

type TAIField struct {
	Mcc string `json:"mcc,omitempty"`
	Mnc string `json:"mnc,omitempty"`
	TAC string `json:"tAC,omitempty"`
}

const (
	TacValueLen2 = 2
	TacValueLen3 = 3
)

func (t *TAIField) MarshalBinary() (data []byte, err error) {
	data, err = marshalMccMnc(t.Mcc, t.Mnc)
	if err != nil {
		return nil, err
	}
	if len(t.TAC) == 4 {
		data = append(data, byte(TacValueLen2&0xff))
		tac, _ := strconv.ParseUint(t.TAC, 16, 32)
		data = append(data, byte(0xff))
		data = append(data, byte((tac&0xff00)>>8))
		data = append(data, byte(tac&0xff))
	} else if len(t.TAC) == 6 {
		data = append(data, byte(TacValueLen3&0xff))
		tac, _ := strconv.ParseUint(t.TAC, 16, 32)
		data = append(data, byte((tac&0xff0000)>>16))
		data = append(data, byte((tac&0xff00)>>8))
		data = append(data, byte(tac&0xff))
	} else {
		return nil, fmt.Errorf("tac value len is wrong: %v", t.TAC)
	}

	return data, nil
}

func (t *TAIField) UnmarshalBinary(data []byte) (err error) {
	length := uint16(len(data))
	if length < 7 {
		return fmt.Errorf("TAIField Inadequate TLV length: %d", length)
	}
	t.Mcc, t.Mnc, err = unmarshalMccMnc(data[0:3])
	if err != nil {
		return err
	}
	tacLen := data[3]
	if tacLen == TacValueLen2 {
		tmp := append([]byte{}, data[5:]...)
		tac := binary.BigEndian.Uint16(tmp)
		t.TAC = fmt.Sprintf("%04s", strconv.FormatUint(uint64(tac), 16))
	} else if tacLen == TacValueLen3 {
		tmp := append([]byte{0}, data[4:]...)
		tac := binary.BigEndian.Uint32(tmp)
		t.TAC = fmt.Sprintf("%06s", strconv.FormatUint(uint64(tac), 16))
	} else {
		return fmt.Errorf("tac len is wrong: %v", tacLen)
	}

	return nil
}

// The coding of RAI (Routing Area Identity) is depicted in Figure 8.21.3-1.
// Only zero or one RAI field shall be present in ULI IE.
//
//	                                        Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +-------------------------------+-------------------------------+
//	   a     |           MCC Digit 2         |          MCC Digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+1    |           MNC Digit 3         |          MCC Digit 3          |
//	         |-------------------------------+-------------------------------|
//	  a+2    |           MNC Digit 2         |          MNC digit 1          |
//	         |-------------------------------+-------------------------------|
//
// a+3 to a+4|                   Location Area Code (LAC)                    |
//
//	|---------------------------------------------------------------|
//
// a+5 to a+6|                    Routing Area Code (RAC)                    |
//
//	+---------------------------------------------------------------+
//
// The Location Area Code (LAC) consists of 2 octets. Bit 8 of Octet c+3 is the
// most significant bit and bit 1 of Octet c+4 the least significant bit. The
// coding of the location area code is the responsibility of each administration.
// Coding using full hexadecimal representation  (binary, not ASCII encoding)
// shall be used (see 3GPP TS 23.003 [2]).
// The Routing Area Code (RAC) consists of 2 octets. Only Octet c+5 contains
// the RAC. Octet c+6 is coded as all 1's (11111111). The RAC is defined by the
// operator. Coding using full hexadecimal representation (binary, not ASCII
// encoding) shall be used (see 3GPP TS 23.003 [2]).
type RAIField struct {
	Mcc string `json:"mcc,omitempty"`
	Mnc string `json:"mnc,omitempty"`
	LAC uint16 `json:"lAC,omitempty"`
	RAC uint16 `json:"rAC,omitempty"`
}

func (r *RAIField) MarshalBinary() (data []byte, err error) {
	data, err = marshalMccMnc(r.Mcc, r.Mnc)
	if err != nil {
		return nil, err
	}

	tmp := make([]byte, 4)
	binary.BigEndian.PutUint16(tmp, r.LAC)
	binary.BigEndian.PutUint16(tmp[2:], r.RAC)
	data = append(data, tmp...)

	return data, nil
}

func (r *RAIField) UnmarshalBinary(data []byte) (err error) {
	length := uint16(len(data))
	if length < 7 {
		return fmt.Errorf("RAIField Inadequate TLV length: %d", length)
	}
	r.Mcc, r.Mnc, err = unmarshalMccMnc(data[0:3])
	if err != nil {
		return err
	}
	r.LAC = binary.BigEndian.Uint16(data[3:])
	r.RAC = binary.BigEndian.Uint16(data[5:])

	return nil
}

// The coding of SAI (Service Area Identifier) is depicted in Figure 8.21.2-1. Only
// zero or one SAI field shall be present in ULI IE.
//
//	                                        Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +-------------------------------+-------------------------------+
//	   a     |           MCC Digit 2         |          MCC Digit 1          |
//	         |-------------------------------+-------------------------------|
//	  a+1    |           MNC Digit 3         |          MCC Digit 3          |
//	         |-------------------------------+-------------------------------|
//	  a+2    |           MNC Digit 2         |          MNC digit 1          |
//	         |-------------------------------+-------------------------------|
//
// a+3 to a+4|                   Location Area Code (LAC)                    |
//
//	|---------------------------------------------------------------|
//
// a+5 to a+6|                      Service Area Code (SAC)                  |
//
//	+---------------------------------------------------------------+
//
// The Location Area Code (LAC) consists of 2 octets. Bit 8 of Octet b+3 is the most
// significant bit and bit 1 of Octet b+4 the least significant bit. The coding of
// the location area code is the responsibility of each administration. Coding using
// full hexadecimal representation (binary, not ASCII encoding) shall be used.
// The Service Area Code (SAC) consists of 2 octets. Bit 8 of Octet b+5 is the most
// significant bit and bit 1 of Octet b+6 the least significant bit. The SAC is defined
// by the operator. See 3GPP TS 23.003 [2] clause 12.5 for more information.
type SAIField struct {
	Mcc string `json:"mcc,omitempty"`
	Mnc string `json:"mnc,omitempty"`
	LAC uint16 `json:"lAC,omitempty"`
	SAC uint16 `json:"sAC,omitempty"`
}

func (s *SAIField) MarshalBinary() (data []byte, err error) {
	data, err = marshalMccMnc(s.Mcc, s.Mnc)
	if err != nil {
		return nil, err
	}

	tmp := make([]byte, 4)
	binary.BigEndian.PutUint16(tmp, s.LAC)
	binary.BigEndian.PutUint16(tmp[2:], s.SAC)
	data = append(data, tmp...)

	return data, nil
}

func (s *SAIField) UnmarshalBinary(data []byte) (err error) {
	length := uint16(len(data))
	if length < 7 {
		return fmt.Errorf("SAIField Inadequate TLV length: %d", length)
	}

	s.Mcc, s.Mnc, err = unmarshalMccMnc(data[0:3])
	if err != nil {
		return err
	}
	s.LAC = binary.BigEndian.Uint16(data[3:])
	s.SAC = binary.BigEndian.Uint16(data[5:])

	return nil
}

func (u *UserLocationInfo) MarshalBinary() (data []byte, err error) {
	tmpOctet := btou(u.NCGI)<<6 |
		btou(u.LAI)<<5 |
		btou(u.ECGI)<<4 |
		btou(u.TAI)<<3 |
		btou(u.RAI)<<2 |
		btou(u.SAI)<<1 |
		btou(u.CGI)
	data = append(data, tmpOctet)

	if u.CGI && u.CGIField != nil {
		tmp, err := u.CGIField.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Marshal CGI failed: %v", err)
		}
		data = append(data, tmp...)
	}

	if u.SAI && u.SAIField != nil {
		tmp, err := u.SAIField.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Marshal SAI failed: %v", err)
		}
		data = append(data, tmp...)
	}

	if u.RAI && u.RAIField != nil {
		tmp, err := u.RAIField.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Marshal RAI failed: %v", err)
		}
		data = append(data, tmp...)
	}

	if u.TAI && u.TAIField != nil {
		tmp, err := u.TAIField.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Marshal TAI failed: %v", err)
		}
		data = append(data, tmp...)
	}

	if u.ECGI && u.ECGIField != nil {
		tmp, err := u.ECGIField.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Marshal ECGI failed: %v", err)
		}
		data = append(data, tmp...)
	}

	if u.LAI && u.LAIField != nil {
		tmp, err := u.LAIField.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Marshal LAI failed: %v", err)
		}
		data = append(data, tmp...)
	}

	if u.NCGI && u.NCGIField != nil {
		tmp, err := u.NCGIField.MarshalBinary()
		if err != nil {
			return nil, fmt.Errorf("Marshal NCGI failed: %v", err)
		}
		data = append(data, tmp...)
	}

	return data, nil
}

func (u *UserLocationInfo) UnmarshalBinary(data []byte) error {
	var idx = 0
	length := len(data)
	if length < 1+idx {
		return fmt.Errorf("Inadequate length of UserLocationInfo, got zero length")
	}
	u.NCGI = utob(data[0] & BitMask7)
	u.LAI = utob(data[0] & BitMask6)
	u.ECGI = utob(data[0] & BitMask5)
	u.TAI = utob(data[0] & BitMask4)
	u.RAI = utob(data[0] & BitMask3)
	u.SAI = utob(data[0] & BitMask2)
	u.CGI = utob(data[0] & BitMask1)

	idx += 1

	if u.CGI {
		if length < idx+UliCgiFieldLen {
			return fmt.Errorf("IE UserLocationInfo Field CGI Inadequate TLV length: %d,idx: %d", length, idx)
		}
		u.CGIField = &CGIField{}
		err := u.CGIField.UnmarshalBinary(data[idx : idx+UliCgiFieldLen])
		if err != nil {
			return fmt.Errorf("Failed CGIField Unmarshal: %v", err)
		}
		idx += UliCgiFieldLen
	}

	if u.SAI {
		if length < idx+UliSaiFieldLen {
			return fmt.Errorf("IE UserLocationInfo Field SAI Inadequate TLV length: %d,idx: %d", length, idx)
		}
		u.SAIField = &SAIField{}
		err := u.SAIField.UnmarshalBinary(data[idx : idx+UliSaiFieldLen])
		if err != nil {
			return fmt.Errorf("Failed SAIField Unmarshal: %v", err)
		}
		idx += UliSaiFieldLen
	}

	if u.RAI {
		if length < idx+UliRaiFieldLen {
			return fmt.Errorf("IE UserLocationInfo Field RAI Inadequate TLV length: %d,idx: %d", length, idx)
		}
		u.RAIField = &RAIField{}
		err := u.RAIField.UnmarshalBinary(data[idx : idx+UliRaiFieldLen])
		if err != nil {
			return fmt.Errorf("Failed RAIField Unmarshal: %v", err)
		}
		idx += UliRaiFieldLen
	}

	if u.TAI {
		if length < idx+UliTaiFieldLen {
			return fmt.Errorf("IE UserLocationInfo Field TAI Inadequate TLV length: %d,idx: %d", length, idx)
		}
		u.TAIField = &TAIField{}
		err := u.TAIField.UnmarshalBinary(data[idx : idx+UliTaiFieldLen])
		if err != nil {
			return fmt.Errorf("Failed TAIField Unmarshal: %v", err)
		}
		idx += UliTaiFieldLen
	}

	if u.ECGI {
		if length < idx+UliEcgiFieldLen {
			return fmt.Errorf("IE UserLocationInfo Field ECGI Inadequate TLV length: %d,idx: %d", length, idx)
		}
		u.ECGIField = &ECGIField{}
		err := u.ECGIField.UnmarshalBinary(data[idx : idx+UliEcgiFieldLen])
		if err != nil {
			return fmt.Errorf("Failed ECGIField Unmarshal: %v", err)
		}
		idx += UliEcgiFieldLen
	}

	if u.LAI {
		if length < idx+UliLaiFieldLen {
			return fmt.Errorf("IE UserLocationInfo Field LAI Inadequate TLV length: %d,idx: %d", length, idx)
		}
		u.LAIField = &LAIField{}
		err := u.LAIField.UnmarshalBinary(data[idx : idx+UliLaiFieldLen])
		if err != nil {
			return fmt.Errorf("Failed LAIField Unmarshal: %v", err)
		}
		idx += UliLaiFieldLen
	}
	if u.NCGI {
		if length < idx+UliNcgiFieldLen {
			return fmt.Errorf("IE UserLocationInfo Field NCGI Inadequate TLV length: %d,idx: %d", length, idx)
		}
		u.NCGIField = &NCGIField{}
		err := u.NCGIField.UnmarshalBinary(data[idx : idx+UliNcgiFieldLen])
		if err != nil {
			return fmt.Errorf("Failed NCGIField Unmarshal: %v", err)
		}
		idx += UliNcgiFieldLen
	}
	return nil
}

package pfcpie

import (
	"bytes"
	"fmt"
)

// TS 29.244 8.2.101: User ID */
/*

The User ID IE type shall be encoded as shown in Figure 8.2.101-1.

+---------------------------------------------------------------------------------------------------------+
|            | Bits                                                                                       |
|------------|--------------------------------------------------------------------------------------------|
| Octets     | 8                                  | 7    | 6     | 5     | 4    | 3       | 2     | 1     |
|------------|--------------------------------------------------------------------------------------------|
| 1 to 2     | Type = 141 (decimal)                                                                       |
|------------|--------------------------------------------------------------------------------------------|
| 3 to 4     | Length = n                                                                                 |
|------------|--------------------------------------------------------------------------------------------|
| 5          | Spare                              | PEIF | GPSIF | SUPIF | NAIF | MSISDNF | IMEIF | IMSIF |
|------------|--------------------------------------------------------------------------------------------|
| 6          | Length of IMSI                                                                             |
|------------|--------------------------------------------------------------------------------------------|
| 7 to a     | IMSI                                                                                       |
|------------|--------------------------------------------------------------------------------------------|
| b          | Length of IMEI                                                                             |
|------------|--------------------------------------------------------------------------------------------|
| (b+1) to c | IMEI                                                                                       |
|------------|--------------------------------------------------------------------------------------------|
| d          | Length of MSISDN                                                                           |
|------------|--------------------------------------------------------------------------------------------|
| (d+1) to e | MSISDN                                                                                     |
|------------|--------------------------------------------------------------------------------------------|
| f          | Length of NAI                                                                              |
|------------|--------------------------------------------------------------------------------------------|
| (f+1) to g | NAI                                                                                        |
|------------|--------------------------------------------------------------------------------------------|
| h          | Length of SUPI                                                                             |
|------------|--------------------------------------------------------------------------------------------|
| (h+1) to i | SUPI                                                                                       |
|------------|--------------------------------------------------------------------------------------------|
| j          | Length of GPSI                                                                             |
|------------|--------------------------------------------------------------------------------------------|
| (j+1) to k | GPSI                                                                                       |
|------------|--------------------------------------------------------------------------------------------|
| l          | Length of PEI                                                                              |
|------------|--------------------------------------------------------------------------------------------|
| (l+1) to m | PEI                                                                                        |
|------------|--------------------------------------------------------------------------------------------|
| p to (n+4) | These octet(s) is/are present only if explicitly specified                                 |
+------------+--------------------------------------------------------------------------------------------+

The following flags are coded within Octet 5:

- Bit 1 – IMSIF: If this bit is set to "1", then the Length of IMSI and IMSI fields shall be present, otherwise these
  fields shall not be present.
- Bit 2 – IMEIF: If this bit is set to "1", then the Length of IMEI and IMEI fields shall be present, otherwise these
  fields shall not be present.
- Bit 3 – MSISDNF: If this bit is set to "1", then the Length of MSISDN and MSISDN fields shall be present, otherwise
  these fields shall not be present.
- Bit 4 – NAIF: If this bit is set to "1", then the Length of NAI and NAI fields shall be present, otherwise these
  fields shall not be present.
- Bit 5 – SUPIF: If this bit is set to "1", then the Length of SUPI and SUPI fields shall be present, otherwise these
  fields shall not be present.
- Bit 6 – GPSIF: If this bit is set to "1", then the Length of GPSI and GPSI fields shall be present, otherwise these
  fields shall not be present.
- Bit 7 – PEIF: If this bit is set to "1", then the Length of PEI and PEI fields shall be present, otherwise these
  fields shall not be present.
- Bit 8: Spare, for future use and set to "0".

One or more flags may be set to "1".

For 5GS User Identities:

- The SUPI field shall only be used for carrying a Global Cable Identifier (GCI) or a Global Line Identifier (GLI). The
  IMSI and NAI, if received by the SMF in the SUPI, shall be included in the IMSI and NAI field respectively.
- The GPSI field shall only be used for carrying an External Identifier. The MSISDN, if received by the SMF in the
  SUPI, shall be included in the MSISDN field.
- The PEI field shall only be used for carrying an MAC address or an Extended Unique Identifier. The IMEI, if received
  by the SMF in the PEI, shall be included in the IMEI field.

The coding of IMSI field, from octets 7 to 'a' shall be encoded as the octets 5 to n+4 of the IMSI IE type specified in
clause 8.3 of 3GPP TS 29.274.

The coding of IMEI field, in octets 'b+1' to 'c' shall be encoded as the octets 5 to n+4 of the MEI IE type specified
in clause 8.10 of 3GPP TS 29.274.

The coding of MSISDN field, in octets 'd+1' to 'e' shall be encoded as the octets 5 to n+4 of the MSISDN IE type
specified in clause 8.11 of 3GPP TS 29.274.

The coding of the SUPI field, in octets 'h+1' to 'i' shall be encoded as the Supi data type specified in clause 5.3.2
of 3GPP TS 29.571.

The coding of the GPSI field, in octets 'j+1' to 'k' shall be encoded as the Gpsi data type specified in clause 5.3.2
of 3GPP TS 29.571.

The coding of the PEI field, in octets 'l+1' to 'm' shall be encoded as the Pei data type specified in clause 5.3.2 of
3GPP TS 29.571.

The NAI field, in octets 'f+1' to 'g' shall be encoded as an Octet String (see IETF RFC 4282).

*/

// application use UserIdData, for encoding/decoding use UserId
type UserID struct {
	Imsi   string `json:"imsi,omitempty"`
	Imei   string `json:"imei,omitempty"`
	Msisdn string `json:"msisdn,omitempty"`
	Nai    string `json:"nai,omitempty"`
	Supi   string `json:"supi,omitempty"`
	Gpsi   string `json:"gpsi,omitempty"`
	Pei    string `json:"pei,omitempty"`
}

func (d *UserID) MarshalBinary() (data []byte, err error) {
	var flags byte
	data = append(data, 0 /* flags */)

	if d.Imsi != "" {
		flags |= BitMask1

		data = append(data, byte((len(d.Imsi)+1)/2))
		data = append(data, []byte(ConvertStringToBCDCode(d.Imsi))...)
	}

	if d.Imei != "" {
		flags |= BitMask2

		data = append(data, byte((len(d.Imei)+1)/2))
		data = append(data, []byte(ConvertStringToBCDCode(d.Imei))...)
		if len(d.Imei) == 15 {
			data[len(data)-1] |= 0xf0
		}
	}

	if d.Msisdn != "" {
		flags |= BitMask3

		data = append(data, byte((len(d.Msisdn)+1)/2))
		data = append(data, []byte(ConvertStringToBCDCode(d.Msisdn))...)
	}

	if d.Nai != "" {
		flags |= BitMask4

		data = append(data, byte(len(d.Nai)))
		data = append(data, []byte(d.Nai)...)
	}

	if d.Supi != "" {
		flags |= BitMask5

		data = append(data, byte(len(d.Supi)))
		data = append(data, []byte(d.Supi)...)
	}

	if d.Gpsi != "" {
		flags |= BitMask6

		data = append(data, byte(len(d.Gpsi)))
		data = append(data, []byte(d.Gpsi)...)
	}

	if d.Pei != "" {
		flags |= BitMask7

		data = append(data, byte(len(d.Pei)))
		data = append(data, []byte(d.Pei)...)
	}

	if flags == 0 {
		return nil, fmt.Errorf("at least one user identifier should be specified: [userid: %#v]", d)
	}

	data[0] = flags
	return data, nil
}

func (d *UserID) UnmarshalBinary(data []byte) error {
	reader := bytes.NewReader(data)
	flags, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("reading the flags has failed: [encoded: %x]", data)
	}
	var length byte

	if flags&BitMask1 != 0 {
		length, err = reader.ReadByte()
		if err != nil {
			return fmt.Errorf("reading the imsi length has failed: [encoded: %x]", data)
		}

		imsi := make([]byte, length)
		_, err = reader.Read(imsi)
		if err != nil {
			return fmt.Errorf("reading the imsi has failed: [encoded: %x]", data)
		}
		d.Imsi = ConvertBCDCodeToString(imsi)
	}

	if flags&BitMask2 != 0 {
		length, err = reader.ReadByte()
		if err != nil {
			return fmt.Errorf("reading the imei length has failed: [encoded: %x]", data)
		}

		imei := make([]byte, length)
		_, err = reader.Read(imei)
		if err != nil {
			return fmt.Errorf("reading the imei has failed: [encoded: %x]", data)
		}
		d.Imei = ConvertBCDCodeToString(imei)
	}

	if flags&BitMask3 != 0 {
		length, err = reader.ReadByte()
		if err != nil {
			return fmt.Errorf("reading the msisdn length has failed: [encoded: %x]", data)
		}

		msisdn := make([]byte, length)
		_, err = reader.Read(msisdn)
		if err != nil {
			return fmt.Errorf("reading the msisdn has failed: [encoded: %x]", data)
		}
		d.Msisdn = ConvertBCDCodeToString(msisdn)
	}

	if flags&BitMask4 != 0 {
		length, err = reader.ReadByte()
		if err != nil {
			return fmt.Errorf("reading the nai length has failed: [encoded: %x]", data)
		}

		nai := make([]byte, length)
		_, err = reader.Read(nai)
		if err != nil {
			return fmt.Errorf("reading the nai has failed: [encoded: %x]", data)
		}
		d.Nai = string(nai)
	}

	if flags&BitMask5 != 0 {
		length, err = reader.ReadByte()
		if err != nil {
			return fmt.Errorf("reading the supi length has failed: [encoded: %x]", data)
		}

		supi := make([]byte, length)
		_, err = reader.Read(supi)
		if err != nil {
			return fmt.Errorf("reading the supi has failed: [encoded: %x]", data)
		}
		d.Supi = string(supi)
	}

	if flags&BitMask6 != 0 {
		length, err = reader.ReadByte()
		if err != nil {
			return fmt.Errorf("reading the gpsi length has failed: [encoded: %x]", data)
		}

		gpsi := make([]byte, length)
		_, err = reader.Read(gpsi)
		if err != nil {
			return fmt.Errorf("reading the gpsi has failed: [encoded: %x]", data)
		}
		d.Gpsi = string(gpsi)
	}

	if flags&BitMask7 != 0 {
		length, err = reader.ReadByte()
		if err != nil {
			return fmt.Errorf("reading the pei length has failed: [encoded: %x]", data)
		}

		pei := make([]byte, length)
		_, err = reader.Read(pei)
		if err != nil {
			return fmt.Errorf("reading the pei has failed: [encoded: %x]", data)
		}
		d.Pei = string(pei)
	}

	return nil
}

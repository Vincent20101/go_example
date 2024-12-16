package pfcpie

import (
	"fmt"
)

const (
	UpFunctionFeaturesLoad  uint8 = 0x1
	UpFunctionFeaturesOvrl  uint8 = 0x2
	UpFunctionFeaturesEpfar uint8 = 0x4
	UpFunctionFeaturesSset  uint8 = 0x8
)

// UPFunctionFeatures IE, Refer to [TS29244 8.2.25 UP Function Features]
//
// The UP Function Features IE indicates the features supported by the UP
// function. It is coded as depicted in Figure 8.2.25-1.
//
//	                                       Bits
//	Octets       8       7       6       5       4       3       2       1
//	         +---------------------------------------------------------------+
//	1 to 2   |                       Type = 43(decimal)                      |
//	         |---------------------------------------------------------------|
//	3 to 4   |                           Length = n                          |
//	         |---------------------------------------------------------------|
//	5 to 6   |                       Supported-Features                      |
//	         |---------------------------------------------------------------|
//	7 to 8   |                  Additional Supported-Features 1              |
//	         |---------------------------------------------------------------|
//	9 to 10  |                  Additional Supported-Features 2              |
//	         |---------------------------------------------------------------|
//
// 11 to n+4 |   These octets(s) is/are present only if explicity specified  |
//
//	+---------------------------------------------------------------+
//	                 Figure 8.2.25-1: UP Function Features
//
// The UP Function Features IE takes the form of a bitmask where each bit set
// indicates that the corresponding feature is supported. Spare bits shall be
// ignored by the receiver. The same bitmask is defined for all PFCP
// interfaces.
//
// The following table specifies the features defined on PFCP interfaces and
// the interfaces on which they apply.
//
//	...
//	Table 8.2.25-1: UP Function Features
//
// The table is not shown here, please refer to the spec for more detail.
type UPFunctionFeatures struct {

	// Octets 5 to 6: Supported-Features

	Bucp bool `json:"bucp,omitempty"` // Octet5/bit1
	Ddnd bool `json:"ddnd,omitempty"` // Octet5/bit2
	Dlbd bool `json:"dlbd,omitempty"` // Octet5/bit3
	Trst bool `json:"trst,omitempty"` // Octet5/bit4
	Ftup bool `json:"ftup,omitempty"` // Octet5/bit5
	Pfdm bool `json:"pfdm,omitempty"` // Octet5/bit6
	Heeu bool `json:"heeu,omitempty"` // Octet5/bit7
	Treu bool `json:"treu,omitempty"` // Octet5/bit8

	Empu  bool `json:"empu,omitempty"`  // Octet6/bit1
	Pdiu  bool `json:"pdiu,omitempty"`  // Octet6/bit2
	Udbc  bool `json:"udbc,omitempty"`  // Octet6/bit3
	Quoac bool `json:"quoac,omitempty"` // Octet6/bit4
	Trace bool `json:"trace,omitempty"` // Octet6/bit5
	Frrt  bool `json:"frrt,omitempty"`  // Octet6/bit6
	Pfde  bool `json:"pfde,omitempty"`  // Octet6/bit7
	Epfar bool `json:"epfar,omitempty"` // Octet6/bit8

	// Octets 7 to 8: Additional Supported-Features 1

	Dpdra bool `json:"dpdra,omitempty"` // Octet7/bit1
	Adpdp bool `json:"adpdp,omitempty"` // Octet7/bit2
	Ueip  bool `json:"ueip,omitempty"`  // Octet7/bit3
	Sset  bool `json:"sset,omitempty"`  // Octet7/bit4
	Mnop  bool `json:"mnop,omitempty"`  // Octet7/bit5
	Mte   bool `json:"mte,omitempty"`   // Octet7/bit6
	Bundl bool `json:"bundl,omitempty"` // Octet7/bit7
	Gcom  bool `json:"gcom,omitempty"`  // Octet7/bit8

	Mpas   bool `json:"mpas,omitempty"`   // Octet8/bit1
	Rttl   bool `json:"rttl,omitempty"`   // Octet8/bit2
	Vtime  bool `json:"vtime,omitempty"`  // Octet8/bit3
	Norp   bool `json:"norp,omitempty"`   // Octet8/bit4
	Iptv   bool `json:"iptv,omitempty"`   // Octet8/bit5
	Ip12pl bool `json:"ip12pl,omitempty"` // Octet8/bit0
	Tscu   bool `json:"tscu,omitempty"`   // Octet8/bit7
	Mptcp  bool `json:"mptcp,omitempty"`  // Octet8/bit8

	// Octets 9 to 10: Additional Supported-Features 1

	Atsssll bool `json:"atsssll,omitempty"` // Octet9/bit1
	Qfqm    bool `json:"qfqm,omitempty"`    // Octet9/bit2
	Gpqm    bool `json:"gpqm,omitempty"`    // Octet9/bit3
	Mtedt   bool `json:"mtedt,omitempty"`   // Octet9/bit4
	Ciot    bool `json:"ciot,omitempty"`    // Octet9/bit5
	Ethar   bool `json:"ethar,omitempty"`   // Octet9/bit6
	Ddds    bool `json:"ddds,omitempty"`    // Octet9/bit7
	Rds     bool `json:"rds,omitempty"`     // Octet9/bit8

	Rttwp bool `json:"rttwp,omitempty"` // Octet10/bit1
}

func (c *UPFunctionFeatures) GetSupportedFeatures() uint64 {
	// No need to consider the big and small issues
	tmp1 := btou(c.Bucp) |
		btou(c.Ddnd)<<1 |
		btou(c.Dlbd)<<2 |
		btou(c.Trst)<<3 |
		btou(c.Ftup)<<4 |
		btou(c.Pfdm)<<5 |
		btou(c.Heeu)<<6 |
		btou(c.Treu)<<7
	tmp2 := btou(c.Empu) |
		btou(c.Pdiu)<<1 |
		btou(c.Udbc)<<2 |
		btou(c.Quoac)<<3 |
		btou(c.Trace)<<4 |
		btou(c.Frrt)<<5 |
		btou(c.Pfde)<<6 |
		btou(c.Epfar)<<7
	tmp3 := btou(c.Dpdra) |
		btou(c.Adpdp)<<1 |
		btou(c.Ueip)<<2 |
		btou(c.Sset)<<3 |
		btou(c.Mnop)<<4 |
		btou(c.Mte)<<5 |
		btou(c.Bundl)<<6 |
		btou(c.Gcom)<<7
	tmp5 := btou(c.Atsssll) |
		btou(c.Qfqm)<<1 |
		btou(c.Gpqm)<<2 |
		btou(c.Mtedt)<<3 |
		btou(c.Ciot)<<4 |
		btou(c.Ethar)<<5 |
		btou(c.Ddds)<<6 |
		btou(c.Rds)<<7

	tmp64 := uint64(tmp5) << 32
	tmp64 = tmp64 | uint64(tmp3)<<16
	tmp64 = tmp64 | uint64(tmp2)<<8
	tmp64 = tmp64 | uint64(tmp1)
	return tmp64
}

func (c *UPFunctionFeatures) MarshalBinary() (data []byte, err error) {
	// Octets 5 to 6: Supported-Features
	// Octet 5
	tmpoctet := btou(c.Bucp) |
		btou(c.Ddnd)<<1 |
		btou(c.Dlbd)<<2 |
		btou(c.Trst)<<3 |
		btou(c.Ftup)<<4 |
		btou(c.Pfdm)<<5 |
		btou(c.Heeu)<<6 |
		btou(c.Treu)<<7
	data = append([]byte(""), byte(tmpoctet))

	// Octet 6
	tmpoctet = btou(c.Empu) |
		btou(c.Pdiu)<<1 |
		btou(c.Udbc)<<2 |
		btou(c.Quoac)<<3 |
		btou(c.Trace)<<4 |
		btou(c.Frrt)<<5 |
		btou(c.Pfde)<<6 |
		btou(c.Epfar)<<7
	data = append(data, byte(tmpoctet))

	// Octets 7 to 8: Additional Supported-Features 1
	// Octet 7
	tmpoctet = btou(c.Dpdra) |
		btou(c.Adpdp)<<1 |
		btou(c.Ueip)<<2 |
		btou(c.Sset)<<3 |
		btou(c.Mnop)<<4 |
		btou(c.Mte)<<5 |
		btou(c.Bundl)<<6 |
		btou(c.Gcom)<<7
	data = append(data, byte(tmpoctet))

	// Octet 8
	tmpoctet = btou(c.Mpas) |
		btou(c.Rttl)<<1 |
		btou(c.Vtime)<<2 |
		btou(c.Norp)<<3 |
		btou(c.Iptv)<<4 |
		btou(c.Ip12pl)<<5 |
		btou(c.Tscu)<<6 |
		btou(c.Mptcp)<<7
	data = append(data, byte(tmpoctet))

	// Octets 9 to 10: Additional Supported-Features 1
	// Octet 9
	tmpoctet = btou(c.Atsssll) |
		btou(c.Qfqm)<<1 |
		btou(c.Gpqm)<<2 |
		btou(c.Mtedt)<<3 |
		btou(c.Ciot)<<4 |
		btou(c.Ethar)<<5 |
		btou(c.Ddds)<<6 |
		btou(c.Rds)<<7
	data = append(data, byte(tmpoctet))

	// Octet 10
	tmpoctet = btou(c.Rttwp)
	data = append(data, byte(tmpoctet))

	return data, nil
}

func (c *UPFunctionFeatures) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0

	// Octets 5 to 6: Supported-Features
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("Umarshal UPFunctionFeatures failed, inadequate TLV length: %d", length)
	}
	c.Bucp = utob(uint8(data[idx]) & BitMask1)
	c.Ddnd = utob(uint8(data[idx]) & BitMask2)
	c.Dlbd = utob(uint8(data[idx]) & BitMask3)
	c.Trst = utob(uint8(data[idx]) & BitMask4)
	c.Ftup = utob(uint8(data[idx]) & BitMask5)
	c.Pfdm = utob(uint8(data[idx]) & BitMask6)
	c.Heeu = utob(uint8(data[idx]) & BitMask7)
	c.Treu = utob(uint8(data[idx]) & BitMask8)

	idx = idx + 1

	// Octet 6
	if length < idx+1 {
		return fmt.Errorf("Umarshal UPFunctionFeatures failed, inadequate TLV length: %d", length)
	}
	c.Empu = utob(uint8(data[idx]) & BitMask1)
	c.Pdiu = utob(uint8(data[idx]) & BitMask2)
	c.Udbc = utob(uint8(data[idx]) & BitMask3)
	c.Quoac = utob(uint8(data[idx]) & BitMask4)
	c.Trace = utob(uint8(data[idx]) & BitMask5)
	c.Frrt = utob(uint8(data[idx]) & BitMask6)
	c.Pfde = utob(uint8(data[idx]) & BitMask7)
	c.Epfar = utob(uint8(data[idx]) & BitMask8)

	idx = idx + 1

	// Octets 7 to 8: Additional Supported-Features 1
	// Octet 7
	if length < idx+1 {
		return nil
		//return fmt.Errorf("Umarshal UPFunctionFeatures failed, inadequate TLV length: %d", length)
	}
	c.Dpdra = utob(uint8(data[idx]) & BitMask1)
	c.Adpdp = utob(uint8(data[idx]) & BitMask2)
	c.Ueip = utob(uint8(data[idx]) & BitMask3)
	c.Sset = utob(uint8(data[idx]) & BitMask4)
	c.Mnop = utob(uint8(data[idx]) & BitMask5)
	c.Mte = utob(uint8(data[idx]) & BitMask6)
	c.Bundl = utob(uint8(data[idx]) & BitMask7)
	c.Gcom = utob(uint8(data[idx]) & BitMask8)

	idx = idx + 1

	// Octet 8
	if length < idx+1 {
		return nil
		//return fmt.Errorf("Umarshal UPFunctionFeatures failed, inadequate TLV length: %d, idx:%d", length, idx)
	}
	c.Mpas = utob(uint8(data[idx]) & BitMask1)
	c.Rttl = utob(uint8(data[idx]) & BitMask2)
	c.Vtime = utob(uint8(data[idx]) & BitMask3)
	c.Norp = utob(uint8(data[idx]) & BitMask4)
	c.Iptv = utob(uint8(data[idx]) & BitMask5)
	c.Ip12pl = utob(uint8(data[idx]) & BitMask6)
	c.Tscu = utob(uint8(data[idx]) & BitMask7)
	c.Mptcp = utob(uint8(data[idx]) & BitMask8)

	idx = idx + 1

	// Octets 9 to 10: Additional Supported-Features 1
	// Octet 9
	if length < idx+1 {
		// Compatible with the old version, old spec dont have this feild.
		return nil
		//return fmt.Errorf("Umarshal UPFunctionFeatures failed, inadequate TLV length: %d, idx:%d", length, idx)
	}
	c.Atsssll = utob(uint8(data[idx]) & BitMask1)
	c.Qfqm = utob(uint8(data[idx]) & BitMask2)
	c.Gpqm = utob(uint8(data[idx]) & BitMask3)
	c.Mtedt = utob(uint8(data[idx]) & BitMask4)
	c.Ciot = utob(uint8(data[idx]) & BitMask5)
	c.Ethar = utob(uint8(data[idx]) & BitMask6)
	c.Ddds = utob(uint8(data[idx]) & BitMask7)
	c.Rds = utob(uint8(data[idx]) & BitMask8)

	idx = idx + 1

	// Octet 10
	if length < idx+1 {
		// Compatible with the old version, old spec dont have this feild.
		return nil
		//return fmt.Errorf("Umarshal UPFunctionFeatures failed, inadequate TLV length: %d, idx:%d", length, idx)
	}
	c.Rttwp = utob(uint8(data[idx]) & BitMask1)

	return nil
}

package pfcpie

import "fmt"

const (
	RATTYPE_RESERVED       = 0
	RATTYPE_UTRAN          = 1
	RATTYPE_GERAN          = 2
	RATTYPE_WLAN           = 3
	RATTYPE_GAN            = 4
	RATTYPE_HSPA_EVOLUTION = 5
	RATTYPE_EUTRA          = 6
	RATTYPE_VIRTUAL        = 7
	RATTYPE_EUTRAN_NB_IOT  = 8
	RATTYPE_LTE_M          = 9
	RATTYPE_NR             = 10
)

// RAT Type
// Casa Specific IE for Verizon use
//
//	Octets        8       7       6       5       4       3       2       1
//	          +-------------------------------------------------------------+
//	1         |                   Type = 4(decimal)                         |
//	          |-------------------------------------------------------------|
//	2         |                      Length = n                             |
//	          |-------------------------------------------------------------|
//	3 to n    |	                   RAT Type Value                           |
//	          |-------------------------------------------------------------|
//
//	  +--------------------------+-------------------+
//	  |        RAT Types         |   Value (Decimal) |
//	  +--------------------------+-------------------+
//	  |       <reserved>         |        0          |
//	  |         UTRAN            |        1          |
//	  |         GERAN            |        2          |
//	  |          WLAN            |        3          |
//	  |          GAN             |        4          |
//	  |     HSPA Evolution       |        5          |
//	  |   EUTRAN (WB-E-UTRAN)    |        6          |
//	  |        Virtual           |        7          |
//	  |     EUTRAN-NB-IoT        |        8          |
//	  |        LTE-M             |        9          |
//	  |          NR              |       10          |
//	  |        <spare>           |      11-255       |
//	  +--------------------------+-------------------+
type RATType struct {
	RatTypeVal uint8 `json:"ratTypeVal,omitempty"`
}

func (r *RATType) MarshalBinary() (data []byte, err error) {
	data = append([]byte(""), r.RatTypeVal)
	return data, nil
}

func (r *RATType) UnmarshalBinary(data []byte) error {
	length := len(data)
	if length < 1 {
		return fmt.Errorf("IE RATType Inadequate TLV length: %d", length)
	}
	r.RatTypeVal = uint8(data[0])
	return nil
}

package pfcpie

import "fmt"

const (
	// Bits 6 to 8 defines the timer value unit for the timer as follows:
	// Bits
	// 8 7 6
	TIMER_UNIT_2_SECONDS  uint8 = 0x00 // 0 0 0  value is incremented in multiples of 2 seconds
	TIMER_UNIT_1_MINUTE   uint8 = 0x20 // 0 0 1  value is incremented in multiples of 1 minute
	TIMER_UNIT_10_MINUTES uint8 = 0x40 // 0 1 0  value is incremented in multiples of 10 minutes
	TIMER_UNIT_1_HOUR     uint8 = 0x60 // 0 1 1  value is incremented in multiples of 1 hour
	TIMER_UNIT_10_HOURS   uint8 = 0x80 // 1 0 0  value is incremented in multiples of 10 hours
	TIMER_UNIT_INFINITE   uint8 = 0xE0 // 1 1 1  value indicates that the timer is infinite
)

// TS29.244-g40 8.2.78	Graceful Release Period
// The purpose of the Graceful Release Period IE is to specify a specific time for a graceful release. The Graceful Release Period IE shall be encoded as shown in Figure 8.2.78-1 and table 8.2.78.1.
//
//	                                          Bits
//	 Octets        8       7       6       5       4       3       2       1
//	           +----------------------------------------------------------------+
//	 1 to 2    |                   Type = 112(decimal)                          |
//	           |----------------------------------------------------------------|
//	 3 to 4    |                      Length = n                                |
//	           |-----------------------+----------------------------------------|
//		5         |     Timer unit	      |                Timer value             |
//	           |-----------------------+----------------------------------------|
//		6 to (n+4)|   These octet(s) is/are present only if explicitly specified   |
//	           +----------------------------------------------------------------+
//
// Figure 8.2.78-1: Graceful Release Period
// Table 8.2.78.1: Graceful Release Period information element
// Timer value
// Bits 5 to 1 represent the binary coded timer value.
// Timer unit
// Bits 6 to 8 defines the timer value unit for the timer as follows:
// Bits
// 8 7 6
// 0 0 0  value is incremented in multiples of 2 seconds
// 0 0 1  value is incremented in multiples of 1 minute
// 0 1 0  value is incremented in multiples of 10 minutes
// 0 1 1  value is incremented in multiples of 1 hour
// 1 0 0  value is incremented in multiples of 10 hours
// 1 1 1  value indicates that the timer is infinite
// Other values shall be interpreted as multiples of 1 minute in this version of the protocol.
// Timer unit and Timer value both set to all "zeros" shall be interpreted as an indication that the timer is stopped.
type GracefulReleasePeriod struct {
	// ie
	TimerUnit  uint8 `json:"timerUnit,omitempty"`
	TimerValue uint8 `json:"timerValue,omitempty"`
	// custom field
	IsInfinite bool   `json:"isInfinite,omitempty"`
	IsStopped  bool   `json:"isStopped,omitempty"`
	TotalSec   uint32 `json:"totalSec,omitempty"`
}

func (g *GracefulReleasePeriod) MarshalBinary() (data []byte, err error) {
	// Octet 5
	tmpUint8 := (g.TimerUnit & 0xE0) | (g.TimerValue & 0x1F)
	data = append([]byte(""), byte(tmpUint8))
	return data, nil
}

func (g *GracefulReleasePeriod) UnmarshalBinary(data []byte) error {
	length := uint16(len(data))

	var idx uint16 = 0
	// Octet 5
	if length < idx+1 {
		return fmt.Errorf("IE GracefulReleasePeriod Inadequate TLV length: %d", length)
	}

	g.TimerUnit = (uint8(data[idx]) & 0xE0)
	g.TimerValue = uint8(data[idx]) & 0x1F
	idx = idx + 1

	g.ToCustomField()

	return nil
}

func (g *GracefulReleasePeriod) ToCustomField() {
	// Octet 5
	if g.TimerUnit == TIMER_UNIT_INFINITE {
		g.IsInfinite = true
	} else if g.TimerUnit == 0 && g.TimerValue == 0 {
		g.IsStopped = true
	}

	switch g.TimerUnit {
	case TIMER_UNIT_2_SECONDS:
		g.TotalSec = uint32(g.TimerValue) * 2
	case TIMER_UNIT_1_MINUTE:
		g.TotalSec = uint32(g.TimerValue) * 1 * 60
	case TIMER_UNIT_10_MINUTES:
		g.TotalSec = uint32(g.TimerValue) * 10 * 60
	case TIMER_UNIT_1_HOUR:
		g.TotalSec = uint32(g.TimerValue) * 1 * 60 * 60
	case TIMER_UNIT_10_HOURS:
		g.TotalSec = uint32(g.TimerValue) * 10 * 60 * 60
	}
}

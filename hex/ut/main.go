package main

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultTickerPeriodTime uint32 = 100 // unit: Millisecond
)

func main() {
	fmt.Println(NonGbrBitRateToHexEncode("150 Mbps"))
	duration := time.Duration(DefaultTickerPeriodTime) * time.Millisecond
	fmt.Println(duration)
}

func NonGbrBitRateToHexEncode(rate string) string {
	var s string
	split := strings.Split(rate, " ")
	if len(split) != 2 {
		s = ""
	}
	val, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		s = ""
	}
	switch split[1] {
	case "bps":
		s = strconv.Itoa(int(val / 1000))
	case "Kbps":
		s = strconv.Itoa(int(val * 1000 / 1000))
	case "Mbps":
		s = strconv.Itoa(int(val * 1000000 / 1000))
	case "Gbps":
		s = strconv.Itoa(int(val * 1000000000 / 1000))
	case "Tbps":
		s = strconv.Itoa(int(val * 1000000000000 / 1000))
	default:
		return ""
	}

	b := make([]byte, 4)
	e, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return ""
	}
	b[0] = byte(e >> 24)
	b[1] = byte(e >> 16)
	b[2] = byte(e >> 8)
	b[3] = byte(e)
	return hex.EncodeToString(b)
}

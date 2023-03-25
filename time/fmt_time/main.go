package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
	"unsafe"
)

func main() {
	//var timestamp time.Time
	//itimeStr := fmt.Sprintf("%02d%02d%02d", timestamp.Hour(), timestamp.Minute(), timestamp.Second())
	//fmt.Println(itimeStr)
	fmt.Println(strconv.FormatInt(time.Now().UnixNano(), 16))
	now := time.Now()
	formatStr := "1504-0700"
	fmt.Println(now.Format(formatStr))
	//fmt.Println(now.UTC())
	//fmt.Println(now.UTC().Hour())
	//timeStr := fmt.Sprintf("%02d/%02d/%02d %02d:%02d:%02d + %02d:%02d:%02d",
	//	now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), now.UTC().Hour(), now.UTC().Minute(), now.UTC().Second())
	//fmt.Printf("time:%s\n", timeStr)

	timeStr := fmt.Sprintf("%02d%02d", now.Hour(), now.Minute())
	// determine timezone
	dateArray := strings.Fields(now.String())
	fmt.Println(dateArray)
	// the 2nd field in the date array is the timezone offset
	timeStr += dateArray[2]
	fmt.Println(timeStr)

	ut := time.Now().UnixNano()
	fmt.Println(ut, "-------", ut>>32)

	//s := "2022-08-23T10:21:28.60701395Z"
	//o, err := time.Parse(time.Kitchen, s)
	//fmt.Println(o, err)

	//fmt.Println(time.Now())
	//fmt.Println(time.ParseInLocation("2006-01-02 15:04:05", "2022-08-23T10:21:28.60701395Z", time.Local))

	fmt.Println(len("2C06000003E8000003E8"))
	src := []byte("Hello")
	encodedStr := hex.EncodeToString(src)

	fmt.Printf("%s\n", encodedStr)

	const s = "2C06000003E8000003E8"
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)
	a16 := strconv.FormatInt(int64(1000), 16) //10进制转换为16进制
	a17 := strconv.FormatInt(int64(1000), 16) //10进制转换为16进制
	fmt.Println(a16 + a17)
	fmt.Printf("%v \n", a16)
	fmt.Printf("%.8x \n", "1000")

	fmt.Println(hex.EncodeToString([]byte("1111")))
	fmt.Println(hex.DecodeString("1111111"))
	//bytes := make([]byte, 8)
	var i int64
	i = 1000
	fmt.Println(IntToBytes(i))

	//fmt.Println(StringToBytes(a16))

	//ss := "2C"
	//
	//n, err := strconv.ParseUint("1000", 16, 8)
	//fmt.Println(IntToBytes(int64(n)))
	//
	//n2 := uint8(n)
	//f := int(*(*int8)(unsafe.Pointer(&n2)))
	//fmt.Println(f)

	eci, err := strconv.ParseUint("1111111", 16, 32)
	fmt.Println(eci)
	fmt.Println(strconv.FormatInt(int64(eci), 16))
	if err != nil {
		fmt.Println(err)
		return
	}

	bEcgi := make([]byte, 4)
	binary.BigEndian.PutUint32(bEcgi, uint32(eci))
	fmt.Println(bEcgi)
	//fmt.Println(strconv.ParseInt(string, 10, 64)(bEcgi))
	fmt.Println(binary.BigEndian.Uint32(bEcgi))

	b := make([]byte, 3)
	mcc := "502"
	fmt.Println(mcc[0] - '0')
	fmt.Println(mcc[1] - '0')
	b[0] = ((mcc[1] - '0') << 4) | (mcc[0] - '0')
	fmt.Println(b)

	toInt, err := SbaBitRateRmToInt("1000 Kbps")
	fmt.Println(*toInt, err)

	b1 := make([]byte, 4)
	e, _ := strconv.ParseUint("104", 10, 32)
	binary.BigEndian.PutUint32(b1, uint32(e))
	fmt.Println(e, b1)
	fmt.Println(hex.EncodeToString(b1))

	b2 := make([]byte, 2)
	e2, _ := strconv.ParseUint("104", 10, 16)
	binary.BigEndian.PutUint16(b2, uint16(e2))
	fmt.Println(e2, b2)
	fmt.Println(hex.EncodeToString(b2))

	b3 := make([]byte, 2)
	e3, _ := strconv.ParseUint("6", 10, 16)
	binary.BigEndian.PutUint16(b3, uint16(e3))
	fmt.Println(e3, b3)
	fmt.Println(hex.EncodeToString(b3))

	var ii uint8
	ii = 10
	fmt.Println(hex.EncodeToString([]byte{ii}))
	fmt.Println(hex.EncodeToString([]byte{104}))

	fmt.Println(bitRateToHexEncode("1000000000 bps"))

	bytesBuffer := bytes.NewBuffer([]byte{})
	err = binary.Write(bytesBuffer, binary.BigEndian, int64(10))
	fmt.Println(err)
	fmt.Println("====", bytesBuffer.Bytes())

	toBytes, _ := IntToBytes(10)
	fmt.Println(toBytes)
	fmt.Println(hex.EncodeToString(toBytes))

	formatInt := strconv.FormatInt(0, 2)
	fmt.Println(formatInt)

	parseInt, err := strconv.ParseInt("1101000", 2, 16)
	fmt.Println(parseInt, err)

	var tmp uint16
	tmp = 104
	fmt.Println(IntToBytes(int64(tmp)))
	fmt.Println(IntToBytes(0))
	fmt.Println(hex.EncodeToString([]byte("00")))

	var st uint8
	fmt.Printf("%T\n", st|0x40|0x20|0x8)
	fmt.Println(st | 0x40 | 0x1 | (10 << 2))
	formatInt = strconv.FormatInt(10<<2, 2)
	fmt.Println(formatInt)
	fmt.Println(10 | 0x40 | 0x1)

	var ti uint16
	ti = 1
	fmt.Println(byte(ti))
	fmt.Println(hex.EncodeToString([]byte{byte(ti)}))
	formatInt = strconv.FormatInt(1, 2)
	fmt.Println(formatInt)
	fmt.Printf("%o", 1)
	fmt.Println([]byte("01101000"))

	b11 := make([]byte, 2)
	binary.BigEndian.PutUint16(b11, uint16(9))
	fmt.Println(b11)
	fmt.Println(hex.EncodeToString(b11))

	var ub []byte = make([]byte, 4)
	utf8.EncodeRune(ub, 1000000)
	fmt.Println(ub)
	fmt.Println(utf8.DecodeRune(ub))
	fmt.Println(len("Hello, 世界"))

	fmt.Println(len([]rune("Hello, 世界")))

	ip := net.ParseIP("12::ac1d:9003:0:0")
	fmt.Println(ip)
	fmt.Println(net.HardwareAddr(ip[8:]))

	cidr, ipNet, err := net.ParseCIDR("12::ac1d:9003:0:0/64")
	size, _ := ipNet.Mask.Size()
	fmt.Println(cidr, ipNet.IP, size, err)
	fmt.Println(cidr.To16())
	fmt.Println(net.HardwareAddr(ipNet.IP[8:]))
	fmt.Println(len([]rune("lab.casa.com;386824474;2285846018;450051230000000-69;bell.mnc152.mcc502.gprs;;http:__smfsm:80_nchf-convergedcharging_v1_notify_imsi-450051230000000_69")))
}

func SbaBitRateRmToInt(rate string) (r *int64, err error) {

	// format should be '^\d+(\.\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$' such as: "100.00 kbps"
	split := strings.Split(rate, " ")
	if len(split) != 2 {
		err = fmt.Errorf("unsupport rate format: %s", rate)
		return nil, err
	}

	val, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return nil, err
	}

	switch split[1] {
	case "bps":
		rateVal := int64(val)
		r = &rateVal
	case "Kbps":
		rateVal := int64(val * 1000)
		r = &rateVal
	case "Mbps":
		rateVal := int64(val * 1000000)
		r = &rateVal
	case "Gbps":
		rateVal := int64(val * 1000000000)
		r = &rateVal
	case "Tbps":
		rateVal := int64(val * 1000000000000)
		r = &rateVal
	default:
		err = fmt.Errorf("BitRateRm just support bps|Kbps|Mbps|Gbps|Tbps: %s", rate)
	}

	/*if strings.HasSuffix(rate, "bps") {
		r = new(int64)
		fmt.Sscanf(rate, "%d bps", r)
	} else if strings.HasSuffix(rate, "Kbps") {
		r = new(int64)
		if strings.Contains(rate, ".") {
			var f float64
			fmt.Sscanf(rate, "%f Kbps", &f)
			*r = int64(f * 1000)
		} else {
			fmt.Sscanf(rate, "%d Kbps", r)
			*r = *r * 1000
		}
	} else if strings.HasSuffix(rate, "Mbps") {
		r = new(int64)
		if strings.Contains(rate, ".") {
			var f float64
			fmt.Sscanf(rate, "%f Mbps", &f)
			*r = int64(f * 1000000)
		} else {
			fmt.Sscanf(rate, "%d Mbps", r)
			*r = *r * 1000000
		}
	} else if strings.HasSuffix(rate, "Gbps") {
		r = new(int64)
		if strings.Contains(rate, ".") {
			var f float64
			fmt.Sscanf(rate, "%f Gbps", &f)
			*r = int64(f * 1000000000)
		} else {
			fmt.Sscanf(rate, "%d Gbps", r)
			*r = *r * 1000000000
		}
	} else if strings.HasSuffix(rate, "Tbps") {
		r = new(int64)
		if strings.Contains(rate, ".") {
			var f float64
			fmt.Sscanf(rate, "%f Tbps", &f)
			*r = int64(f * 1000000000000)
		} else {
			fmt.Sscanf(rate, "%d Tbps", r)
			*r = *r * 1000000000000
		}
	} else {
		err = fmt.Errorf("BitRateRm just support bps|Kbps|Mbps|Gbps|Tbps: %s", rate)
	}*/

	return
}

// 整形转换成字节
func IntToBytes(val int64) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, val); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

// 字节转换成整形
func BytesToInt(b []byte) (int64, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var iVal int64
	if err := binary.Read(bytesBuffer, binary.BigEndian, &iVal); err != nil {
		return -1, err
	}
	return iVal, nil
}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func bitRateToHexEncode(rate string) string {
	split := strings.Split(rate, " ")
	if len(split) != 2 {
		return ""
	}

	val, err := strconv.ParseFloat(split[0], 64)
	if err != nil {
		return ""
	}
	var s string
	switch split[1] {
	case "bps":
		s = strconv.Itoa(int(val))
	case "Kbps":
		s = strconv.Itoa(int(val / 1000))
		fmt.Println(s)
	case "Mbps":
		s = strconv.Itoa(int(val * 1000000))
	case "Gbps":
		s = strconv.Itoa(int(val * 1000000000))
	case "Tbps":
		s = strconv.Itoa(int(val * 1000000000000))
	default:
		return ""
	}

	b := make([]byte, 4)
	e, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return ""
	}
	binary.BigEndian.PutUint32(b, uint32(e))
	return hex.EncodeToString(b)
}

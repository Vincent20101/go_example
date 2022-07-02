package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//dstring, err := ImsiToTBCDstring("490154203237518")
	dstring, err := ImsiToTBCDstring("630015002000004")
	fmt.Println(dstring, err)

	fmt.Println(TBCDToString(dstring))
	i2, _ := strconv.ParseInt("F4", 16, 0)
	fmt.Println(i2)
	i3, _ := strconv.ParseInt("54", 16, 10)
	i4, _ := strconv.ParseInt("00", 16, 10)
	i5, _ := strconv.ParseInt("15", 16, 10)
	i6, _ := strconv.ParseInt("32", 16, 10)
	i7, _ := strconv.ParseInt("00", 16, 10)
	i8, _ := strconv.ParseInt("00", 16, 10)
	i9, _ := strconv.ParseInt("00", 16, 10)
	i0, _ := strconv.ParseInt("F0", 16, 10)
	fmt.Println(TBCDToString([]byte{byte(i3), byte(i4), byte(i5), byte(i6), byte(i7), byte(i8), byte(i9), byte(i0)}))
	fmt.Println(TBCDToString([]byte{84, 0, 21, 50, 0, 0, 0, 240}))

	parseInt, _ := strconv.ParseInt("148", 10, 0)
	formatInt := strconv.FormatInt(int64(parseInt), 16)
	fmt.Println(formatInt)
	fmt.Println(solve(formatInt))

	i, _ := strconv.ParseInt("94", 16, 64)
	fmt.Println(i)
	//fmt.Println(IntToBytes(490154203237518))
	//fmt.Println(BytesToInt(dstring))
	//
	//encmsg, err1 := EncodeToTBCD("12345")
	//fmt.Println([]byte(encmsg))
	//dedmsg, err2 := DecodeToTBCD(encmsg)
	//
	//if err1 == nil && err2 == nil {
	//	fmt.Println("Encoded Message", encmsg)
	//	fmt.Println("Decoded Message", dedmsg)
	//} else {
	//	fmt.Println("Encoded and Decoded message not matched")
	//}
	//
	//fmt.Println(btox(encmsg))

	//toInt := BytesToInt([]byte{16})
	//fmt.Println(toInt)
	//toString := hex.EncodeToString(dstring)
	//fmt.Println(toString)

	//bs, _ := hex.DecodeString("0fa8")
	//num := binary.BigEndian.Uint16(bs[:2])
	//fmt.Println(num)

	marshal, _ := json.Marshal([]byte{84, 0, 21, 50, 0, 0, 0, 240})
	fmt.Println(string(marshal))
}

func solve(str string) string {
	// write code here
	if len(str) == 0 {
		return ""
	}
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func Bcd2Number(bcd []byte) string {
	var number string
	for _, i := range bcd {
		number += fmt.Sprintf("%02X", i)
	}
	pos := strings.LastIndex(number, "F")
	if pos == 8 {
		return "0"
	}
	return number[pos+1:]
}

//二进制转十六进制
func btox(b string) string {
	base, _ := strconv.ParseInt(b, 2, 10)
	return strconv.FormatInt(base, 16)
}

//十六进制转二进制
func xtob(x string) string {
	base, _ := strconv.ParseInt(x, 16, 10)
	return strconv.FormatInt(base, 2)
}

func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

// convert string to TBCD string
func ImsiToTBCDstring(imsi string) (b []byte, err error) {

	if len(imsi) > 0 {
		// IMSI ::= TBCDSTRING (SIZE(3..8))
		if len(imsi) > 15 || len(imsi) < 5 {
			err = fmt.Errorf("invalid imsi length")
			return nil, err
		}
		if len(imsi)%2 == 0 {
			b = make([]byte, len(imsi)/2)
			for i := 0; i < len(imsi); {
				b[i/2] = imsi[i] & 0x0f
				i++
				b[i/2] |= imsi[i] << 4
				i++
			}
		} else {
			b = make([]byte, (len(imsi)+1)/2)
			i := 0
			for i < len(imsi)-1 {
				b[i/2] = imsi[i] & 0x0f
				i++
				b[i/2] |= imsi[i] << 4
				i++
			}
			b[i/2] = imsi[i] & 0x0f
			i++
			b[i/2] |= 0xf0
		}

	} else {
		err = fmt.Errorf("imsi is nil")
		return nil, err
	}

	return
}

func TBCDToString(buf []byte) string {
	var Str string
	if len(buf) == 0 {
		return Str
	}

	for _, data := range buf {
		str1 := strconv.Itoa(int(data & 0x0f))        // 0x0f  -241   https://tool.lu/hexconvert/?t=1506303877199
		str2 := strconv.Itoa(int((data & 0xf0) >> 4)) // 0xf0  -16    https://tool.lu/hexconvert/?t=1506303877199

		if len(str2) == 1 {
			Str = Str + str1 + str2
		} else {
			Str = Str + str1
		}

	}

	return Str
}

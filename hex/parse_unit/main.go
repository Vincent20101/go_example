package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math"
	"strconv"
)

func main() {
	var i int64
	toBytes, _ := IntToBytes(i)
	fmt.Println(hex.EncodeToString(toBytes))

	//ss := "2C06000003E8000003E8"
	in := 655350
	bb, _ := IntToBytes(int64(in))
	fmt.Println(hex.EncodeToString(bb))
	decodeString, _ := hex.DecodeString("000000000009fff6")
	fmt.Println(decodeString)
	fmt.Println(BytesToInt(decodeString))

	u1 := int64(0) << 24
	u2 := int64(0) << 16
	u3 := int64(3) << 8
	u4 := int64(232)
	fmt.Println(u1 | u2 | u3 | u4)
	fmt.Println(strconv.FormatInt(u1, 2))
	fmt.Println(strconv.FormatInt(u2, 2))
	fmt.Println(strconv.FormatInt(u3, 2))
	fmt.Println(strconv.FormatInt(u4, 2))

	fmt.Println(hex.DecodeString("e8"))
	fmt.Println([]byte{255})
	fmt.Println('0')
	fmt.Println(hex.DecodeString("68"))
	fmt.Println([]byte("68"))
	fmt.Println([]byte("16i668"))
	s1 := "1111111111111111111"
	fmt.Println([]byte(s1))
	fmt.Println(s1[9])
	fmt.Println((6 << 4) | 8)

	parseUint, _ := strconv.ParseUint("104", 10, 64)
	fmt.Println(parseUint)
	bi := make([]byte, 4)
	binary.BigEndian.PutUint32(bi, 1040000000)
	fmt.Println(bi)
	maxInt32 := math.MaxInt32
	fmt.Println(maxInt32)
	fmt.Println(math.MaxUint32)
}

//整形转换成字节
func IntToBytes(val int64) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	if err := binary.Write(bytesBuffer, binary.BigEndian, val); err != nil {
		return nil, err
	}
	return bytesBuffer.Bytes(), nil
}

//字节转换成整形
func BytesToInt(b []byte) (int64, error) {
	bytesBuffer := bytes.NewBuffer(b)
	var iVal int64
	if err := binary.Read(bytesBuffer, binary.BigEndian, &iVal); err != nil {
		return -1, err
	}
	return iVal, nil
}

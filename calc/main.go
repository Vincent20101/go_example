package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
)

func main() {
	// 十进制
	var a int = 10
	a2 := strconv.FormatInt(int64(a), 2)   //10进制转换为2进制
	fmt.Printf("%v \n", a2)                //1010
	a8 := strconv.FormatInt(int64(a), 8)   //10进制转换为8进制
	fmt.Printf("%v \n", a8)                //12
	a16 := strconv.FormatInt(int64(a), 16) //10进制转换为16进制
	fmt.Printf("%v \n", a16)               //a
	var b string = "1010"
	p, _ := strconv.ParseInt(b, 2, 64) // 2进制转10进制
	fmt.Printf("%v\n", p)              //10
	var b2 string = "345"
	p2, _ := strconv.ParseInt(b2, 8, 64) // 8进制转10进制
	fmt.Printf("%v\n", p2)               //229

	toBytes := IntToBytes(257)
	fmt.Println(toBytes)
	fmt.Println(BytesToInt(toBytes))

	var v3 uint32
	var b3 [4]byte
	v3 = 257
	b3[3] = uint8(v3)
	b3[2] = uint8(v3 >> 8)
	b3[1] = uint8(v3 >> 16)
	b3[0] = uint8(v3 >> 24)
	fmt.Println("b3:", b3)
}

//整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

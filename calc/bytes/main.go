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
	var i int64 = 23423423430000
	fmt.Println(i)
	a2 := strconv.FormatInt(i, 2)
	fmt.Println(a2)

	//offsetLow := uint32(offset & 0xffffffff)
	//offsetHigh := uint32((offset >> 32) & 0xffffffff)
	fmt.Println(strconv.FormatInt(i&0xffffffff, 2))
	fmt.Println(strconv.FormatInt((i>>32)&0xffffffff, 2))
	fmt.Println("二进制：")
	fmt.Printf("%b\n", math.Float64bits(52.0))
	fmt.Println("二进制：end")

	fmt.Println(strconv.ParseInt("000000000111010100110000", 2, 64))
	i = 30000
	a3 := strconv.FormatInt(i, 2)
	fmt.Println(a3)
	fmt.Println(strconv.FormatInt(i&0xff, 2))
	fmt.Println(strconv.FormatInt((i>>8)&0xff, 2))

	bs := []byte{0x00, 0xfD}
	for _, n := range bs {
		fmt.Printf("%08b", n) // prints 00000000 11111101
	}

	fmt.Printf("%b\n", math.Float64bits(52.0))

	fmt.Println(strconv.ParseInt("02", 16, 10))
	fmt.Println(strconv.ParseInt("AF", 16, 10))
	fmt.Println(strconv.ParseInt("08", 16, 10))
	fmt.Println(strconv.ParseInt("73", 16, 10))

	fmt.Println(hex.DecodeString("1111"))

	fmt.Println(strconv.ParseInt("16", 16, 8))
}

// 整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)

	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// 字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)

	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)

	return int(x)
}

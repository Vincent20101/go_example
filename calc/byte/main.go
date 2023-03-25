package main

import (
	"fmt"
	"math/big"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println(strconv.FormatInt(100, 2))
	fmt.Println(strconv.FormatInt(10, 2))
	fmt.Println(strconv.FormatInt((100^10)&10, 2))
	var i uint64
	i = 234234234234234234
	formatBit := strconv.FormatInt(int64(i), 2)
	fmt.Println("i的二进制：", formatBit)
	zy32 := i << 32
	fmt.Printf("zy32: %v, type: %T\n", zy32, zy32)
	fmt.Println("zy32的二进制：", strconv.FormatInt(int64(zy32), 2))
	zy24 := i << 24
	fmt.Printf("zy24: %v, type: %T\n", zy24, zy24)
	fmt.Println("zy24的二进制：", strconv.FormatInt(int64(zy24), 2))
	zy8 := i << 8
	fmt.Printf("zy8: %v, type: %T\n", zy8, zy8)
	fmt.Println("zy8的二进制：", strconv.FormatInt(int64(zy8), 2))
	tt := i<<32 | i<<24 | i<<16 | i<<8 | i
	fmt.Println(strconv.FormatInt(int64(tt), 2))

	big.NewInt(1)
	var ss string
	fmt.Println(unsafe.Sizeof(ss))
	fmt.Println(unsafe.Sizeof(i))
	data := "1000"
	bytes := *(*[]byte)(unsafe.Pointer(&data))
	fmt.Println(string(bytes))
	fmt.Println([]byte(data))
	fmt.Println([]byte(data))

	ti := 10000
	bb := make([]byte, 5)
	bb[0] = byte(ti)
	bb[1] = byte(ti >> 8)
	bb[2] = byte(ti >> 16)
	bb[3] = byte(ti >> 24)
	bb[4] = byte(ti >> 32)

	var is uint16
	is = 300
	fmt.Println(strconv.FormatInt(int64(is), 2))
	fmt.Println(strconv.FormatInt(int64(is&0xf), 2))
	fmt.Printf("%T\n", is&0xf)
	fmt.Println(is & 0xf)
	fmt.Println(is & 0x0f)

	fmt.Println(bb)
	fmt.Println(string(bb))
}

package main

import (
	"bytes"
	"encoding/binary"
	fmt "fmt"
	"log"
	"os"
	"strconv"
	"unsafe"
)

func main() {
	//var s = "123"
	//atoi, err := strconv.Atoi(s)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(atoi)
	//
	log.New(os.Stdout, "\r\n", log.LstdFlags)

	x := 0x12345678

	p := unsafe.Pointer(&x)
	n := (*[4]byte)(p)

	for i := 0; i < len(n); i++ {
		fmt.Printf("%X\n ", n[i])
		//fmt.Printf("%c ", n[i])
	}

	a := 10
	a16 := strconv.FormatInt(int64(a), 16) //10进制转换为16进制
	fmt.Printf("%v \n", a16)               //a

	xs := "12345678"
	parseUint, _ := strconv.ParseUint(xs, 16, 32)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, uint32(parseUint))
	fmt.Println(bytesBuffer.Bytes())
	for _, b := range bytesBuffer.Bytes() {
		fmt.Printf("%x\n", b)
	}

	var body interface{}
	body = &[]int{3}[0]
	fmt.Println(body)
	v, ok := body.(*int)
	fmt.Println(v, ok)

}

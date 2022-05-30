package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

// string转ytes
func Str2sbyte(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s	// 把s的地址付给b
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s)	// 修改容量为长度
	return
}

// []byte转string
func Sbyte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func main() {
	str := "this is a string"
	var bs []byte

	//bs = Str2sbyte(str)
	//fmt.Println(bs)
	//fmt.Println(Sbyte2str(bs))

	bs = string2bytes(str)
	fmt.Println(bs)
	fmt.Println(bytes2string(bs))
}

func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string{
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}

	return *(*string)(unsafe.Pointer(&sh))
}
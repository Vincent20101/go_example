package main

import (
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func main() {
	bytes := []byte{104, 101, 108, 108, 111}
	fmt.Println(string(bytes))
	p := unsafe.Pointer(&bytes) //强制转换成unsafe.Pointer，编译器不会报错
	str := *(*string)(p)        //然后强制转换成string类型的指针，再将这个指针的值当做string类型取出来
	fmt.Println(str)            //输出 "hello"

	s := utf8.RuneCountInString("Hello, 世界")
	s1 := len("Hello, 世界")
	fmt.Println(s, s1)
}

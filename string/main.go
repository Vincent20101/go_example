package main

import (
	bytes2 "bytes"
	"fmt"
	"strconv"
	"strings"
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

	index := strings.Index("死神来了, 死神bye bye", ", ")
	fmt.Println(len("死神来了死神"))
	fmt.Println(index)
	fmt.Println(strings.Index("asdfgh", "g"))

	type MsgType string
	type M2 string
	var ts1 MsgType = "sdf66666"
	fmt.Println(M2(ts1))
	fmt.Println(test())
	var b []byte = make([]byte, 10)
	fmt.Println(string(bytes2.Repeat(b, 10)))
	fmt.Println(len(b))
	fmt.Println(strconv.FormatInt(int64(0|0xff>>1), 2))
	fmt.Printf("%s\n", bytes2.Repeat([]byte("1"), 10))
	fmt.Println([]string{"a", "b"})
	fmt.Printf("ba%s", bytes2.Repeat([]byte("na"), 2))
	//Repeat
	b = []byte("ha")
	fmt.Println(string(bytes2.Repeat(b, 1))) //ha
	fmt.Println(string(bytes2.Repeat(b, 2))) //haha
}

func test() (retry bool) {
	return
}

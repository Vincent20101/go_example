package main

import (
	"fmt"
	"net"
	"time"
	"unsafe"
)

func main() {
	fmt.Println(net.ParseIP("12::/64"))

	fmt.Println(unsafe.Sizeof("aaa"))

	var t time.Time
	fmt.Println(t)
	fmt.Println(t.IsZero())

	var s *string
	s1 := "test"
	s = &s1
	fmt.Println(*s)

	ss := "12::/64"
	byte := []byte(ss)
	fmt.Println(byte)
	fmt.Println(byte[len(byte)-1])
	fmt.Println(string(byte[len(byte)-1]))
}

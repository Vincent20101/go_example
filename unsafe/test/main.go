package main

import (
	"fmt"
	"unsafe"
)

// https://qcrao.com/post/dive-into-go-unsafe/
func double(x *int) {
	*x += *x // x 跟 a 指向同一地址
	x = nil  // x 指向另一个地址
}

func main() {
	var a = 3
	double(&a)
	fmt.Println(a) // 6

	p := &a
	double(p)
	fmt.Println(a, p == nil) // 12 false

	mp := make(map[string]int)
	mp["qcrao"] = 100
	mp["stefno"] = 18

	count := **(**int)(unsafe.Pointer(&mp))
	fmt.Println(count, len(mp)) // 2 2

	s := make([]int, 9, 20)
	var Len = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(8)))
	fmt.Println(Len, len(s)) // 9 9

	var Cap = *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + uintptr(16)))
	fmt.Println(Cap, cap(s)) // 20 20

}

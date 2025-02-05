package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	c := 1
	fmt.Println(c)
	f1(c)
	fmt.Println(c)
	f2(&c)
	fmt.Println(c)

	var a int64
	fmt.Println(&a)
	ptr := &a
	fmt.Println(ptr)
	pptr := &ptr
	fmt.Println(pptr)
	ppptr := &pptr
	fmt.Println(ppptr)
	a = 1
	fmt.Println(**ppptr)
	fmt.Println(***ppptr)

	atomic.AddInt64(&a, 1)
	fmt.Println(**ppptr)
	fmt.Println(***ppptr)
}

func f1(c int) {
	c = 10
}

func f2(c *int) {
	*c = 10
	c = nil
}

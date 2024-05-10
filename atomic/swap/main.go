package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var t *uint32 = new(uint32)
	fmt.Println(atomic.CompareAndSwapUint32(t, 1, 2))
	fmt.Println(*t)
}

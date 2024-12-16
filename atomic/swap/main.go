package main

import (
	"bytes"
	"fmt"

	"go.uber.org/atomic"
)

func main() {
	//var t *uint32 = new(uint32)
	//fmt.Println(atomic.CompareAndSwapUint32(t, 0, 2))
	//fmt.Println(*t)

	var connState atomic.Bool
	fmt.Println(connState.Load())
	b := connState.CAS(connState.Load(), true)
	fmt.Println(connState.Load())
	fmt.Println(b)

	body := bytes.NewBuffer(nil)
	fmt.Println(body.Bytes())

	var tt, tt1 uint64
	tt = 0x200000
	tt1 = 4611686018832138240
	tt2 := 4611686018830041088
	fmt.Println(int64(tt2))
	fmt.Println(tt1 & tt)
}

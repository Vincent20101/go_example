package main

import (
	"fmt"
	"math"
)

const (
	KBSize = 1024
	MBSize = KBSize * 1024
	GBSize = MBSize * 1024
)

func main() {
	var c int64
	c = int64(1 * GBSize / 4000)
	fmt.Println(c)

	b := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = make([]int, len(b))
	copy(a, b)
	fmt.Println(a)
	fmt.Println(math.MaxInt)
}

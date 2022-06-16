package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for x := 0; x < 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}
	fmt.Println(4 & 15)

	// 49 -->  148
	fmt.Println((4 & 15), (9 << 4))
	fmt.Println((4 & 15) | (9 << 4))
	// 148 --> 49
	fmt.Println((148 & 15), (148 >> 4))
	fmt.Println((15 & 15), (15 >> 4))
}

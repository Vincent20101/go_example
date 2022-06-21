package main

import (
	"fmt"
	"math/rand"
)

const (
	_ = iota

	KB = 1 << (10 * iota)

	MB = 1 << (10 * iota)

	GB = 1 << (10 * iota)

	TB = 1 << (10 * iota)

	PB = 1 << (10 * iota)
)

func main() {
	// 一个巧妙的技巧是：你可以用 & 操作去测试一个数字是奇数还是偶数。原因是当一个数字的二进制的最低位是 1 的时候，那他就是奇数。
	// 我们可以用一个数字和 1 进行 & 操作，然后在和 1 做 AND 运算，如果的到的结果是 1 ，那么这个原始的数字就是奇数
	for x := 0; x < 100; x++ {
		num := rand.Int()
		if num&1 == 1 {
			fmt.Printf("%d is odd\n", num)
		} else {
			fmt.Printf("%d is even\n", num)
		}
	}

	fmt.Println(KB, MB, GB, TB, PB)

	fmt.Println(4 & 15)

	// 49 -->  148
	fmt.Println((4 & 15), (9 << 4))
	fmt.Println((4 & 15) | (9 << 4))
	// 148 --> 49
	fmt.Println((148 & 15), (148 >> 4))
	fmt.Println((15 & 15), (15 >> 4))
}

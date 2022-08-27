package main

import (
	//"crypto/rand"
	"fmt"
	"math/rand"
	"strings"
	"time"
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
	//for x := 0; x < 100; x++ {
	//	num := rand.Int()
	//	if num&1 == 1 {
	//		fmt.Printf("%d is odd\n", num)
	//	} else {
	//		fmt.Printf("%d is even\n", num)
	//	}
	//}

	fmt.Println(KB, MB, GB, TB, PB)

	fmt.Println(4 & 15)

	// 49 -->  148
	fmt.Println((4 & 15), (9 << 4))
	fmt.Println((4 & 15) | (9 << 4))
	// 148 --> 49
	fmt.Println((148 & 15), (148 >> 4))
	fmt.Println((15 & 15), 15>>4)

	fmt.Println(122 | 239)
	fmt.Println(239 | 122)

	unix := time.Now().Unix()
	fmt.Println(unix)
	fmt.Printf("%x\n", unix)
	fmt.Printf("%X\n", unix)
	//r, _ := rand.Int(rand.Reader, big.NewInt(1000))
	//fmt.Println(r)

	i := rand.Int()
	fmt.Println(i)
	fmt.Println(rand.Intn(10))

	fmt.Println(strings.Split("vzwaudit.mnc010.mcc310.gprs", ".")[0])

	// format should be '^\d+(\.\d+)? (bps|Kbps|Mbps|Gbps|Tbps)$' such as: "100.00 kbps"
	split := strings.Split("100.00 kbps", " ")
	fmt.Println(split[0])

	fmt.Println("  sdf sdfdsf  ")
	fmt.Println(strings.TrimSpace("  sdf sdfdsf  "))

}

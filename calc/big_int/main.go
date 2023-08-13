package main

import (
	"fmt"
	"math/big"
)

const NumFlags = 10

type BitFlag struct {
	value big.Int
}

func (b *BitFlag) Set(bit int) {
	if bit < NumFlags {
		b.value.SetBit(&b.value, bit, 1)
	}
}

func (b *BitFlag) Get(bit int) bool {
	if bit < NumFlags {
		return b.value.Bit(bit) == 1
	}
	return false
}

func (b *BitFlag) Clean(bit int) {
	if bit < NumFlags {
		b.value.SetBit(&b.value, bit, 0)
	}
}

func main() {
	var flags BitFlag

	flags.Set(0) // 设置第一位
	flags.Set(3) // 设置第四位
	flags.Set(8) // 设置第九位

	fmt.Println(flags.Get(0)) // 输出 true
	fmt.Println(flags.Get(3)) // 输出 true
	fmt.Println(flags.Get(8)) // 输出 true
	fmt.Println(flags.Get(1)) // 输出 false

	flags.Clean(0) // 清除第一位
	flags.Clean(3) // 清除第四位
	flags.Clean(8) // 清除第九位

	fmt.Println(flags.Get(0)) // 输出 false
	fmt.Println(flags.Get(3)) // 输出 false
	fmt.Println(flags.Get(8)) // 输出 false
}

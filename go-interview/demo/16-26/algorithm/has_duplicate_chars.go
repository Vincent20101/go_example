package main

import (
	"fmt"
)

func hasDuplicateChars(str string) bool {
	//定义长度为8的数组，每个int有32位，共256位
	var bitset [8]int32

	for i := 0; i < len(str); i++ {
		char := str[i]
		// 计算字符对应的数组索引和位索引
		idx := char / 32                     // 数组索引
		bit := char % 32                     // 位索索引
		if (bitset[idx] & (1 << bit)) != 0 { // 检查该位是否已设置
			return true
		}
		bitset[idx] |= (1 << bit) // 设置该位
	}
	return false
}

func main() {
	fmt.Println(hasDuplicateChars("good"))                       // 输出：true
	fmt.Println(hasDuplicateChars("abc"))                        // 输出：false
	fmt.Println(hasDuplicateChars("abcdefghijklmnopqrstuvwxyz")) // 输出：false
}

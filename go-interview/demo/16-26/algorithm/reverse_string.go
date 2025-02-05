package main

import (
	"fmt"
)

// reverseString 反转字符串
func reverseString(s string) string {
	// 将字符串转换为 rune 切片，支持多字节字符
	runes := []rune(s)
	// 定义两个指针，一个指向切片的起始位置，一个指向切片的结束位置
	i, j := 0, len(runes)-1
	// 当左指针小于右指针时，交换两者指向的元素，并向中心移动
	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}
	// 将 rune 切片转换回字符串
	return string(runes)
}

func main() {
	s := "Hello, 世界"
	reversed := reverseString(s)
	fmt.Println("原字符串:", s)
	fmt.Println("反转后的字符串:", reversed)
}

package main

import (
	"fmt"
)

// isAnagram 判断两个字符串是否为换位字符串
func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	counter := make([]int, 128) // UTF-8编码的ASCII字符集基数

	for _, ch := range s1 {
		counter[ch]++
	}
	for _, ch := range s2 {
		counter[ch]--
		if counter[ch] < 0 {
			return false
		}
	}
	return true
}

func main() {
	s1 := "aaaabbc"
	s2 := "abcbaaa"
	fmt.Println("两字符串是否为换位字符串：", isAnagram(s1, s2))
}

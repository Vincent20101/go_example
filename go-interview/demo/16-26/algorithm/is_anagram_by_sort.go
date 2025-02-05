package main

import (
	"fmt"
	"sort"
)

// sortString 对字符串中的字符进行排序
func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

// isAnagramBySort 通过排序判断两个字符串是否为换位字符串
func isAnagramBySort(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return sortString(s1) == sortString(s2)
}

func main() {
	s1 := "aaaabbc"
	s2 := "abcbaaa"
	fmt.Println("两字符串是否为换位字符串（排序法）：", isAnagramBySort(s1, s2))
}

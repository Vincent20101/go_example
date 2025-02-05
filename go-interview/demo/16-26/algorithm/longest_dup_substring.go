package main

import (
	"fmt"
)

// Rabin-Karp 字符串哈希方法
func search(L int, a int, modulus int64, n int, s string) (start int, ok bool) {
	// 计算第一个长度为L的子串的哈希值
	h := int64(0)
	for i := 0; i < L; i++ {
		h = (h*int64(a) + int64(s[i]-'a')) % modulus
	}

	// 存储已经看到的哈希值
	seen := make(map[int64]bool)
	seen[h] = true

	// 常数值用于滚动哈希
	aL := int64(1)
	for i := 1; i <= L; i++ {
		aL = (aL * int64(a)) % modulus
	}

	// 滑动窗口以查找新的哈希值
	for i := 1; i < n-L+1; i++ {
		h = (h*int64(a) - int64(s[i-1]-'a')*aL + int64(s[i+L-1]-'a')) % modulus
		if h < 0 {
			h += modulus
		}
		if seen[h] {
			start = i
			return start, true
		}
		seen[h] = true
	}

	return 0, false
}

// 主函数
func longestDupSubstring(s string) string {
	n := len(s)
	// 搜索最长重复子串的长度范围
	left, right := 1, n-1
	modulus := int64(1<<31) - 1
	a := 26
	start, length := 0, 0

	for left <= right {
		L := left + (right-left)/2
		if pos, ok := search(L, a, modulus, n, s); ok {
			left = L + 1
			start = pos
			length = L
		} else {
			right = L - 1
		}
	}

	return s[start : start+length]
}

func main() {
	fmt.Println("Longest duplicate substring in 'banana' is:", longestDupSubstring("banana")) // 输出：ana
}

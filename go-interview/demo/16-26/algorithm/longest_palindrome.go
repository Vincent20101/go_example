package main

import "fmt"

// 从给定中心扩展找到回文
func expandAroundCenter(s string, left, right int) (int, int) {
	// 向两边扩展，直到不再是回文
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return left + 1, right - 1
}

// 寻找最长的回文子串
func longestPalindrome(s string) string {
	start, end := 0, 0 // 记录最长回文子串的起始和结束位置

	for i := 0; i < len(s); i++ {
		// 奇数长度回文的中心扩展
		l1, r1 := expandAroundCenter(s, i, i)
		// 偶数长度回文的中心扩展
		l2, r2 := expandAroundCenter(s, i, i+1)

		// 更新找到的最长回文子串的位置
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}

	// 提取子串并返回
	return s[start : end+1]
}

func main() {
	s := "babad"
	result := longestPalindrome(s)
	fmt.Println("最长回文子串是:", result)
}

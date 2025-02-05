package main

import (
	"fmt"
)

// FindLexicographicalMaxSubsequence 通过逆序遍历找到字典序最大子序列
func FindLexicographicalMaxSubsequence(s string) string {
	n := len(s)
	if n == 0 {
		return ""
	}

	// 结果数组，最初存储逆序结果
	result := []byte{}

	// 从后向前遍历，记录当前遇到的最大字符
	maxChar := byte(0)
	for i := n - 1; i >= 0; i-- {
		if s[i] >= maxChar {
			maxChar = s[i]
			result = append(result, maxChar)
		}
	}

	// 反转结果数组以得到正确的顺序
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func main() {
	input := "acbdxmng"
	result := FindLexicographicalMaxSubsequence(input)
	fmt.Println("The lexicographical max subsequence is:", result)
}

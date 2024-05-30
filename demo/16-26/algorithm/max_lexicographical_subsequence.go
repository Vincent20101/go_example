package main

//给定一个字符串，求串中字典序最大的子序列。字典序最大的子序列是这样构造的：给定字符串a0a1…an-1，首先在字符串a0a1…an-1找到值最大的字符ai，然后在剩余的字符串ai+1…an-1中找到值最大的字符aj，然后在剩余的aj+1…an-1中找到值最大的字符ak…直到字符串的长度为0，则aiajak…即为答案
import (
	"fmt"
)

func FindMaxLexicographicalSubsequence(s string) string {
	var result []byte
	startIndex := 0
	for startIndex < len(s) {
		maxChar := s[startIndex]
		maxIndex := startIndex
		// 从当前起点开始找最大字符
		for i := startIndex; i < len(s); i++ {
			if s[i] > maxChar {
				maxChar = s[i]
				maxIndex = i
			}
		}
		// 将找到的最大字符加入结果中
		result = append(result, maxChar)
		// 更新起点为找到字符后的一个位置
		startIndex = maxIndex + 1
	}

	return string(result)
}

func main() {
	input := "acbdxmng"
	result := FindMaxLexicographicalSubsequence(input)
	fmt.Printf("The max lexicographical subsequence of '%s' is: %s\n", input, result)
}

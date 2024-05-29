package main

import "fmt"

//动态规划法求最长公共子串
func longestCommonSubstring(s1, s2 string) string {
	len1 := len(s1)
	len2 := len(s2)

	// 初始化一个二维切片来存储动态规划的结果
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	maxLen := 0
	endIndex := 0

	// 动态规划填表
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > maxLen {
					maxLen = dp[i][j]
					endIndex = i // 记录最长公共子串的结束位置
				}
			} else {
				dp[i][j] = 0
			}
		}
	}

	// 从dp表中提取最长公共子串
	longestCommon := s1[endIndex-maxLen : endIndex]
	return longestCommon
}

func main() {
	s1 := "abccade"
	s2 := "dgcadde"
	result := longestCommonSubstring(s1, s2)
	fmt.Println("The longest common substring is:", result)
}

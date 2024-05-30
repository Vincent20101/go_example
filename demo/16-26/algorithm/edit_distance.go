package main

import (
	"fmt"
)

// 计算最小值
func min(values ...int) int {
	minVal := values[0]
	for _, v := range values[1:] {
		if v < minVal {
			minVal = v
		}
	}
	return minVal
}

// 计算编辑距离
func editDistance(s1, s2 string, replaceWeight int) int {
	len1, len2 := len(s1), len(s2)
	dp := make([][]int, len1+1)
	for i := range dp {
		dp[i] = make([]int, len2+1)
	}

	// 初始化
	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}

	// 动态规划计算所有dp值
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(
					dp[i-1][j]+1,               // 删除
					dp[i][j-1]+1,               // 插入
					dp[i-1][j-1]+replaceWeight, // 替换，代价为2
				)
			}
		}
	}

	return dp[len1][len2]
}

func main() {
	//1. 删除位在第1个字的k
	//2. 在第1个字之前加入s
	//3. 删除位在第4个字的e
	//4. 在第4个字之前加入i
	//5. 在第6个字之后加入g
	s1 := "kitten"
	s2 := "sitting"
	fmt.Println("Edit distance:", editDistance(s1, s2, 2)) // 应输出 5
}

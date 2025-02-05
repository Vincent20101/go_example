package main

import (
	"fmt"
)

// 计算最长递增子序列的长度
func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 初始化dp数组，每个元素至少可以构成长度为1的递增子序列
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLIS := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLIS = max(maxLIS, dp[i])
	}

	return maxLIS
}

// 辅助函数，计算两个数的最大值
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println("Length of the longest increasing subsequence is:", lengthOfLIS(nums))
}

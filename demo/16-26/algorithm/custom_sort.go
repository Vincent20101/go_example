package main

import (
	"fmt"
	"sort"
)

// 自定义排序
func customSort(input []string, order string) []string {
	// 创建一个映射，存储每个字母的优先级
	orderMap := make(map[rune]int)
	for index, char := range order {
		orderMap[char] = index
	}

	// 定义比较函数
	lessFunc := func(i, j int) bool {
		s1, s2 := input[i], input[j]
		minLen := len(s1)
		if len(s2) < minLen {
			minLen = len(s2)
		}

		for k := 0; k < minLen; k++ {
			if s1[k] != s2[k] {
				// 比较两个字符的优先级
				return orderMap[rune(s1[k])] < orderMap[rune(s2[k])]
			}
		}
		// 如果所有比较的字符都相同，则较短的字符串应该排在前面
		return len(s1) < len(s2)
	}

	// 使用 sort.Slice 根据自定义比较函数排序
	sort.Slice(input, func(i, j int) bool {
		return lessFunc(i, j)
	})

	return input
}

func main() {
	order := "dgecfboa"
	input := []string{"bed", "dog", "dear", "eye"}
	sortedInput := customSort(input, order)
	for _, str := range sortedInput {
		fmt.Println(str)
	}
}

package main

import "fmt"

//滑动窗口法
func longestCommonSubstringSlideWindow(s1, s2 string) string {
	len1 := len(s1)
	len2 := len(s2)
	longestCommon := ""

	// 遍历较短字符串
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			temp := ""
			tempLen := 0
			// 从当前位置开始比较两个字符串
			for i+tempLen < len1 && j+tempLen < len2 && s1[i+tempLen] == s2[j+tempLen] {
				temp += string(s1[i+tempLen])
				tempLen++
			}
			// 更新最长公共子串
			if len(temp) > len(longestCommon) {
				longestCommon = temp
			}
		}
	}

	return longestCommon
}

func main() {
	s1 := "abccade"
	s2 := "dgcadde"
	result := longestCommonSubstringSlideWindow(s1, s2)
	fmt.Println("The longest common substring is:", result)
}

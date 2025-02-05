package main

//给定一个能判断一个单词是否为另一个单词的子字符串的方法，记为isSubstring。如何判断s2是否能通过旋转s1得到（只能使用一次isSubstring方法）
import (
	"fmt"
	"strings"
)

// 假设isSubstring函数已经给定，这里使用strings包中的Contains函数模拟
func isSubstring(sub, main string) bool {
	return strings.Contains(main, sub)
}

// 判断s2是否能通过旋转s1得到
func canRotateTo(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1s1 := s1 + s1
	return isSubstring(s2, s1s1)
}

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	result := canRotateTo(s1, s2)
	fmt.Println(result) // 输出: true
}

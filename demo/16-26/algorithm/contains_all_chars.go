package main

//给定由字母组成的字符串s1和s2，其中，s2中字母的个数少于s1，如何判断s1是否包含s2？即出现在s2中的字符在s1中都存在。
//例如s1=“abcdef”，s2=“acf”，那么s1就包含s2；如果s1=“abcdef”，s2=“acg”，那么s1就不包含s2，因为s2中有“g”，但是s1中没有g。
import (
	"fmt"
)

// containsAllChars 判断 s1 是否包含 s2 中的所有字符
func containsAllChars(s1, s2 string) bool {
	if len(s2) > len(s1) {
		return false
	}

	counter := make([]int, 128) // ASCII 字符计数器

	// 计数 s1 中的字符
	for _, char := range s1 {
		counter[char]++
	}

	// 检查 s2 中的字符是否都在 s1 中
	for _, char := range s2 {
		if counter[char] == 0 {
			return false // s2 中有字符在 s1 中不存在
		}
	}

	return true
}

func main() {
	s1 := "abcdef"
	s2 := "acf"
	s3 := "acg"

	fmt.Printf("s1 是否包含 s2 的所有字符: %v\n", containsAllChars(s1, s2))
	fmt.Printf("s1 是否包含 s3 的所有字符: %v\n", containsAllChars(s1, s3))
}

package main

import (
	"fmt"
)

// 用递归的方式实现一个求字符串中连续出现相同字符的最大值，
// 例如字符串“aaabbcc”中连续出现字符‘a’的最大值为3，字符串“abbc”中连续出现字符‘b’的最大值为2

// 递归函数
func maxConsecutiveChars(s string, index int, prevChar rune, currentCount int, maxCount int) int {
	if index == len(s) {
		if currentCount > maxCount {
			return currentCount
		}
		return maxCount
	}

	currentChar := rune(s[index])
	if currentChar == prevChar {
		currentCount++
	} else {
		if currentCount > maxCount {
			maxCount = currentCount
		}
		currentCount = 1
	}
	return maxConsecutiveChars(s, index+1, currentChar, currentCount, maxCount)
}

// 入口函数
func maxConsecutive(input string) int {
	if len(input) == 0 {
		return 0
	}
	firstChar := rune(input[0])
	return maxConsecutiveChars(input, 1, firstChar, 1, 0)
}

func main() {
	fmt.Println("Max consecutive 'a':", maxConsecutive("aaabbcc")) // 3
	fmt.Println("Max consecutive 'b':", maxConsecutive("abbc"))    // 2
}

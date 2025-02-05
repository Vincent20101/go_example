package main

//有一个由大小写字母组成的字符串，请对它进行重新组合，使得其中的所有小写字母排在大写字母的前面（大写或小写字母之间不要求保持原来次序）
import (
	"fmt"
)

// partitionLetters 将小写字母移至字符串前部
func partitionLetters(s []rune) []rune {
	left := 0
	for i := range s {
		// 检查当前字符是否为小写字母
		if s[i] >= 'a' && s[i] <= 'z' {
			// 是小写字母时，交换到 left 指针的位置
			s[i], s[left] = s[left], s[i]
			// 移动 left 指针
			left++
		}
	}
	return s
}

func main() {
	s := "AbCDeFgHiJkL"
	fmt.Println("原字符串:", string(s))
	rearranged := partitionLetters([]rune(s)) // 字符串转换为 rune 切片以处理可能的多字节字符
	fmt.Println("重新组合后:", string(rearranged))
}

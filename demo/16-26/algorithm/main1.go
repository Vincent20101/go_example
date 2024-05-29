package main

import "fmt"

// permute 函数用于生成字符串的所有排列
// 参数：
// str: 输入的字符串
// l: 当前处理的字符串的起始索引
// r: 当前处理的字符串的结束索引
// 返回值：无
func permute(str []rune, l, r int) {
	// 如果l等于r，说明已经处理到字符串的最后一个字符，输出当前排列
	if l == r {
		fmt.Println(string(str))
	} else {
		// 从l到r遍历字符串
		for i := l; i <= r; i++ {
			// 交换当前字符和第i个字符
			str[l], str[i] = str[i], str[l]
			// 递归处理剩余的字符串
			permute(str, l+1, r)
			// 恢复原始字符串，以便进行下一次迭代
			str[l], str[i] = str[i], str[l]
		}
	}
}

func main() {
	str := "abcde"
	n := len(str)
	permute([]rune(str), 0, n-1)
}

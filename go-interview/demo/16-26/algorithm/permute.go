package main

import (
	"log"
)

// permute 递归函数以交换字符实现排列
func permute(a []rune, left int, result *map[string]struct{}) {
	if left == len(a)-1 {
		//防止重复字符串生成重复组合
		if _, ok := (*result)[string(a)]; !ok {
			(*result)[string(a)] = struct{}{}
		}

	} else {
		for i := left; i < len(a); i++ {

			//交换left与i位置的字符
			a[left], a[i] = a[i], a[left]
			// 固定第一个字符，对剩余字符进行全排列
			permute(a, left+1, result)
			//还原left与i位置的字符
			a[left], a[i] = a[i], a[left]
		}
	}
}

// permuteString 入口函数，接收字符串参数并初始化递归
func permuteString(s string) map[string]struct{} {
	// 将字符串转换成rune数组以支持Unicode字符
	a := []rune(s)
	result := make(map[string]struct{})
	permute(a, 0, &result)
	return result
}

func main() {
	s := "abb"
	// 获取所有排列
	permutations := permuteString(s)
	// 打印结果
	log.Println(permutations)

}

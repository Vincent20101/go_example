package main

import (
	"fmt"
	"sort"
)

// nextPermutation 寻找下一个排列
func nextPermutation(a []int) bool {
	// Step 1: 从后向前查看逆序区域，找到逆序区域的前一位，也就是字符置换的边界
	i := len(a) - 2
	for i >= 0 && a[i] >= a[i+1] {
		i--
	}
	if i < 0 {
		// 如果是完全逆序，则已经是最后一个排列
		return false
	}

	// Step 2: 在逆序区域中从后向前查找第一个元素，使其大于边界点的值
	j := len(a) - 1
	for a[j] <= a[i] {
		j--
	}

	// Step 3: 交换边界点和找到的元素
	a[i], a[j] = a[j], a[i]

	// Step 4: 将原本的逆序区域转为顺序
	reverse(a, i+1)
	return true
}

// reverse 反转数组的指定部分
func reverse(a []int, start int) {
	i, j := start, len(a)-1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

// allPermutations 获取所有排列
func allPermutations(s string) []string {
	// 先将输入字符串转化为rune切片，并排序
	a := []rune(s)
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })

	// 用于存储所有排列的切片
	var result []string
	// 添加第一个排列
	result = append(result, string(a))

	for {
		// 申请一个新切片，保存当前排列的副本
		b := make([]int, len(a))
		for i, c := range a {
			b[i] = int(c)
		}

		if !nextPermutation(b) {
			break
		}

		// 转换回 rune 切片
		for i, v := range b {
			a[i] = rune(v)
		}
		result = append(result, string(a))
	}

	return result
}

func main() {
	s := "12345"
	permutations := allPermutations(s)
	for _, p := range permutations {
		fmt.Println(p)
	}
}

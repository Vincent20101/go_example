package main

import (
	"fmt"
	"math"
)

// 获取数字的指定位数的值
func getDigit(num, digit int) int {
	return num / int(math.Pow(10, float64(digit-1))) % 10
}

// 获取数字的位数（最大位数）
func getMaxDigits(arr []int) int {
	max := 0
	for _, num := range arr {
		digits := int(math.Log10(float64(num))) + 1
		if digits > max {
			max = digits
		}
	}
	return max
}

// 基数排序
func radixSort(arr []int) []int {
	maxDigits := getMaxDigits(arr)
	fmt.Println("maxDigits: ", maxDigits)
	for digit := 1; digit <= maxDigits; digit++ {
		// 创建10个桶
		buckets := make([][]int, 10)
		// 将数字放入对应的桶中
		for _, num := range arr {
			d := getDigit(num, digit)
			buckets[d] = append(buckets[d], num)
		}
		fmt.Println("buckets:", buckets)
		// 从桶中取出数字重新组成数组
		arrIndex := 0
		for _, bucket := range buckets {
			for _, num := range bucket {
				arr[arrIndex] = num
				arrIndex++
			}
		}
		fmt.Println("arr:", arr)
	}
	return arr
}

func main() {
	arr := []int{170, 45, 75, 90, 802, 24, 3, 2, 66}
	fmt.Println("原始数组：", arr)
	sortedArr := radixSort(arr)
	fmt.Println("排序后数组：", sortedArr)
}

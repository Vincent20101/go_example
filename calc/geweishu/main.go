package main

import "fmt"

func main() {
	num := 1234 // 你要获取位数的数字

	// 获取个位数
	unit := num % 10

	// 获取十位数
	tenth := (num / 10) % 10

	// 获取百位数
	hundredth := (num / 100) % 10

	// 获取千位数
	thousandth := (num / 1000) % 10

	fmt.Printf("千位数：%d\n", thousandth)
	fmt.Printf("百位数：%d\n", hundredth)
	fmt.Printf("十位数：%d\n", tenth)
	fmt.Printf("个位数：%d\n", unit)
}

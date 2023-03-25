package main

import (
	"fmt"
)

//  使用第三变量交换a,b值
func swap(a *int, b *int) {
	tem := *a
	*a = *b
	*b = tem
	return
}

//  使用第三变量交换a,b值:go 直接交换值
func swapTwo(a *int, b *int) {
	*a, *b = *b, *a
}

//  不使用第三变量交换a,b值：直接返回
func swapReturn(a int, b int) (int, int) {
	return b, a
}

//  不使用第三变量交换a,b值：数学运算
func swapThree(a *int, b *int) {
	*a = *a + *b
	*b = *a - *b
	*a = *a - *b
}

//  不使用第三变量交换a,b值：位异或运算
func swapFour(a *int, b *int) {
	*a = *a ^ *b
	*b = *a ^ *b
	*a = *a ^ *b
}

func main() {

	//  第三参数交换值
	f := 3
	s := 4
	swap(&f, &s)
	fmt.Println(f, s)

	// 直接 ruturn
	g := 3
	u := 4
	fmt.Println(swapReturn(g, u))

	//  go 直接交换值
	m := 3
	n := 4
	swapTwo(&m, &n)
	fmt.Println(m, n)

	//  数学运算交换值
	a := 3
	b := 4
	swapThree(&a, &b)
	fmt.Println(a, b)

	//  位异或运算交换值
	h := 3
	w := 4
	swapFour(&h, &w)
	fmt.Println(h, w)
}

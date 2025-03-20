package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Max[T constraints.Ordered](slice []T) T {
	if len(slice) == 0 {
		var zero T
		return zero
	}
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}

// 解释：
//
// 1）函数 Max 使用了类型参数 T， constraints.Ordered 约束，它确保了类型 T 是一个有序类型，因此可以使用 > 运算符进行比较。这样，代码就可以正确编译和运行了。
//
// 2）函数可以接受任何可比较类型的切片，并返回其中的最大值。
//
// 3）在 main 函数中，我们分别传入 int、float64 和 string 类型的切片，展示了泛型函数的通用性。
func main() {
	ints := []int{1, 3, 2, 5, 4}
	floats := []float64{1.1, 3.3, 2.2, 5.5, 4.4}
	strings := []string{"apple", "banana", "cherry"}

	fmt.Println("Max int:", Max(ints))
	fmt.Println("Max float:", Max(floats))
	fmt.Println("Max string:", Max(strings))
}

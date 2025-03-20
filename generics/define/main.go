package main

import (
	"fmt"
	"os"

	"golang.org/x/exp/constraints"
)

type AliasSupiMapOption = SupiMapOption

type PcapCtxOption interface {
	pcapCtxOption()
}

type SupiMapOption struct {
	IsOnlyForSupi bool
	IsUpdate      bool
}

func (SupiMapOption) pcapCtxOption() {}

// 定义一个新的接口，用于约束类型为 SupiMapOption 或其底层类型
type SupiMapOptionLike interface {
	PcapCtxOption
	SupiMapOption // 约束底层类型为 SupiMapOption
	//~SupiMapOption // 约束底层类型为 SupiMapOption
}

type pcapTrcMgr[T PcapCtxOption] struct{}

func GMin[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

func Min[T constraints.Ordered](s []T) T {
	if len(s) == 0 {
		panic("empty slice")
	}

	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}

func main() {
	fmt.Println(GMin[int](2, 3))
	intSlice := []int{3, 1, 4, 2}
	floatSlice := []float64{3.14, 1.61, 4.67, 2.71}

	fmt.Println(Min[int](intSlice))       // Output: 1
	fmt.Println(Min[float64](floatSlice)) // Output: 1.61
	fmt.Println(os.Hostname())
}

package main

import (
	"fmt"
	"sort"
)

// SliceFn implements sort.Interface for a slice of T.
type SliceFn[T any] struct {
	s    []T
	less func(T, T) bool
}

func (s SliceFn[T]) Len() int {
	return len(s.s)
}
func (s SliceFn[T]) Swap(i, j int) {
	s.s[i], s.s[j] = s.s[j], s.s[i]
}
func (s SliceFn[T]) Less(i, j int) bool {
	return s.less(s.s[i], s.s[j])
}

// SortFn sorts s in place using a comparison function.
func SortFn[T any](s []T, less func(T, T) bool) {
	sort.Sort(SliceFn[T]{s, less})
}

// Person struct for demonstration
type Person struct {
	Name string
	Age  int
}

func main() {
	var mm map[string]int

	if err, ok := mm["name"]; ok {
		fmt.Println(err, ok)
	} else {
		fmt.Println("error: ", err, ok)
	}

	// 示例 1: 对整数切片进行排序
	ints := []int{5, 2, 3, 1, 4}
	fmt.Println("Before sorting:", ints)

	SortFn(ints, func(a, b int) bool {
		return a < b // 升序排序
	})

	fmt.Println("After sorting:", ints)

	// 示例 2: 对自定义结构体切片进行排序
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
	}

	fmt.Println("Before sorting:", people)

	SortFn(people, func(a, b Person) bool {
		return a.Age < b.Age // 按年龄升序排序
	})

	fmt.Println("After sorting:", people)

	strings := []string{"apple", "banana", "cherry", "date"}
	SortFn(strings, func(a, b string) bool {
		return a > b // Less function for descending order
	})
	fmt.Println("Sorted strings in reverse order:", strings)
}

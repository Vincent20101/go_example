package main

import (
	"fmt"
	"runtime"
	"sort"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(0))
	// 创建一个字符串切片
	names := []string{"Alice", "Bob", "Eve", "Charlie", "David"}

	// 使用 sort.Slice 自定义排序
	sort.Slice(names, func(i, j int) bool {
		return len(names[i]) < len(names[j])
	})

	// 打印排序后的结果
	fmt.Println("Sorted names by length:", names)
}

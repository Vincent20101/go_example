package main

import (
	"fmt"
	"strings"
)

// 用于判断两个值的最小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 计算两个文件路径的相对路径
func relativePath(a, b string) string {
	aParts := strings.Split(a, "/")
	bParts := strings.Split(b, "/")
	i := 0

	// 找出公共前缀的长度
	for i < min(len(aParts), len(bParts)) && aParts[i] == bParts[i] {
		i++
	}

	// 构造返回上层目录的部分
	upLevels := len(aParts) - i - 1
	relPath := strings.Repeat("../", upLevels)

	// 添加b路径的不同部分
	relPath += strings.Join(bParts[i:], "/")

	return relPath
}

func main() {
	a := "/qihoo/app/a/b/c/d/new.c"
	b := "/qihoo/app/1/2/test.c"
	fmt.Println("Relative path from a to b is:", relativePath(a, b)) // "../../../../1/2/test.c"
}

package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个时间对象，未设置任何值，即 Go 语言的零值
	var t time.Time

	// 判断时间对象是否为零值
	if t.IsZero() {
		fmt.Println("Time is zero") // Output: Time is zero
	} else {
		fmt.Println("Time is not zero")
	}

	// 设置时间对象的值
	t = time.Now()

	// 再次判断时间对象是否为零值
	if t.IsZero() {
		fmt.Println("Time is zero")
	} else {
		fmt.Println("Time is not zero") // Output: Time is not zero
	}

	// 假设时间戳的整数值为 0
	timestamp := int64(0)

	// 使用 time.Unix() 将时间戳转换为 time.Time 对象
	t = time.Unix(timestamp, 0)

	fmt.Println(t)
	fmt.Println(t.IsZero())
	fmt.Println(time.Unix(1, 0))
}

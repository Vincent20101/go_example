package main

import (
	"fmt"
	"time"
	"unsafe"
)

func main() {
	// 获取当前时间
	currentTime := time.Now()

	// 假设你有另一个时间值 otherTime
	otherTime := time.Now().Add(-10 * time.Minute) // 假设 otherTime 是当前时间的10分钟之前
	afterTime := time.Now().Add(10 * time.Minute)
	//var otherTime time.Time

	// 计算两个时间的差值
	timeDiff := otherTime.Sub(currentTime)
	fmt.Println("time:", timeDiff)
	fmt.Println(time.Until(afterTime))
	// 判断差值是否大于5分钟
	if timeDiff > 5*time.Minute {
		fmt.Println("时间差大于5分钟")
	} else {
		fmt.Println("时间差小于或等于5分钟")
	}

}

func StringToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(
		&struct {
			string
			Cap int
		}{s, len(s)},
	))
}

func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

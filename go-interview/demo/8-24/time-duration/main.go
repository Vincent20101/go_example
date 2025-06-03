package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	//afterUnixMilli1()
	//afterUnixMilli2()
	timeStr := "2020-01-02T15:04:12Z"

	// 解析时间字符串
	layout := time.RFC3339 // 使用 RFC3339 格式
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}
	fmt.Println(t)
	fmt.Println(strings.HasPrefix(t.String(), "2020-01-02 15:04"))
	baseTime, _ := time.Parse(time.RFC3339, "2022-01-01T01:00:00Z")
	fmt.Println(baseTime)
	diff := t.Unix() - baseTime.Unix()
	t1 := time.Now().Add(time.Duration(diff) * time.Second)
	fmt.Println(t1)

	fmt.Println(t.Second())
	t2 := time.Now().Add(time.Duration(t.Second()) * time.Second)
	fmt.Println(t2)
}

// 获取500ms后的时间戳，单位ms
func afterUnixMilli1() {
	nowUnixMilli := time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).UnixMilli()
	fmt.Println("afterUnixMilli1: now", nowUnixMilli)
	//延迟后的时间戳 = 指定的时间戳+延时时间
	end := nowUnixMilli + int64(time.Millisecond*500)
	end1 := time.Duration(nowUnixMilli) + time.Millisecond*500
	fmt.Println(end)
	fmt.Println(int64(end1), end1)
}

// 获取500ms后的执行戳，单位ms
func afterUnixMilli2() {
	nowUnixMilli := time.Date(2023, 10, 10, 0, 0, 0, 0, time.UTC).UnixMilli()
	end := nowUnixMilli + 500
	delayduration := 500 * time.Millisecond
	end1 := nowUnixMilli + delayduration.Milliseconds()
	fmt.Println(end)
	fmt.Println(end1)
}

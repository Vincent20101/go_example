package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(strconv.ParseFloat(0, 64))
	fmt.Println(time.Until(time.Now().Add(time.Second * 10)).Milliseconds())
	task := []string{"a", "b", "c", "d", "e", "f"}

	var m = make(map[int]*string)
	for k, t := range task {
		m[k] = &t
	}
	fmt.Println(m)
	//Process1(task)
	Process2(task)
	time.Sleep(time.Hour)
}

func Process1(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func() {
			fmt.Printf("Worker start process task: %s\n", task)
		}()
	}
}

func Process2(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func(t string) {
			fmt.Printf("Worker start process task: %s\n", t)
		}(task)
	}
}

func Double(a int) int {
	return a * 2
}

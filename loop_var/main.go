package main

import (
	"fmt"
	"time"
)

func main() {
	// Process1([]string{"a","b","c","d"})
	Process2([]string{"a","b","c","d"})
}

func Process1(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func() {
			fmt.Printf("Worker start process task: %s\n", task)
		}()
	}
	time.Sleep(time.Second * 10)
}

func Process2(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func(t string) {
			fmt.Printf("Worker start process task: %s\n", t)
		}(task)
	}
	time.Sleep(time.Second * 10)
}

func Double(a int) int {
	return a * 2
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var flag atomic.Bool
	var wg sync.WaitGroup

	// 启动多个 Goroutine 并发读取
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			value := flag.Load()
			fmt.Println("Read value:", value)
		}()
	}

	// 并发写入
	flag.Store(true)

	// 等待所有 Goroutine 完成
	wg.Wait()
}

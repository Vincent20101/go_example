package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/groupcache/singleflight"
)

func main() {
	// 创建一个 SingleFlight 实例
	var group singleflight.Group

	// 使用 WaitGroup 来等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 模拟并发调用的场景
	for i := 1; i <= 15; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// key 用于标识唯一的调用
			key := "example-key"

			// 使用 Do 方法来确保只执行一次函数
			val, err := group.Do(key, func() (interface{}, error) {
				// 模拟一些耗时的操作
				time.Sleep(2 * time.Second)

				// 返回一个结果
				return fmt.Sprintf("Result for ID %d", id), nil
			})

			// 输出结果或错误
			if err != nil {
				fmt.Printf("Error for ID %d: %v\n", id, err)
			} else {
				fmt.Printf("ID %d: %v\n", id, val)
			}
		}(i)
	}

	// 等待所有 goroutine 完成
	wg.Wait()
}

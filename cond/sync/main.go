package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]int, 0)

	// 生产者协程
	go func() {
		for i := 1; i <= 10; i++ {
			cond.L.Lock()
			queue = append(queue, i)
			cond.Signal() // 发送信号给等待的协程
			cond.L.Unlock()
			fmt.Println("push end")
			//time.Sleep(time.Second * 2) // 模拟生产速度
		}
	}()

	// 消费者协程
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for {
				cond.L.Lock()

				// 判断队列是否为空，如果为空则等待信号
				for len(queue) == 0 {
					cond.Wait() // 等待信号
				}
				fmt.Println("goroutine： ", i)
				// 取出队列中的第一个元素
				item := queue[0]
				queue = queue[1:]

				cond.L.Unlock()

				fmt.Printf("Goroutine received item: %d\n", item)

				// 模拟消费速度
				//time.Sleep(time.Second)

				// 判断队列是否已经全部消费完，如果是则退出循环
				cond.L.Lock()
				if len(queue) == 0 {
					cond.L.Unlock()
					continue
				}
				cond.L.Unlock()
			}
		}(i)
	}

	wg.Wait() // 等待所有协程执行完毕
	//time.Sleep(time.Second * 100)
}

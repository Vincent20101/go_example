package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]int, 0)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for {
				//runtime.Gosched()

				if len(queue) == 0 {
					fmt.Println("waiting...")
				}
				fmt.Println("go routing: ", i)
				cond.L.Lock()
				for len(queue) == 0 {
					cond.Wait() // 等待条件满足
				}
				fmt.Println("Goroutine received signal")
				fmt.Println(queue[0])
				queue = queue[1:]
				cond.L.Unlock()
				time.Sleep(time.Second * 2)

				fmt.Println("sleep 2")
			}
		}(i)
	}

	go func() {
		//time.Sleep(time.Second * 10)
		for i := 0; i < 50; i++ {
			cond.L.Lock()
			queue = append(queue, i)
			cond.Signal() // 发送信号给一个等待的协程
			cond.L.Unlock()
			//time.Sleep(time.Second * 2)
		}
	}()

	//time.Sleep(time.Second * 3) // 确保所有协程进入等待状态

	wg.Wait() // 等待所有协程执行完毕
	//time.Sleep(time.Second * 20)
}

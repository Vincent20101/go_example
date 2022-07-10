package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(in1, in2 <-chan int) {
	defer wg.Done()
	for {
		select {
		case v, ok := <-in1:
			if !ok {
				fmt.Println("收到退出信号 in1")
				in1 = nil
			}
			// do something
			fmt.Println(v)
		case v, ok := <-in2:
			if !ok {
				fmt.Println("收到退出信号 in2")
				in2 = nil
			}
			// do something
			fmt.Println(v)
		}
		// select已经结束，我们需要判断两个通道的状态
		// 都为nil则结束当前goroutine
		if in1 == nil && in2 == nil {
			return
		}
	}
}
func main() {
	in1 := make(chan int) // 退出通道，接收
	in2 := make(chan int)
	wg.Add(2)
	go worker(in1, in2)
	go worker(in2, in2)
	for i := 0; i < 3; i++ {
		in1 <- i
		time.Sleep(1 * time.Second)
		in2 <- i
	}
	close(in1)
	close(in2)
	wg.Wait()

	time.Sleep(100 * time.Second)
}

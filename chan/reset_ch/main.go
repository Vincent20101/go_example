package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	// 启动一个 Goroutine 来发送数据到通道
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	// 从通道中读取数据
	for value := range ch {
		fmt.Println("Received:", value)
	}

	// 通道已关闭，创建一个新的通道
	ch = make(chan int)

	// 启动一个 Goroutine 来发送数据到新通道
	go func() {
		ch <- 42
		close(ch)
	}()

	// 从新通道中读取数据
	value := <-ch
	fmt.Println("Received from new channel:", value)
}

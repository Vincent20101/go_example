package main

import (
	"fmt"
	"runtime"
	"sync"

	_ "go.uber.org/automaxprocs"
)

var wg sync.WaitGroup

func main() { // main => go
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(1) // 设置逻辑CPU的个数
	ch := make(chan int)  // 缓冲区
	wg.Add(1)
	go send(ch)
	wg.Wait()
	// time.Sleep(1e9)
}
func send(ch chan int) { // 通道参数引用
	defer wg.Done()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go recv(ch)
		ch <- i
	}
}

func recv(ch chan int) { // 通道参数引用
	defer wg.Done()
	fmt.Println("接收通道信息", <-ch) // 通道没有数据？取不出 =》 阻塞
}

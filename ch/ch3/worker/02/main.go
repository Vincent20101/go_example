package main

import (
	"fmt"
)

//有两个协程方法；方法send可以循环发送指定范围的数字，方法recv可以接收并打印出send新的；并且main方法必须等两个协程完成后再结束

func main() {
	ch := make(chan int)
	done := make(chan bool)

	// 阻塞之后是否影响到 main 运行
	// 死锁的情况 1
	// <-done // 发生阻塞 ==》 阻塞main运行 -》不可取
	// go recv(ch, done)

	// 死锁的情况 2
	go send(1, 4, ch) // 执行 1~ 4 就会结束执行
	go recv(ch, done) // 因为采用循环 =》从通道中获取信息 =》 后面的done通道无法设置参数
	<-done            // 就不能执行 =》就会影响到main

	//time.Sleep(1e9)
}

func send(start, end int, ch chan<- int) {
	for i := start; i < end; i++ {
		ch <- i
	}
	close(ch)
}

func recv(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Println("num : ", num)
	}
	done <- true
}

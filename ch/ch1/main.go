package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var sendChOnly chan<- string // 单向发送通道
var recvChOnly <-chan string // 单向获取通道

// main => 协程
// main 结束之后内部所有的协程都是不会再运行
// 协程切换条件？ 阻塞
func main() {
	// go func() {
	// 	fmt.Println("1")
	// 	time.Sleep(time.Second * 2)
	// 	fmt.Println("y")
	// }()
	// r := 0
	// go func(i int) {
	// 	fmt.Println("2")
	// 	fmt.Println("ko")
	// }(r)
	//
	// time.Sleep(time.Second * 2)
	// fmt.Println("main")

	//go demoRun1()
	// go demoRun2()
	// runtime.GOMAXPROCS(3)
	//
	// go B()
	// go C()
	// go A()
	//
	// time.Sleep(time.Second)
	// []value (0)

	// []类型
	// chan 类型

	// 建议：通道变量名 变量声明具有 ch | chan
	// ch := make(chan string) // 创建通道 - chan string 字符串类型的通道
	// go send1(ch)
	// go recv1(ch)
	//
	// time.Sleep(time.Second)

	// ch := make(chan int, 1)
	// ch <- 10
	// x := <-ch
	// // xs := <-ch
	// fmt.Println("x", x)
	// fmt.Println("x", xs)

	// ch := make(chan string, 1)
	//
	// go recv(ch)
	// go recv(ch)
	// ch <- "shineyork 666"
	// ch <- "shineyork 666"
	// // fmt.Println(<-ch)
	// ch <- "shineyork 666"
	// fmt.Println("ok")

	// ch := make(chan string)
	// ch <- "10"
	// go recv(ch)
	// fmt.Println("ok")

	// ch := make(chan string)
	// go recv(ch)
	// ch <- "shineyork 666"
	// close(ch)
	// v, ok := <-ch
	// fmt.Println("v", v, "ok", ok)

	ch := make(chan string, 1) // 1 声明通道容量 、 缓存
	// sendChOnly = ch
	// sendChOnly <- "shineyork"
	// recvChOnly = ch
	// fmt.Println(<-recvChOnly)

	go sendfChOnly(ch)
	go recvfChOnly(ch)
	time.Sleep(time.Second)
}

func sendfChOnly(ch chan<- string) {
	ch <- "shineyork" //
}

func recvfChOnly(ch <-chan string) {
	fmt.Println("接收通道信息", <-ch)
}

func recv(ch chan string) {
	fmt.Println("接收通道信息", <-ch)
}

func send1(ch chan string) { // 通道参数引用
	ch <- "shineyork" //
	fmt.Println("1  ")
	ch <- "sixstar" // 因为通道满了 =》阻塞
	fmt.Println("2  ")
	ch <- "go"
	fmt.Println("3  ")
}

func recv1(ch chan string) { // 通道参数引用
	s := <-ch
	fmt.Println(" s : ", s)
	fmt.Println("接收通道信息", <-ch) // 通道没有数据？取不出 =》 阻塞

	if <-ch == "go" {
		fmt.Println("go")
	}
}

func A() {
	fmt.Println("A")
}
func B() {
	fmt.Println("B")
}
func C() {
	fmt.Println("C")
}

/*
go func main(){
  go func() {
    fmt.Println("k")
  }()

  go func() {
    fmt.Println("ko")
  }()
}
*/

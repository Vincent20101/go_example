package main

import (
	"fmt"
	"testing"
	"time"
)

func timeAfter(done chan bool) {
	time.Sleep(2e9)
	done <- true
}

func TestTimer(t *testing.T) {
	ticker := time.Tick(1e9)
	for {
		<-ticker
		go curl()
	}
}
func curl() {
	fmt.Println("请求某一个接口")
}

func TestSelect(t *testing.T) {
	goodsCh := make(chan string) // 存储数据
	orderCh := make(chan string) // 存储数据

	done := make(chan bool)
	go goods(goodsCh)
	go order(orderCh) //

	go timeAfter(done)
	fmt.Println("my")

SHINEYORK:
	for {
		select { // 检查 =》 选择可以运行的通道进行执行 （非阻塞的）
		case order := <-orderCh:
			fmt.Println("order", order)
			// return // 结束当前的程序
			break SHINEYORK
		case goods := <-goodsCh:
			fmt.Println("goods", goods)
			// case <-done:
		case <-time.After(2e9):
			fmt.Println("超时")
			break SHINEYORK
			// default:
			// 	fmt.Println("default")
		}
	}
	fmt.Println("kk")

	// MySQL逻辑处理

	// goods逻辑处理
	// fmt.Println(data)
}

func goods(ch chan<- string) {
	time.Sleep(1e9)
	ch <- "goods"
}
func order(ch chan<- string) {
	time.Sleep(4e9)
	ch <- "order"
}
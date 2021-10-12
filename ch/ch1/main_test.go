package main

import (
	"fmt"
	"testing"
	"time"
)

func goods(ch chan<- string) {
	time.Sleep(1e9)
	ch <- "goods"
}
func order(ch chan<- string) {
	time.Sleep(3e9)
	ch <- "order"
}
func TestSelect(t *testing.T) {
	goodsCh := make(chan string) // 存储数据
	orderCh := make(chan string) // 存储数据
	go goods(goodsCh)
	go order(orderCh)

	ticker := time.After(2e9)
FCAK:
	for {
		select { // 检查 =》 选择可以运行的通道进行执行 （非阻塞的）
		case order := <-orderCh:
			fmt.Println("order", order)
			break FCAK
		case goods := <-goodsCh:
			fmt.Println("goods", goods)
		case <-ticker: // 产生新的chan =》 case每一次都是获取新的chan的信息
			fmt.Println("超时")
			break FCAK
			// default: // 当其他的分支阻塞的时候才执行
			// fmt.Println("default")
		}
	}

	// select { // switch
	// case order := <-orderCh:
	// 	fmt.Println("order", order)
	// case goods := <-goodsCh:
	// 	fmt.Println("goods", goods)
	// }

	// 	fmt.Println("kk")
}

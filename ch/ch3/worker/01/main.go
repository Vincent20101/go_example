package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan string)     // 存储数据
	done := make(chan []string) // 启动阻塞切换协程的执行
	data := make([]string, 0)
	wg.Add(1)
	go redis(ch) // defer wg.Done() // => 1

	go func(data []string, ch chan string, done chan []string) {
		for {
			d, ok := <-ch
			fmt.Println(ok)
			if  ok { // 只有在协程通道关闭的时候才能跳出循环
				data = append(data, d)
			} else {
				break
			}
		}
		fmt.Println("data:", data)
		done <- data
	}(data, ch, done)
	wg.Wait() // 在添加的协程个数，执行完之后就会取消阻塞
	close(ch) // 这是关闭ch通道

	fmt.Println("<-done", <-done) // 是不是 需要在 done 通道有信息的时候才执行

	// for _, ch := range chs {
	// 	data = append(data, <-ch)
	// }
	// for {
	// 	data = append(data, <-ch)
	// }

	// for v := range ch {
	// 	data = append(data, v)
	// }
	// fmt.Println(data)

	// for i := 0; i < 3; i++ { // 灵活扩展？？
	// 	data = append(data, <-ch)
	// }

}
func redis(ch chan<- string) {
	defer wg.Done()
	ch <- "redis"
}
func mysql(ch chan<- string) {
	defer wg.Done()
	ch <- "mysql"
}
func es(ch chan<- string) {
	defer wg.Done()
	ch <- "es"
}

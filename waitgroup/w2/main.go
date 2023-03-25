package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	for i := 0; i < 1; i++ {
		go test()
		time.Sleep(1 * time.Millisecond)
	}
	//模拟 永远不死...
	for {

	}
}

func test() {

	var ch chan int
	ch = make(chan int)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(index int) {
			defer func() {
				err := recover()
				if err != nil {
					fmt.Println(err)
				}
				wg.Done()
			}()
			ch <- index
			fmt.Println(fmt.Sprintf("我是第：%d个", index))
		}(i)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()

	for v := range ch {
		fmt.Println(v)
	}

}

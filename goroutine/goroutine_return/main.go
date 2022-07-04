package main

import (
	"context"
	"fmt"
	"time"
)

var (
	c chan interface{}
)

func main() {

	c = make(chan interface{}, 10)
	ctx, _ := context.WithCancel(context.Background())
	go func() {
		c <- test(ctx, c)
	}()
	t := time.Tick(2 * time.Second)
	//for {
	select {
	case <-c:
		//cancel()
		fmt.Println("c")
	case <-t:
		fmt.Println("timeout")
		//cancel()
		//return
		//time.Sleep(10 * time.Second)
		//default:
		//	fmt.Println("return")
		//	cancel()
		//	//time.Sleep(5 * time.Second)
		//	return
	}
	//}
	//time.Sleep(10 * time.Second)
	select {}
}

func test(ctx context.Context, c chan interface{}) (err error) {
	//time.Sleep(3 * time.Second)

	for i := 0; i < 3; i++ {
		go func(i int) {
			for {
				i++
				//c <- "sleep count"
				time.Sleep(2 * time.Second)
				fmt.Println("test", i)
				if i > 0 {
					break
				}
			}
		}(i)

	}
	//<-ctx.Done()
	return
}

package main

import (
	"fmt"
	"time"
)

func main() {

	//forrangetest()
	//selecttest()

	//ch := make(chan int)
	//done := make(chan struct{})
	//go selectblock(ch, done)
	//go func() {
	//	for {
	//		fmt.Println(<-ch)
	//	}
	//}()
	//
	////select {}
	//<-done

	go func() { time.Sleep(10 * time.Second) }()
	select {}
}

func selectblock(ch chan int, done chan struct{}) {
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		//close(ch)
	}()
	time.Sleep(10 * time.Second)
	done <- struct{}{}
}

func selecttest() {
	ch := make(chan *int)
	c2 := make(chan int)
	go func() {
		//time.Sleep(10e9)
		//for i := 0; i < 2; i++ {
		//	ch <- i
		//}
		//close(ch)
	}()
	go func() {
		//time.Sleep(20e9)
		c2 <- 10000
		//close(c2)
	}()

	for {
		select {
		case i := <-ch:
			fmt.Println(i)
			//return
		//default:
		//	fmt.Println("default")
		case k := <-c2:
			fmt.Println("k的值：", k)
		}
	}
	fmt.Println("done === done")
}

func forrangetest() {
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		for k := range ch {
			fmt.Println("num:", k)
		}
		done <- struct{}{}
	}()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
	<-done
}

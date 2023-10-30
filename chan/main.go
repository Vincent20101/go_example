package main

import (
	"fmt"
	"time"
)

func main() {

	//test1()
	stopCh := make(chan bool, 1)
	go test2(stopCh)
	time.Sleep(time.Second * 100)
	fmt.Println("done")
}

func test2(stopCh chan bool) {
	var i int

	for {
		select {
		case <-stopCh:
			fmt.Println("stop loop")
			goto LOOP
		default:
			//do something
			fmt.Println("doing something...")
			time.Sleep(time.Second)
			i++
			if i == 5 {
				stopCh <- true
			}
		}
	}
LOOP:
	fmt.Println("111111111")
	return
}

func test1() {
	c1 := make(chan interface{})
	c2 := make(chan interface{})

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- nil
		//close(c1)
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- nil
		//close(c2)
	}()

	for i := 0; i < 20; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received1", msg1)
			fmt.Println(c1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
			//default:
			//	fmt.Println("default")
		}
	}
}

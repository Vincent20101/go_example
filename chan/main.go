package main

import (
	"fmt"
	"time"
)

func main() {

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
		default:
			fmt.Println("default")
		}
	}
}

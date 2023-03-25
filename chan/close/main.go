package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan struct{})
	nq := make(chan int, 100)

	go func() {
		for i := 0; ; i++ {
			select {
			case s, ok := <-q:
				fmt.Println("leave", s, ok)
				return
			case n := <-nq:
				time.Sleep(1 * time.Second)
				fmt.Println("n=", n, ", i= ", i)
			}
		}
	}()
	for i := 0; i < 10; i++ {
		nq <- i
	}
	close(q)
	time.Sleep(3 * time.Second)
	fmt.Println("done ?")
}

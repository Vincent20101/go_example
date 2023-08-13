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
				fmt.Println("leave1", s, ok)
				return
			case n := <-nq:
				time.Sleep(1 * time.Second)
				fmt.Println("n1=", n, ", i1= ", i)
			}
		}
	}()
	go func() {
		for i := 0; ; i++ {
			select {
			case s, ok := <-q:
				fmt.Println("leave2", s, ok)
				return
			case n := <-nq:
				time.Sleep(1 * time.Second)
				fmt.Println("n2=", n, ", i2= ", i)
			}
		}
	}()
	for i := 0; i < 10; i++ {
		nq <- i
	}
	time.Sleep(30 * time.Second)
	close(q)
	fmt.Println("done ?")
}

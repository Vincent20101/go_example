package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	ch := make(chan int)
	done := make(chan bool)
	wg.Add(2)
	go send(1, 10, ch)
	go recv(ch, done)
	<-done
	wg.Wait()
}

func send(start, end int, ch chan<- int) {
	for i := start; i < end; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func recv(in <-chan int, done chan<- bool) {
	for k := range in {
		fmt.Println(k)
	}
	done <- true
	wg.Done()
}

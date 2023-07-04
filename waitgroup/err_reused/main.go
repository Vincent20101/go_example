package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		fmt.Println("Goroutine 1")
		wg.Done()
	}()

	wg.Wait()

	wg.Add(1) // 这里会导致报错

	go func() {
		fmt.Println("Goroutine 2")
		wg.Done()
	}()

	wg.Wait()
}

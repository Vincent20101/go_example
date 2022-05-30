package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan *struct{})
	wg.Add(1)
	go func() {
		ch <- nil
		wg.Done()
	}()
	fmt.Println(<-ch)
	wg.Wait()
	fmt.Println("end")
	ctx := context.TODO()
	select {
	case <-ctx.Done():
		fmt.Println("end======")
	default:


	}
}

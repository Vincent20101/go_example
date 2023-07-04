package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	ch := make(chan *struct{})
	//wg.Add(1)
	go func() {
		time.Sleep(time.Second * 3)
		//ch <- nil
		close(ch)
		//wg.Done()
	}()
	//fmt.Println(<-ch)
	//wg.Wait()
	fmt.Println("end")
	ctx := context.TODO()
	select {
	case <-ctx.Done():
		fmt.Println("end======")
	case c, ok := <-ch:
		fmt.Println(c, ok)
		//default:
		//	fmt.Println("test")

	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	var done = make(chan struct{})
	t := time.NewTicker(time.Second * 2)
	//go func() {
	//	time.Sleep(time.Second * 10)
	//	close(done)
	//}()
	var i = 100
	for {
		fmt.Println("1111")
		select {
		//case <-t.C:
		//	fmt.Println("continue")
		//	continue
		//	fmt.Println("continue222")
		case <-t.C:
			switch i {
			case 100:
				fmt.Println("lhb 100")
				continue
				i++
			case 101:
				fmt.Println("lhb 101")
				continue
			}
		case <-done:
			fmt.Println("done")
		}
		fmt.Println("lhb")
		time.Sleep(time.Second * 3)

	}
}

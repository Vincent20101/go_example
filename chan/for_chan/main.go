package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Unix(0, 0))
	var (
		t       = time.NewTicker(time.Second * 2)
		startCh = make(chan struct{}, 1)
		endCh   = make(chan struct{}, 1)
		a       int
	)
	for {
		select {
		case <-t.C:
			startCh <- struct{}{}
		case <-startCh:
			fmt.Println("lhb: starting")
			a++
			if a == 2 {
				endCh <- struct{}{}
			}
		case <-endCh:
			fmt.Println("lhb:ending")
			break
		}
	}
	fmt.Println("exit")
	time.Sleep(time.Second * 1000)
}

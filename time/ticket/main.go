package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	//defer ticker.Stop()
	done := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		done <- true
	}()
	go func() {
		for range ticker.C {
			fmt.Println("==============")
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		select {
		case <-done:
			fmt.Println("Done!")
			//return
		case t := <-ticker.C:
			fmt.Println("Current time: ", t)
		}
	}

}

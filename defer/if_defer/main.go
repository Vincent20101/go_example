package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println("start time: ", startTime)
	defer fmt.Println("duration:", time.Now().Sub(startTime))
	if true {
		defer func() {
			fmt.Println("duration upgraded:", time.Now().Sub(startTime))
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("finish time: ", time.Now())
}

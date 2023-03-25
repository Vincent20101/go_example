package main

import (
	"fmt"
	"time"
)

func main() {
	deferGuess()
	fmt.Println(deferReturn())
}

func deferGuess() {
	startTime := time.Now()
	fmt.Println("start time: ", startTime)
	defer fmt.Println("duration:", time.Now().Sub(startTime))
	defer func() {
		fmt.Println("duration upgraded:", time.Now().Sub(startTime))
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("finish time: ", time.Now())
}

func deferReturn() (ret int) {
	defer func() {
		ret++
	}()
	return 10
}

package main

import (
	"fmt"
	"time"
)

func main() {
	time.AfterFunc(0*time.Millisecond, func() {
		fmt.Println("test")
	})
	time.Sleep(1e9 * time.Second)
}

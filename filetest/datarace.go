package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("i的值为：", i)
		}(i)
	}
	time.Sleep(2 * time.Second)
}

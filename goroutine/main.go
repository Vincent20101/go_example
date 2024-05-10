package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	fmt.Println(runtime.NumCPU())
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

func tg() {
	for i := 0; i < 10; i++ {
		go func(i int) {
			go func(ii int) {
				go func(iii int) {
					time.Sleep(5 * time.Second)
					var a map[int]int
					a[iii] = iii
				}(ii)
			}(i)
		}(i)
	}
}

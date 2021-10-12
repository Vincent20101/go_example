package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	// 1.14之前加上该行打印从0开始，之后版本则不影响
	//time.Sleep(time.Second)
	wg.Wait()
}

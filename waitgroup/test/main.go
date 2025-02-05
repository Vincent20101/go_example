package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	var m2 sync.Map
	var swg sync.WaitGroup
	for i := 0; i < 20; i++ {
		if i == 2 {
			time.Sleep(2 * time.Second)
		}
		swg.Add(1)
		go func(i int) {
			key := strconv.Itoa(i)
			m2.Store(key, i)
			value, ok := m2.Load(key)
			if !ok {
				fmt.Println("There are something error")
			}
			fmt.Printf("key:%s value:%d\n", key, value)
			swg.Done()
		}(i)
	}
	swg.Wait()

	//var wg sync.WaitGroup
	//wg.Add(1)
	//go func() {
	//	time.Sleep(time.Millisecond)
	//	wg.Done()
	//	wg.Add(1)
	//}()
	//wg.Wait()
}

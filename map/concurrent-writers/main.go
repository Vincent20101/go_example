package main

import (
	"fmt"
	"sync"
)

func main() {
	var m = make(map[string]int)
	m = map[string]int{"one": 1, "two": 2, "three": 3}
	var wg sync.WaitGroup

	// Goroutine 1
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//m = make(map[string]int) // 覆盖整个 map
			//m = map[string]int{
			//	"a": 1,
			//}
			m["a"] = 1
			wg.Done()
		}()
	}

	// Goroutine 2
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//m = make(map[string]int) // 覆盖整个 map
			m = map[string]int{
				"a1": 11,
			}
			//m["a"] = 2
			wg.Done()
		}()
	}
	wg.Wait()
	//time.Sleep(1 * time.Hour)
	fmt.Println(m)
}

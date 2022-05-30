package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

func main() {
	var buffer unsafe.Pointer
	var wg sync.WaitGroup

	var writeFn = func(index int) {
		b := make([]int, 0)
		b = append(b, index)
		b = append(b, index)
		b = append(b, index)
		fmt.Println("---------",b)

		atomic.StorePointer(&buffer, unsafe.Pointer(&b))
	}

	var readFn = func() {
		b := atomic.LoadPointer(&buffer)
		data := *(*[]int)(b)

		fmt.Println(b, data)
	}

	// 初始化
	writeFn(0)

	// 写入
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			writeFn(index)
			time.Sleep(time.Millisecond * 100)

			wg.Done()
		}(i)
	}

	//time.Sleep(5*time.Second)

	// 读取
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			readFn()
			time.Sleep(time.Millisecond * 100)

			wg.Done()
		}()
	}

	wg.Wait()
}

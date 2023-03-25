package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// 原子替换内存地址
func TestBuffer(t *testing.T) {
	var buffer unsafe.Pointer
	var wg sync.WaitGroup

	var writeFn = func(index int) {
		b := make([]int, 0)
		b = append(b, index)
		b = append(b, index)
		b = append(b, index)

		atomic.StorePointer(&buffer, unsafe.Pointer(&b))
	}

	var readFn = func() {
		b := atomic.LoadPointer(&buffer)
		data := *(*[]int)(b)

		t.Log(b, data)
	}

	// 初始化
	writeFn(0)

	// 写入
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			writeFn(i)
			time.Sleep(time.Millisecond * 100)

			wg.Done()
		}(i)
	}

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

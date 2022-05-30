package main

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

// string转ytes
func Str2sbyte(s string) (b []byte) {
	*(*string)(unsafe.Pointer(&b)) = s	// 把s的地址付给b
	*(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&b)) + 2*unsafe.Sizeof(&b))) = len(s)	// 修改容量为长度
	return
}

// []byte转string
func Sbyte2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BenchmarkByteString2(b *testing.B) {
	str := "this is a string"
	var bs []byte

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bs = Str2sbyte(str)
		str = Sbyte2str(bs)
		b.Log(bs,str)
	}

	b.StopTimer()
}

// 原子替换内存地址
func TestBuffer(t *testing.T) {
	var buffer unsafe.Pointer
	var wg sync.WaitGroup

	var writeFn = func(index int) {
		b := make([]int, 0)
		b = append(b, index)
		b = append(b, index)
		b = append(b, index)
		t.Log("---------",b)

		atomic.StorePointer(&buffer, unsafe.Pointer(&b))
	}

	var readFn = func() {
		b := atomic.LoadPointer(&buffer)
		data := *(*[]int)(b)

		t.Log(b, data)
	}

	// 初始化
	writeFn(0)
	writeFn(1)
	writeFn(2)
	writeFn(3)
	writeFn(4)
	writeFn(5)
	writeFn(6)

	// 写入
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(index int) {
			t.Log("i的值为：",i)
			writeFn(i)
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
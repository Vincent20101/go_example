package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex
var holder map[string]bool = make(map[string]bool)

// 临时对象池
var p = sync.Pool{
	New: func() interface{} {
		fmt.Println("new")
		buffer := make([]byte, 256)
		return &buffer
	},
}

// wg 是一个指针类型,必须是一个内存地址
func readContent(wg *sync.WaitGroup) {
	defer wg.Done()
	client := &http.Client{
		Timeout: time.Second * 30, // 设置30秒超时
	}
	resp, err := client.Get("http://my.oschina.net/xinxingegeya/home")
	if err != nil {
		// handle error
		return
		//panic(err)
	}

	defer resp.Body.Close()

	byteSlice := p.Get().(*[]byte) //类型断言

	key := fmt.Sprintf("%p", byteSlice)
	//fmt.Println("lhb2:", byteSlice)
	//fmt.Println("lhb2:", key)

	// 互斥锁,实现同步操作
	mu.Lock()
	_, ok := holder[key]
	if !ok {
		holder[key] = true
	}
	mu.Unlock()

	//numBytesReadAtLeast, err := io.ReadFull(resp.Body, *byteSlice)
	_, err = io.ReadFull(resp.Body, *byteSlice)
	if err != nil {
		// handle error
		fmt.Println("err:", err.Error())
		return
	}
	//fmt.Println("lhb:", byteSlice)

	p.Put(byteSlice)

	//log.Printf("Number of bytes read: %d\n", numBytesReadAtLeast)
	//fmt.Println(string((*byteSlice)[:256]))
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go readContent(&wg)
	}

	wg.Wait()

	fmt.Println(len(holder))

	for key, val := range holder {
		fmt.Println("Key:", key, "Value:", val)
	}

	fmt.Println("end...")
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// 模拟并发调用实例对象
	for i := 0; i < 100; i++ {
		go func() {
			GetInstance()
		}()
	}
	time.Sleep(time.Second)
}

type Instance struct {
}

// 实例全局对象
var lazyInstance *Instance
var lock = &sync.Mutex{}

func GetInstance() *Instance {
	// 判断第一次如果为空，则给lazy对象重新赋值。
	//Double-Checked Locking
	if lazyInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if lazyInstance == nil {
			fmt.Println("创建实例")
			lazyInstance = &Instance{}
		}
	}

	return lazyInstance
}

package main

import (
	"fmt"
	"sync"
	"time"
)

type Instance struct{}

var lazyInstance *Instance

var once = &sync.Once{}

func GetInstance() *Instance {
	if lazyInstance == nil {
		once.Do(func() {
			fmt.Println("创建实例")
			lazyInstance = &Instance{}
		})
	}
	return lazyInstance
}

func main() {
	// 模拟并发调用实例对象
	for i := 0; i < 100; i++ {
		go func() {
			GetInstance()
		}()
	}
	time.Sleep(time.Second)
}

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Instance struct {
}

// 实例全局对象
var lazyInstance *Instance
var lock = &sync.Mutex{}
var initialized uint32

func GetInstance() *Instance {
	// 使用原子操作检查是否已经初始化
	if atomic.LoadUint32(&initialized) == 1 {
		return lazyInstance
	}
	return doSlow(&initialized, lock)
}

func doSlow(initialized *uint32, l *sync.Mutex) *Instance {
	l.Lock()
	defer l.Unlock()

	// 再次检查，因为在获取锁之前可能已经被其他goroutine初始化
	if *initialized == 0 {
		fmt.Println("创建实例")
		lazyInstance = &Instance{}
		atomic.StoreUint32(initialized, 1)
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

}

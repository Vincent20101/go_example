package main

func main() {
	// 调用实例对象
	GetInstance()
}

type Instance struct {
}

// 实例全局对象
var lazyInstance *Instance

func GetInstance() *Instance {
	// 判断第一次如果为空，则给lazy对象重新赋值。
	if lazyInstance == nil {
		lazyInstance = &Instance{}
	}
	return lazyInstance
}

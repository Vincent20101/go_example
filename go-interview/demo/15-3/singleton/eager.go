package main

func main() {
	GetEagerInstance()
}

// 对象
type EagerInstance struct {
}

var eagerInstance *EagerInstance

// 饿汉式：在程序初始化的时候或者类加载的时候就已经创建好对象，加载速度快。
func init() {
	eagerInstance = &EagerInstance{}
}

// 饿汉式调用
func GetEagerInstance() *EagerInstance {
	return eagerInstance
}

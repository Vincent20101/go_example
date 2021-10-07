package main

import "time"

// go vet可以避免一些坑，比如：锁拷贝、闭包引用循环变量
// go vet ./filetest/vet.go
func main() {
	a := map[int]int{1: 1, 2: 3}
	b := map[int]*int{}
	for k, r := range a {
		go func() {
			b[k] = &r
		}()
	}
	time.Sleep(3 * time.Second)
}

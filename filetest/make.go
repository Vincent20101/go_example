package main

import (
	"fmt"
	"time"
)

// go build make.go && go tool objdump ./make | grep -E "make.go:6|make.go:10|make.go:14"
func main1() {
	// make slice
	// 空间开得比较大，是为了让这个 slice 分配在堆上，栈上的 slice 结果不太一样
	var sl = make([]int,100000)
	println(sl)

	// make channel
	var ch = make(chan int, 5)
	println(ch)

	// make map
	var m = make(map[int]int,22)
	println(m)
}

func main(){
	sl := make([]int,10)
	ssl := sl[10:]
	fmt.Println(ssl)
	println(ssl)
	
	fmt.Println(time.Now())
}
package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func main() {
	//分析过程：定义x = 5，然后将返回值int赋值为x也就是5，此时defer函数执行，x变为6，但此时int值并不受影响（可以理解为两块内存地址），再执行RET所以返回值为5
	fmt.Println(f1()) // 5
	//分析过程：调用f2函数，将y返回值x 赋值为5，此时y执行defer函数将x 变为6，同时返回值变为6（可以理解为同一块内存地址），再执行RET所以返回值为6
	fmt.Println(f2()) // 6
	//分析过程：定义x = 5，然后将返回值y赋值为x也就是5，此时defer函数执行，x变为6，但此时y值并不受影响（可以理解为两块内存地址），再执行RET所以返回值为5
	fmt.Println(f3()) // 5
	//分析过程：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值，因此在执行到defer时，x = 0。执行到return时，将 返回值x赋值为5，此时defer函数中x 变为1，但对返回值x没有影响（可以理解为两块内存地址）， 再执行RET所以返回值为5
	fmt.Println(f4()) // 5
}

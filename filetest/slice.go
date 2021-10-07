package main

import "fmt"

type MySlice []int

func (m MySlice) append(i int) {
	m = append(m, i)
	fmt.Println(m) // [0 1]

}

// receiver 是值不是指针时
func main() {
	var m MySlice = make(MySlice, 1, 1)
	fmt.Println(m) // [0]
	m.append(1)
	fmt.Println(m) // [0]
}

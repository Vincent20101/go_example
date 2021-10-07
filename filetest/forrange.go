package main

import "fmt"

func main() {
	var m = map[int]person{
		1: {
			11, "abc",
		},
		2: {
			22, "dev",
		},
	}

	var mm = map[int]*person{}
	// for range 生成的k v始终是同一个对象
	for k, v := range m {
		mm[k] = &v
	}

	fmt.Println(mm[1], mm[2])
}

type person struct {
	age  int
	name string
}

package main

import "fmt"

func main() {
	m := make(map[string]string, 10)

	m["s"] = "sdf"
	fmt.Println(m["s"])
	m["s"] = "abc"
	fmt.Println(m)
	ab, ok := m["a"]
	fmt.Println(&ab, ok)

	var a *int
	fmt.Printf("地址是：%p\n", a)

	var strSlice []string
	var intSlice []int
	var emptySlice = []int{}                                   //分配到了内存，但无元素
	fmt.Printf("%p, %p, %p\n", strSlice, intSlice, emptySlice) //0x0, 0x0, 0x1195ab8

	c := make(chan struct{}, 1)
	d := make(chan struct{}, 1)
	c <- struct{}{}
	d <- struct{}{}

	anies := <-c
	fmt.Println(anies)
	a2 := <-d
	fmt.Println(&a2)
	fmt.Println(&c)
	fmt.Println(&d)
}

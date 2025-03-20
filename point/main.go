package main

import (
	"fmt"
	"strings"
)

type student struct {
	name       string
	age        int
	level      int
	work       *string
	secondWork []*string
}

func main() {
	var point = new(int)
	*point = 10

	pp := point

	var point2 = new(int)
	*point2 = 100

	point = point2

	fmt.Println(*pp)
	fmt.Println(*point)
	fmt.Println(*point2)

	fmt.Println(128 & 128)
	var a *int
	b := 1
	a = &b

	c := new(int)
	//c = a
	*c = *a
	fmt.Println(c, a)
	fmt.Println(*c, *a)
	*c = 18
	fmt.Println(*a, b, *c)

	s1 := &student{
		name:  "s1",
		age:   10,
		level: 1,
	}

	s2 := new(student)
	fmt.Println("s2", s2.work, s2.secondWork)
	fmt.Println("s2", s2.work == nil, s2.secondWork == nil)

	s := "s:d:1"
	fmt.Println(strings.Split(s, ":")[0])
	s2 = s1
	//*s2 = *s1
	s2.age = 18
	fmt.Println(s1, s2)

	newA := &[]uint{3}[0]
	fmt.Println(*newA)

	var aa, bb int
	aa, bb = 3, 4
	swapTwo(&aa, &bb)
	fmt.Println(aa, bb)

	var cc = &[]int{0}[0]
	*cc = 2
	fmt.Println(*cc)
}

// 使用第三变量交换a,b值:go 直接交换值
func swapTwo(a *int, b *int) {
	//*a, *b = *b, *a
	a, b = b, a
}

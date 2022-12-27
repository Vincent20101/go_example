package main

import "fmt"

type student struct {
	name  string
	age   int
	level int
}

func main() {
	var a *int
	b := 1
	a = &b

	c := new(int)
	//c = a
	*c = *a
	fmt.Println(c, a)
	*c = 18
	fmt.Println(*a, b, *c)

	s1 := &student{
		name:  "s1",
		age:   10,
		level: 1,
	}

	s2 := new(student)
	s2 = s1
	//*s2 = *s1
	s2.age = 18
	fmt.Println(s1, s2)
}

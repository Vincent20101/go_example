package main

import "fmt"

type student struct {
	age      int
	name     *string
	homework string
}

type person struct {
	s    *student
	name string
}

func main() {
	p := &person{
		s: &student{
			age:      10,
			name:     nil,
			homework: "hw",
		},
		name: "new test",
	}

	p.s.age = 22
	change(p)

	fmt.Printf("%+v", *p.s)
}

func change(p *person) {
	//s := "next change"
	p.s.name = &p.name

}

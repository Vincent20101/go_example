package main

import "fmt"

type People struct {
	name string
	age  int
}

func (p *People) GetName() string {
	return p.name
}

func (p People) Test() {

}

func main() {
	var tt []*int
	fmt.Println(len(tt))
}

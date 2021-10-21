package main

import "fmt"

type human interface {
	speak()
	sing()
}

type man struct {
}

func (m man) speak() {
	fmt.Println("speaking")
}

func (m *man) sing() {
	fmt.Println("singing")
}

func main() {
	var h human = &man{}
	//var h human = man{} // 调用sing()会报错
	h.speak()
	h.sing()
}

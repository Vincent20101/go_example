package main

import "fmt"

func main() {
	var x interface{} = nil
	var y *int = nil
	var s struct{}
	fmt.Println(x == nil)
	fmt.Println(y == nil)
	fmt.Println(s)
	interfaceIsNil(x)
	interfaceIsNil(y)
	interfaceIsNil(s)
}

func interfaceIsNil(x interface{}) {
	fmt.Printf("%#v\n", x)
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

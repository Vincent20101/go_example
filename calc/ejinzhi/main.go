package main

import "fmt"

func main() {
	a := 12
	b := 4
	c := a &^ b
	fmt.Printf("a: %08b\n", a)
	fmt.Printf("b：%08b\n", b)
	fmt.Printf("c: %#v 二进制：%08b\n", c, c)
	/*输出：
	a: 00001100
	b：00000100
	c: 8 二进制：00001000
	*/
}

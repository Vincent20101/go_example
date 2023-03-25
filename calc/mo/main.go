package main

import "fmt"

func main() {
	var a = 10
	a = a & ^(1 << 3)
	fmt.Println(a)
	fmt.Println(^a) //取反
	fmt.Println("============")
	b := 13634534
	fmt.Println(b % 5)
	fmt.Println(b & 0x03)
	fmt.Println(b & 3)
}

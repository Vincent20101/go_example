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
	fmt.Println("============")
	fmt.Println(0 ^ 0)
	fmt.Println(0 ^ 1)
	fmt.Println(1 ^ 0)
	fmt.Println(1 ^ 1)
}

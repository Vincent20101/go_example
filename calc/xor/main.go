package main

import "fmt"

func main() {
	var a = 300
	var b = 500
	fmt.Println(fmt.Sprintf("%.16b", a))
	fmt.Println(fmt.Sprintf("%.16b", b))
	fmt.Println(fmt.Sprintf("%.16b", ^a&b))
	fmt.Println(fmt.Sprintf("%.16b", b&^a))
	fmt.Println((b ^ (a & b)) == b&^a)
	fmt.Println(fmt.Sprintf("%.16b", a^a))
	fmt.Println(fmt.Sprintf("%.16b", a|a))
}

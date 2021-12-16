package main

import "fmt"

func modify01(s []int) {
	s[0] = 90
	s = append(s, 8888)
	fmt.Println(s)
}
func main() {
	a := [3]int{89, 90, 91}
	modify01(a[:])
	fmt.Println(a)
}

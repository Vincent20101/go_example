package main

import "fmt"

func main() {
	var s []string

	b := []string{"a", "b"}

	s = append(s, "d", "e")
	c := append(s, b...)

	fmt.Println(s)
	fmt.Println(b)
	fmt.Println(c)
}

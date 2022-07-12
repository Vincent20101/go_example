package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 58, 9, 0, 6, 4, 4, 2}

	for _, v := range s {
		fmt.Println(v)
	}

	var s1 []string
	fmt.Println(s1[0])
}

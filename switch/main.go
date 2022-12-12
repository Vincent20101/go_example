package main

import "fmt"

func main() {
	var a = 1
	switch a {
	case 1:
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
}

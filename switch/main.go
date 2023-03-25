package main

import "fmt"

func main() {
	through()
	//swith()
	delete(nil, 1)
}

func through() {
	var a = 1
	switch a {
	case 1:
		fmt.Println("sdfsdf")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
}

func swith() {
	var a = 1
	switch a {
	case 1:
		fmt.Println("coming...")
		break
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	}
}

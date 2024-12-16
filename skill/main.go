package main

import "fmt"

func main() {
	number := 10000000
	better := 10_000_000

	fmt.Println(number == better)

	name := "Bob"
	fmt.Printf("My name is %[1]s. Yes, you heard that right: %[1]s\n", name)
}

package main

import "fmt"

func main() {
	var name = new(string)
	stringPointer(name)

	fmt.Println(*name)
}

func stringPointer(name *string) {
	*name = "test"
}

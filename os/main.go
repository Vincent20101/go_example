package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Create("./a.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
}

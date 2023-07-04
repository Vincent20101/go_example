package main

import (
	"fmt"
	"io"
)

func main() {
	fmt.Println(io.ReadAll(nil))
}

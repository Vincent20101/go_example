package main

import (
	"fmt"

	luhn "github.com/vincent20101/go-example/luhn/tool"
)

func main() {
	n, _ := luhn.Generate("49015420323751")
	fmt.Println(n)
}

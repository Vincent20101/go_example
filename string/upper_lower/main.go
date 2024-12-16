package main

import (
	"fmt"
	"strings"
)

func isAllUpper(s string) bool {
	return s == strings.ToUpper(s)
}

func isAllLower(s string) bool {
	return s == strings.ToLower(s)
}

func main() {
	str1 := "HELLO WORLD - 123 - 应用@!#$"
	str2 := "hello world - 123 - 应用@!#$"
	str3 := "Hello World - 123 - 应用@!#$"

	fmt.Printf("Is '%s' all uppercase? %v\n", str1, isAllUpper(str1)) // true
	fmt.Printf("Is '%s' all uppercase? %v\n", str2, isAllUpper(str2)) // false
	fmt.Printf("Is '%s' all lowercase? %v\n", str2, isAllLower(str2)) // true
	fmt.Printf("Is '%s' all lowercase? %v\n", str3, isAllLower(str3)) // false
}

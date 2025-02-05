package main

import (
	"fmt"
)

func truncateStringByRune(s string, maxBytes int) string {
	a := []rune(s)
	if maxBytes > len(a) {
		return ""
	}
	return string(a[:maxBytes])
}

func main() {
	s1 := "人生ABC"
	fmt.Println("Truncated string:", truncateStringByRune(s1, 4))

	s2 := "人AB们DEF"
	fmt.Println("Truncated string:", truncateStringByRune(s2, 6))
}

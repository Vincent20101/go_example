package main

import (
	"fmt"
	"unicode/utf8"
)

func truncateStringByBytes(s string, maxBytes int) string {
	var size int
	var result []rune

	for _, runeValue := range s {
		byteSize := utf8.RuneLen(runeValue)
		if size+byteSize > maxBytes {
			break
		}
		size += byteSize
		result = append(result, runeValue)
	}

	return string(result)
}

func main() {
	s1 := "人生ABC"
	fmt.Println("Truncated string:", truncateStringByBytes(s1, 4))

	s2 := "人AB们DEF"
	fmt.Println("Truncated string:", truncateStringByBytes(s2, 6))
}

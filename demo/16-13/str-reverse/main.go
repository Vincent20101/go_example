package main

import "log"

func reverseWithTmp(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
	}
	return string(s)

}

func reverseWithXOR(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i] = s[i] ^ s[j]
		s[j] = s[i] ^ s[j]
		s[i] = s[i] ^ s[j]
	}
	return string(s)
}

func main() {
	a := 6
	b := 5
	log.Println(a ^ a)
	log.Println(a ^ 0)
	log.Println(a ^ b)
	log.Println(b ^ a)
	return
	str := "hello world"
	reversedStr := reverseWithTmp(str)
	println(reversedStr)

	reversedStr = reverseWithXOR(str)
	println(reversedStr)

}

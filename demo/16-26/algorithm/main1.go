package main

import "fmt"

func permute(str []rune, l, r int) {
	if l == r {
		fmt.Println(string(str))
	} else {
		for i := l; i <= r; i++ {
			str[l], str[i] = str[i], str[l]
			permute(str, l+1, r)
			str[l], str[i] = str[i], str[l]
		}
	}
}

func main() {
	str := "abcde"
	n := len(str)
	permute([]rune(str), 0, n-1)
}

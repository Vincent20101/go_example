package luhn

import "unicode/utf8"

func calculateDigit(digit int) int {
	digit = digit * 2

	if digit > 9 {
		digit = digit - 9
	}

	return digit
}

func reverse(s string) string {
	cs := make([]rune, utf8.RuneCountInString(s))
	i := len(cs)
	for _, c := range s {
		i--
		cs[i] = c
	}
	return string(cs)
}

func even(number int) bool {
	return number%2 == 0
}

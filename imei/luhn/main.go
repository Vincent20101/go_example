package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	//bool, err :=Verify("49015420323751")
	//fmt.Println(bool,err)
	//nl, err :=Generate("35576205279323")
	nl, err :=Generate("26053179311383")
	fmt.Println(nl,err)
}

// Verify verifies if the last digit is a valid luhn check digit
func Verify(n string) (v bool, err error) {
	// Whole string except last digit
	cn := n[:len(n)-1]
	rcn, err := Generate(cn)
	if err != nil {
		return
	}

	v = n == rcn
	return
}

// Generate generates a luhn check digit
func Generate(n string) (nl string, err error) {
	// Reverse string
	rn := reverse(n)
	sum := 0

	// Range over string
	for i, r := range rn {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			err = fmt.Errorf("cannot convert to digit: %s at index %d", string(r), i)
			return "", err
		}

		// When the check digit does not exist: take every uneven digit and double it
		if even(i) {
			digit = calculateDigit(digit)
		}
		sum = sum + digit
	}

	fmt.Println(strconv.FormatInt(int64((sum*9)%10), 10))
	nl = n + strconv.FormatInt(int64((sum*9)%10), 10)
	return
}


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
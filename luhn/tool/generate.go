package luhn

import (
	"fmt"
	"strconv"
)

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

	nl = n + strconv.FormatInt(int64((sum*9)%10), 10)
	return
}

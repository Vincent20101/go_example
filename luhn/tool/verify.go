package luhn

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

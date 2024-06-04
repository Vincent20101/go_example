package main

func reverseWithTmp(str string) string {
	s := []rune(str)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		tmp := s[i]
		s[i] = s[j]
		s[j] = tmp
	}
	return string(s)

}

// a=a^b
// b=a^b  //b=a^b=(a^b)^b=a^(b^b)=a^0=a
// a=a^b  //a=a^b=(a^b)^a=b^(a^a)=b^0=b
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
	//a := 6
	//b := 5
	//log.Println(a ^ a)
	//log.Println(a ^ 0)
	//log.Println(a ^ b)
	//log.Println(b ^ a)
	//return
	str := "hello world"
	reversedStr := reverseWithTmp(str)
	println(reversedStr)

	reversedStr = reverseWithXOR(str)
	println(reversedStr)

}

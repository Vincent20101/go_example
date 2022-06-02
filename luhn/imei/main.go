package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(calculateLuhnCheckDigit("49015420323751"))

	fmt.Println(len("49015420323751"))

	str := "49015420323751"
	var strPointer *string
	strPointer = &str
	*strPointer = strings.Join([]string{*strPointer,"0"},"")
	fmt.Println(*strPointer)


	//s := strings.HasPrefix("imei-311570000010010","imei-")
	//fmt.Println(s)
	strSlice := strings.Split("imei-311570000010010", "-")
	fmt.Println(strSlice[0],strSlice[1])

	//s1 := strings.Replace("imei-311570000010010","imei-","imeisv-",-1)
	//fmt.Println(s1)
	//var strPointer2 = new(string)
	//*strPointer2 = "49015420323751"
	//var byteBuffer bytes.Buffer
	//byteBuffer.WriteString(*strPointer2)
	//byteBuffer.WriteString("0")
	//*strPointer2 = byteBuffer.String()
	//fmt.Println(*strPointer2)
}
func calculateLuhnCheckDigit(imei string) (int) {
	num, _ := strconv.ParseInt(imei, 10, 64)
	var sum	int = 0
	var tmp int = 0

	// calculate the luhn check digit
	var checkDigit int = 0
	for i := 0; i < len(imei); i++ {
		tmp = int((num % int64(math.Pow10(i+1))) / int64(math.Pow10(i)))
		if i%2 == 0 {
			sum += ((tmp * 2) / 10 + (tmp * 2) % 10)
		}
	}
	if sum % 10 == 0 {
		checkDigit = 0
	} else {
		checkDigit = 10 - sum % 10
	}

	return  checkDigit
}
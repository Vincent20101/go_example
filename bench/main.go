package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {

}

func ConcatStr1(str1, str2 string) string {
	return str1 + str2
}

func ConcatStr2(str1, str2 string) string {
	return fmt.Sprintf("%s%s", str1, str2)
}

func ConcatStr3(str1, str2 string) string {
	return strings.Join([]string{str1, str2}, "")
}

func ConcatStr4(str1, str2 string) string {
	b := bytes.Buffer{}
	b.WriteString(str1)
	b.WriteString(str2)
	return b.String()
}

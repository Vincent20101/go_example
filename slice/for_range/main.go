package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4, 58, 9, 0, 6, 4, 4, 2}

	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Println("=========")
	var s1 []string
	//fmt.Println(s1[0])
	s1 = nil
	s2 := &s1
	fmt.Println(*s2)
	for _, v := range s1 {
		fmt.Println(v)
	}
	if s1 == nil {
		if true {
			goto cleanUp
		}
	}
	fmt.Println("clean up done")
cleanUp:
	fmt.Println("clean up ing...")
	fmt.Println("I am clean up xxxx")
	//return

	var s3 *string
	fmt.Println(s3)
	s4 := ""
	s3 = &s4
	fmt.Println(s3)
	fmt.Println("==", *s3)
	fmt.Println(s1, s2, s3, s4)
	if s4 == "xx" || (s1 == nil && s4 != "") || (s2 != nil && (s1 == nil || s3 != &s4)) {
		fmt.Println("true")
	}
}

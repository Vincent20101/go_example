package main

import (
	"fmt"
	"unicode/utf8"
)

/**
 * 反转字符串
 * @param str string字符串
 * @return string字符串
 */
func solve(str string) string {
	// write code here
	if len(str) == 0 {
		return ""
	}
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverse(str string) string {
	rs := []rune(str)
	len := len(rs)
	var tt []rune

	tt = make([]rune, 0)
	for i := 0; i < len; i++ {
		tt = append(tt, rs[len-i-1])
	}
	return string(tt[0:])
}
func reverse2(s string) string {
	ss := make([]rune, utf8.RuneCountInString(s))
	i := len(ss)
	for _, c := range s {
		i--
		ss[i] = c
	}
	return string(ss)
}

func testReverseStringV2() {
	var str = "hello中文 "
	var r []rune = []rune(str)

	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}

	str = string(r)
	fmt.Println(str)
}

func testHuiWen() {
	var str = "上海自来水来自海上"
	var r []rune = []rune(str)

	for i := 0; i < len(r)/2; i++ {
		tmp := r[len(r)-i-1]
		r[len(r)-i-1] = r[i]
		r[i] = tmp
	}

	str2 := string(r)
	if str2 == str {
		fmt.Println(str, " is huiwen")
	} else {
		fmt.Println(str, " is not huiwen")
	}
}

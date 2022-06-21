package main

import "unicode/utf8"

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

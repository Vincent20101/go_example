package main

import "log"

// ComputeLPSArray 计算部分匹配表
func ComputeLPSArray(p string) []int {
	lps := make([]int, len(p))
	length := 0 // 长度 of the previous longest prefix suffix
	i := 1

	// lps[0] is always 0
	lps[0] = 0

	// loop calculates lps[i] for i from 1 to len(p)-1
	for i < len(p) {
		if p[i] == p[length] {
			length++
			lps[i] = length
			i++
		} else { // (p[i] != p[length])
			if length != 0 {
				length = lps[length-1]
				// We do not increment i here
			} else { // if (length == 0)
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// KMPSearch executes the KMP search algorithm
func KMPSearch(s, p string) int {
	lps := ComputeLPSArray(p)
	i := 0 // index for s
	j := 0 // index for p

	for i < len(s) {
		if p[j] == s[i] {
			i++
			j++
		}

		if j == len(p) {
			return i - j // Found pattern at index i-j
		} else if i < len(s) && p[j] != s[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i = i + 1
			}
		}
	}
	return -1 // Pattern not found
}

func main() {
	s := "ABABDABACDABABCABAB"
	p := "ABABCABAB"
	result := KMPSearch(s, p)
	if result != -1 {
		log.Printf("Found pattern at index %d\n", result)
	} else {
		log.Println("Pattern not found")
	}
}

package main

import (
	"fmt"
	"sort"
)

// 给定一个字符串数组，找出数组中最长的字符串，使其能由数组中其他的字符串组成
func canFormByCombining(str string, wordSet map[string]bool, memo map[string]bool) bool {
	if val, found := memo[str]; found {
		return val
	}
	for i := 1; i < len(str); i++ {
		prefix, suffix := str[:i], str[i:]
		if wordSet[prefix] && (wordSet[suffix] || canFormByCombining(suffix, wordSet, memo)) {
			memo[str] = true
			return true
		}
	}
	memo[str] = false
	return false
}

func longestWord(words []string) string {
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j]) || (len(words[i]) == len(words[j]) && words[i] < words[j])
	})

	wordSet := make(map[string]bool)
	for _, word := range words {
		wordSet[word] = true
	}

	memo := make(map[string]bool)
	for _, word := range words {
		if canFormByCombining(word, wordSet, memo) {
			return word
		}
		if wordSet[word] {
			delete(wordSet, word) // 避免自身引用的情况
		}
	}
	return ""
}

func main() {
	words := []string{"test", "tester", "testertest", "testing", "apple", "seattle", "banana", "batting", "ngcat", "batti", "bat", "testingtester", "testbattingcat"}
	fmt.Println("The longest word that can be constructed from others is:", longestWord(words))
}

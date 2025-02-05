package main

import (
	"log"
)

func dfs(chars []rune, path []rune, used []bool, results *[]string) {
	if len(path) == len(chars) {
		*results = append(*results, string(path))
		return
	}
	for i, c := range chars {
		if !used[i] {
			used[i] = true
			path = append(path, c)
			dfs(chars, path, used, results)
			path = path[:len(path)-1]
			used[i] = false
		}
	}
}

func permuteDFS(s string) []string {
	chars := []rune(s)
	var results []string
	used := make([]bool, len(chars))
	dfs(chars, []rune{}, used, &results)
	return results
}

func main() {
	s := "abc"
	results := permuteDFS(s)
	log.Println(results)
}
